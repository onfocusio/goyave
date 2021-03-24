package goyave

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"net/http"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"goyave.dev/goyave/v3/config"
	"goyave.dev/goyave/v3/validation"
)

var urlParamFormat = regexp.MustCompile(`{\w+(:.+?)?}`)

// SaveOpenAPISpec PROTOTYPE function printing generated OpenAPI 3 spec to stdout
func (r *Router) SaveOpenAPISpec() { // TODO how to call this?
	spec := &openapi3.Swagger{
		OpenAPI: "3.0.0",
		Info: &openapi3.Info{
			Title:   config.GetString("app.name"),
			Version: "0.0.0",
		},
		Paths:   make(openapi3.Paths),
		Servers: makeServers(),
	}

	convertRouter(r, spec)

	json, err := spec.MarshalJSON()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(json))
}

func makeServers() openapi3.Servers {
	return openapi3.Servers{
		&openapi3.Server{
			URL: BaseURL(),
		},
	}
}

func convertRouter(r *Router, spec *openapi3.Swagger) {
	for _, route := range r.routes {

		desc, annotations := readDescription(route.handler)

		for _, m := range route.methods {
			if m == http.MethodHead || m == http.MethodOptions {
				continue
			}
			op := openapi3.NewOperation()
			op.Description = desc
			// TODO handle OPTIONS response (with CORS)

			if route.validationRules != nil {
				hasBody := canHaveBody(m)
				// TODO generate schema ref instead of duplicating
				// But it's a bit hard to identify identical schemas and parameters
				if hasBody {
					schema := openapi3.NewObjectSchema()
					for name, field := range route.validationRules.Fields {
						schema.Properties[name] = &openapi3.SchemaRef{Value: makeSchemaFromField(field)}
						if field.IsRequired() {
							schema.Required = append(schema.Required, name)
						}
					}

					var content openapi3.Content
					if hasFile(route.validationRules) {
						content = openapi3.NewContentWithFormDataSchema(schema)
						if hasOnlyOptionalFiles(route.validationRules) {
							jsonSchema := openapi3.NewObjectSchema()
							jsonSchema.Required = schema.Required
							for name, prop := range schema.Properties {
								if prop.Value.Format != "binary" && prop.Value.Format != "bytes" {
									jsonSchema.Properties[name] = prop
								}
							}
							content["application/json"] = openapi3.NewMediaType().WithSchema(jsonSchema)
						}
					} else {
						content = openapi3.NewContentWithJSONSchema(schema)
					}
					body := openapi3.NewRequestBody().WithContent(content)
					if hasRequired(route.validationRules) {
						body.Required = true
					}
					op.RequestBody = &openapi3.RequestBodyRef{
						Value: body,
					}
				} else {
					for name, field := range route.validationRules.Fields {
						param := openapi3.NewQueryParameter(name)
						param.Schema = &openapi3.SchemaRef{Value: makeSchemaFromField(field)}
						format := param.Schema.Value.Format
						if format != "binary" && format != "bytes" {
							param.Required = field.IsRequired()
							op.Parameters = append(op.Parameters, &openapi3.ParameterRef{Value: param})
						}
					}
				}
			}

			// Regex are not allowed in URI, generate it without format definition
			params := make([]string, 0, len(route.parameters))
			for _, p := range route.parameters {
				params = append(params, "{"+p+"}")
			}

			uri := route.BuildURI(params...)
			if i := strings.Index(uri[1:], "/"); i != -1 {
				op.Tags = []string{uri[1 : i+1]}
			} else {
				op.Tags = []string{uri[1:]}
			}
			op.Responses = openapi3.Responses{}
			processAnnotations(op, annotations)
			if len(op.Responses) == 0 {
				op.Responses["default"] = &openapi3.ResponseRef{Value: openapi3.NewResponse().WithDescription("")}
			}
			spec.AddOperation(uri, m, op)
			path := spec.Paths[uri]
			formats := urlParamFormat.FindAllStringSubmatch(route.GetFullURI(), -1)
			for i, p := range route.parameters {
				if parameterExists(path, p) {
					continue
				}
				param := openapi3.NewPathParameter(p)
				schema := openapi3.NewStringSchema()
				if len(formats[i]) == 2 {
					schema.Pattern = formats[i][1]
					if schema.Pattern != "" {
						// Strip the colon
						schema.Pattern = schema.Pattern[1:]
					}
					if schema.Pattern == "[0-9]+" {
						schema.Type = "integer"
					}
				}
				param.Schema = &openapi3.SchemaRef{Value: schema}
				ref := &openapi3.ParameterRef{Value: param}
				path.Parameters = append(path.Parameters, ref)
			}
		}
	}

	for _, subrouter := range r.subrouters {
		convertRouter(subrouter, spec)
	}
}

func processAnnotations(op *openapi3.Operation, annotations []*annotation) {
	for _, a := range annotations { // TODO better annotations architecture
		switch a.Type {
		case "Response":
			params := strings.Split(a.Value, "\t")
			r := responseAnnotation{
				Name:        params[0],
				Description: params[1],
			}
			op.Responses[r.Name] = &openapi3.ResponseRef{
				Value: &openapi3.Response{
					Description: &r.Description,
				},
			}
		default:
			fmt.Println("WARNING: unsupported annotation:", a.Type)
		}
	}
}

func findFirstTypeRule(field *validation.Field) *validation.Rule {
	for _, rule := range field.Rules {
		if rule.IsType() || rule.Name == "file" {
			return rule
		}
	}
	return nil
}

