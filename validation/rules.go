package validation

import (
	"math"
	"reflect"
	"strconv"
	"strings"

	"github.com/Code-Hex/uniseg"
	"goyave.dev/goyave/v4/helper"
	"goyave.dev/goyave/v4/helper/filesystem"
)

func validateRequired(ctx *Context) bool {
	if !ctx.Field.IsNullable() && ctx.Value == nil {
		return false
	}
	if str, okStr := ctx.Value.(string); okStr && str == "" {
		return false
	}
	return true
}

func validateMin(ctx *Context) bool {
	min, err := strconv.ParseFloat(ctx.Rule.Params[0], 64)
	if err != nil {
		panic(err)
	}
	switch GetFieldType(ctx.Value) {
	case "numeric":
		floatValue, _ := helper.ToFloat64(ctx.Value)
		return floatValue >= min
	case "string":
		return uniseg.GraphemeClusterCount(ctx.Value.(string)) >= int(min)
	case "array":
		list := reflect.ValueOf(ctx.Value)
		return list.Len() >= int(min)
	case "file":
		files, _ := ctx.Value.([]filesystem.File)
		for _, file := range files {
			if file.Header.Size < int64(min)*1024 {
				return false
			}
		}
		return true
	}

	return true // Pass if field type cannot be checked (bool, dates, ...)
}

func validateMax(ctx *Context) bool {
	max, err := strconv.ParseFloat(ctx.Rule.Params[0], 64)
	if err != nil {
		panic(err)
	}
	switch GetFieldType(ctx.Value) {
	case "numeric":
		floatValue, _ := helper.ToFloat64(ctx.Value)
		return floatValue <= max
	case "string":
		return uniseg.GraphemeClusterCount(ctx.Value.(string)) <= int(max)
	case "array":
		list := reflect.ValueOf(ctx.Value)
		return list.Len() <= int(max)
	case "file":
		files, _ := ctx.Value.([]filesystem.File)
		for _, file := range files {
			if file.Header.Size > int64(max)*1024 {
				return false
			}
		}
		return true
	}

	return true // Pass if field type cannot be checked (bool, dates, ...)
}

func validateBetween(ctx *Context) bool {
	min, errMin := strconv.ParseFloat(ctx.Rule.Params[0], 64)
	max, errMax := strconv.ParseFloat(ctx.Rule.Params[1], 64)
	if errMin != nil {
		panic(errMin)
	}
	if errMax != nil {
		panic(errMax)
	}

	switch GetFieldType(ctx.Value) {
	case "numeric":
		floatValue, _ := helper.ToFloat64(ctx.Value)
		return floatValue >= min && floatValue <= max
	case "string":
		length := uniseg.GraphemeClusterCount(ctx.Value.(string))
		return length >= int(min) && length <= int(max)
	case "array":
		list := reflect.ValueOf(ctx.Value)
		length := list.Len()
		return length >= int(min) && length <= int(max)
	case "file":
		files, _ := ctx.Value.([]filesystem.File)
		for _, file := range files {
			minSize := int64(min) * 1024
			maxSize := int64(max) * 1024
			if file.Header.Size < minSize || file.Header.Size > maxSize {
				return false
			}
		}
		return true
	}

	return true // Pass if field type cannot be checked (bool, dates, ...)
}

func validateGreaterThan(ctx *Context) bool {
	valueType := GetFieldType(ctx.Value)

	_, compared, _, exists := GetFieldFromName(ctx.Rule.Params[0], ctx.Data)
	if !exists || valueType != GetFieldType(compared) {
		return false // Can't compare two different types or missing field
	}

	switch valueType {
	case "numeric":
		floatValue, _ := helper.ToFloat64(ctx.Value)
		comparedFloatValue, _ := helper.ToFloat64(compared)
		return floatValue > comparedFloatValue
	case "string":
		return uniseg.GraphemeClusterCount(ctx.Value.(string)) > uniseg.GraphemeClusterCount(compared.(string))
	case "array":
		return reflect.ValueOf(ctx.Value).Len() > reflect.ValueOf(compared).Len()
	case "file":
		files, _ := ctx.Value.([]filesystem.File)
		comparedFiles, _ := compared.([]filesystem.File)
		for _, file := range files {
			for _, comparedFile := range comparedFiles {
				if file.Header.Size <= comparedFile.Header.Size {
					return false
				}
			}
		}
		return true
	}

	return true // Pass if field type cannot be checked (bool, dates, ...)
}

func validateGreaterThanEqual(ctx *Context) bool {
	valueType := GetFieldType(ctx.Value)

	_, compared, _, exists := GetFieldFromName(ctx.Rule.Params[0], ctx.Data)
	if !exists || valueType != GetFieldType(compared) {
		return false // Can't compare two different types or missing field
	}

	switch valueType {
	case "numeric":
		floatValue, _ := helper.ToFloat64(ctx.Value)
		comparedFloatValue, _ := helper.ToFloat64(compared)
		return floatValue >= comparedFloatValue
	case "string":
		return uniseg.GraphemeClusterCount(ctx.Value.(string)) >= uniseg.GraphemeClusterCount(compared.(string))
	case "array":
		return reflect.ValueOf(ctx.Value).Len() >= reflect.ValueOf(compared).Len()
	case "file":
		files, _ := ctx.Value.([]filesystem.File)
		comparedFiles, _ := compared.([]filesystem.File)
		for _, file := range files {
			for _, comparedFile := range comparedFiles {
				if file.Header.Size < comparedFile.Header.Size {
					return false
				}
			}
		}
		return true
	}

	return true // Pass if field type cannot be checked (bool, dates, ...)
}

