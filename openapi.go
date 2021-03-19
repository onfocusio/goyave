package goyave

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/getkin/kin-openapi/openapi3"
	"goyave.dev/goyave/v3/config"
	"goyave.dev/goyave/v3/validation"
)

var urlParamFormat = regexp.MustCompile(`{\w+(:.+?)?}`)

func (r *Router) SaveOpenAPISpec() { // TODO how to call this?
	spec := &openapi3.Swagger{
		OpenAPI: "3.0.0",
		Info: &openapi3.Info{
			Title:   config.GetString("app.name"),
			Version: "0.0.0",
		},
		Paths: make(openapi3.Paths),
	}

	convertRouter(r, spec)

	json, err := spec.MarshalJSON()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(json))
}

func convertRouter(r *Router, spec *openapi3.Swagger) {
	for _, route := range r.routes {
		for _, m := range route.methods {
			if m == http.MethodHead || m == http.MethodOptions {
				continue
			}
			op := openapi3.NewOperation()
			op.Responses = openapi3.NewResponses()
			// TODO handle OPTIONS response (with CORS)

			if route.validationRules != nil {
				hasBody := canHaveBody(m)
				// TODO generate schema ref instead of duplicating
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

			uri := route.BuildURL(params...)[len(BaseURL()):] // FIXME not optimized
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