func canHaveBody(method string) bool {
	return method == http.MethodDelete ||
		method == http.MethodPatch ||
		method == http.MethodPost ||
		method == http.MethodPut
}

func parameterExists(path *openapi3.PathItem, param string) bool {
	for _, p := range path.Parameters { // TODO if refs used?
		if p.Value.Name == param {
			return true
		}
	}
	return false
}

func hasFile(rules *validation.Rules) bool {
	return has(rules, "file")
}

func hasRequired(rules *validation.Rules) bool {
	return has(rules, "required")
}

func has(rules *validation.Rules, ruleName string) bool {
	for _, f := range rules.Fields {
		for _, r := range f.Rules {
			if r.Name == ruleName {
				return true
			}
		}
	}
	return false
}

func hasOnlyOptionalFiles(rules *validation.Rules) bool {
	for _, f := range rules.Fields {
		for _, r := range f.Rules {
			if r.Name == "file" && f.IsRequired() {
				return false
			}
		}
	}
	return true
}

func makeSchemaFromField(field *validation.Field) *openapi3.Schema {
	s := openapi3.NewSchema()
	if rule := findFirstTypeRule(field); rule != nil {
		switch rule.Name {
		case "numeric":
			s.Type = "number"
		case "bool":
			s.Type = "boolean"
		case "file":
			// TODO format "binary" (or "bytes" ?)
			s.Type = "string"
			s.Format = "binary"
		case "array": // TODO multidimensional arrays
			s.Type = "array"
			schema := openapi3.NewSchema()
			schema.Type = ruleNameToType(rule.Name)
			s.Items = &openapi3.SchemaRef{Value: schema}
		// TODO objects and string formats
		// TODO email, uuid, uri, ipv4, ipv6, date, date-time (not types but patterns)
		default:
			s.Type = rule.Name
		}
	}

	for _, r := range field.Rules {
		convertRule(r, s)
	}
	s.Nullable = field.IsNullable()
	return s
}

func ruleNameToType(name string) string {
	switch name {
	case "numeric":
		return "number"
	case "bool":
		return "boolean"
	case "file":
		return "string"
	default:
		return name
	}
	// TODO match type rules with correct openapi types defined in spec
}

func convertRule(r *validation.Rule, s *openapi3.Schema) {
	// TODO minimum, maximum, string formats, arrays, uniqueItems (distinct)
	switch r.Name {
	case "min":
		switch s.Type {
		case "string":
			min, _ := strconv.ParseUint(r.Params[0], 10, 64)
			s.MinLength = min
		case "number", "integer":
			min, _ := strconv.ParseFloat(r.Params[0], 64)
			s.Min = &min
		case "array":
			min, _ := strconv.ParseUint(r.Params[0], 10, 64)
			s.MinItems = min
		}
	case "max":
		switch s.Type {
		case "string":
			max, _ := strconv.ParseUint(r.Params[0], 10, 64)
			s.MaxLength = &max
		case "number", "integer":
			max, _ := strconv.ParseFloat(r.Params[0], 64)
			s.Max = &max
		case "array":
			max, _ := strconv.ParseUint(r.Params[0], 10, 64)
			s.MaxItems = &max
		}
	}
}

func readDescription(handler Handler) (string, []*annotation) {
	pc := reflect.ValueOf(handler).Pointer()
	handlerValue := runtime.FuncForPC(pc)
	file, _ := handlerValue.FileLine(pc)
	funcName := handlerValue.Name()

	src, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	fset := token.NewFileSet() // positions are relative to fset

	f, err := parser.ParseFile(fset, file, src, parser.ParseComments)

	if err != nil {
		panic(err)
	}

	var doc *ast.CommentGroup

	// TODO optimize, this re-inspects the whole file for each route. Maybe cache already inspected files
	ast.Inspect(f, func(n ast.Node) bool { // TODO what would it do with closures and implementations?
		// Example output of "funcName" value for controller: goyave.dev/goyave/v3/auth.(*JWTController).Login-fm
		fn, ok := n.(*ast.FuncDecl)
		if ok {
			if fn.Name.IsExported() {
				if fn.Recv != nil {
					for _, f := range fn.Recv.List {
						if expr, ok := f.Type.(*ast.StarExpr); ok {
							if id, ok := expr.X.(*ast.Ident); ok {
								strct := fmt.Sprintf("(*%s)", id.Name) // TODO handle expr without star (no ptr)
								name := funcName[:len(funcName)-3]     // strip -fm
								expectedName := strct + "." + fn.Name.Name
								if name[len(name)-len(expectedName):] == expectedName {
									doc = fn.Doc
									return false
								}
							}
						}
					}
				}
				lastIndex := strings.LastIndex(funcName, ".")
				if funcName[lastIndex+1:] == fn.Name.Name {
					doc = fn.Doc
					return false
				}
			}
		}
		return true
	})

	if doc != nil {
		annotations := []*annotation{}
		text := ""
		for _, line := range strings.Split(doc.Text(), "\n") {
			trimmed := strings.TrimSpace(line)
			if len(trimmed) >= 1 && trimmed[0] == '@' {
				i := strings.Index(trimmed, " ")
				if i != -1 {
					annotations = append(annotations, &annotation{
						Type:  trimmed[1:i],
						Value: strings.TrimSpace(trimmed[i:]),
					})
					continue
				}
			}
			// TODO find better way to specify responses
			// FIXME only works with one-line

			text += "\n" + line
		}
		return strings.TrimSpace(text), annotations
	}

	return "", []*annotation{}
}

type annotation struct {
	Type  string
	Value string
}

type responseAnnotation struct {
	Name        string
	Description string
}