func validateLowerThan(ctx *Context) bool {
	valueType := GetFieldType(ctx.Value)

	_, compared, _, exists := GetFieldFromName(ctx.Rule.Params[0], ctx.Data)
	if !exists || valueType != GetFieldType(compared) {
		return false // Can't compare two different types or missing field
	}

	switch valueType {
	case "numeric":
		floatValue, _ := helper.ToFloat64(ctx.Value)
		comparedFloatValue, _ := helper.ToFloat64(compared)
		return floatValue < comparedFloatValue
	case "string":
		return uniseg.GraphemeClusterCount(ctx.Value.(string)) < uniseg.GraphemeClusterCount(compared.(string))
	case "array":
		return reflect.ValueOf(ctx.Value).Len() < reflect.ValueOf(compared).Len()
	case "file":
		files, _ := ctx.Value.([]filesystem.File)
		comparedFiles, _ := compared.([]filesystem.File)
		for _, file := range files {
			for _, comparedFile := range comparedFiles {
				if file.Header.Size >= comparedFile.Header.Size {
					return false
				}
			}
		}
		return true
	}

	return true // Pass if field type cannot be checked (bool, dates, ...)
}

func validateLowerThanEqual(ctx *Context) bool {
	valueType := GetFieldType(ctx.Value)

	_, compared, _, exists := GetFieldFromName(ctx.Rule.Params[0], ctx.Data)
	if !exists || valueType != GetFieldType(compared) {
		return false // Can't compare two different types or missing field
	}

	switch valueType {
	case "numeric":
		floatValue, _ := helper.ToFloat64(ctx.Value)
		comparedFloatValue, _ := helper.ToFloat64(compared)
		return floatValue <= comparedFloatValue
	case "string":
		return uniseg.GraphemeClusterCount(ctx.Value.(string)) <= uniseg.GraphemeClusterCount(compared.(string))
	case "array":
		return reflect.ValueOf(ctx.Value).Len() <= reflect.ValueOf(compared).Len()
	case "file":
		files, _ := ctx.Value.([]filesystem.File)
		comparedFiles, _ := compared.([]filesystem.File)
		for _, file := range files {
			for _, comparedFile := range comparedFiles {
				if file.Header.Size > comparedFile.Header.Size {
					return false
				}
			}
		}
		return true
	}

	return true // Pass if field type cannot be checked (bool, dates, ...)
}

func validateBool(ctx *Context) bool {
	rv := reflect.ValueOf(ctx.Value)
	kind := rv.Kind().String()
	switch {
	case kind == "bool":
		return true
	case strings.HasPrefix(kind, "int"), strings.HasPrefix(kind, "uint") && kind != "uintptr":
		v, _ := helper.ToFloat64(ctx.Value)
		if v == 1 {
			ctx.Value = true
			return true
		} else if v == 0 {
			ctx.Value = false
			return true
		}
	case kind == "string":
		v, _ := ctx.Value.(string)
		switch v {
		case "1", "on", "true", "yes":
			ctx.Value = true
			return true
		case "0", "off", "false", "no":
			ctx.Value = false
			return true
		}
	}
	return false
}

func validateSame(ctx *Context) bool {
	_, other, _, exists := GetFieldFromName(ctx.Rule.Params[0], ctx.Data)
	if exists {
		valueType := GetFieldType(ctx.Value)
		otherType := GetFieldType(other)
		if valueType == otherType {
			switch valueType {
			case "numeric":
				f1, _ := helper.ToFloat64(ctx.Value)
				f2, _ := helper.ToFloat64(other)
				return f1 == f2
			case "string":
				s1, _ := ctx.Value.(string)
				s2, _ := other.(string)
				return s1 == s2
			case "array":
				return helper.SliceEqual(ctx.Value, other)
			case "object":
				return reflect.DeepEqual(ctx.Value, other)
			}
			// Don't check files
		}
	}
	return false
}

func validateDifferent(ctx *Context) bool {
	return !validateSame(ctx)
}

func validateSize(ctx *Context) bool {
	size, err := strconv.Atoi(ctx.Rule.Params[0])
	if err != nil {
		panic(err)
	}

	switch GetFieldType(ctx.Value) {
	case "numeric":
		floatVal, _ := helper.ToFloat64(ctx.Value)
		return floatVal == float64(size)
	case "string":
		return uniseg.GraphemeClusterCount(ctx.Value.(string)) == size
	case "array":
		list := reflect.ValueOf(ctx.Value)
		return list.Len() == size
	case "file":
		files, _ := ctx.Value.([]filesystem.File)
		for _, file := range files {
			if int64(math.Round(float64(file.Header.Size)/1024.0)) != int64(size) {
				return false
			}
		}
		return true
	}

	return true // Pass if field type cannot be checked (bool, dates, ...)
}

func validateObject(ctx *Context) bool {
	_, ok := ctx.Value.(map[string]interface{})
	return ok
}
