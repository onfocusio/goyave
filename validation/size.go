package validation

import (
	"math"
	"reflect"
	"strconv"

	"github.com/Code-Hex/uniseg"
	"goyave.dev/goyave/v4/util/fsutil"
)

func validateSizeV5(ctx *ContextV5, v func(size int) bool) bool {
	val := reflect.ValueOf(ctx.Value)
	switch getFieldType(val) {
	// TODO document it doesn't support numbers (because it wouldn't make a lot of sense)
	case FieldTypeString:
		return v(uniseg.GraphemeClusterCount(ctx.Value.(string)))
	case FieldTypeArray, FieldTypeObject: // TODO document it also works for objects (number of keys)
		return v(val.Len())
	case FieldTypeFile:
		files, _ := ctx.Value.([]fsutil.File)
		for _, file := range files {
			if !v(int(math.Ceil(float64(file.Header.Size) / 1024.0))) {
				return false
			}
		}
	}
	return true // Pass if field type cannot be checked (bool, dates, ...)
}

func numberAsFloat64(n any) (float64, bool) {
	switch val := n.(type) {
	case float32:
		return float64(val), true
	case float64:
		return val, true
	case int:
		return float64(val), true
	case int8:
		return float64(val), true
	case int16:
		return float64(val), true
	case int32:
		return float64(val), true
	case int64:
		return float64(val), true
	case uint:
		if val >= math.MaxInt64 {
			return float64(val), false
		}
		return float64(val), true
	case uint8:
		return float64(val), true
	case uint16:
		return float64(val), true
	case uint32:
		return float64(val), true
	case uint64:
		if val >= math.MaxInt64 {
			return float64(val), false
		}
		return float64(val), true
	}
	return 0, false
}

// SizeValidator validates the field under validation depending on its type.
//   - Strings must have a length of n characters (calculated based on the number of grapheme clusters)
//   - Arrays must have n elements
//   - Objects must have n keys
//   - Files must weight n KiB (for multi-files, all files must match this criteria). The number of KiB is rounded up (ceil).
type SizeValidator struct {
	BaseValidator
	Size int
}

// Validate checks the field under validation satisfies this validator's criteria.
func (v *SizeValidator) Validate(ctx *ContextV5) bool {
	return validateSizeV5(ctx, func(size int) bool {
		return size == v.Size
	})
}

// Name returns the string name of the validator.
func (v *SizeValidator) Name() string { return "size" }

// IsTypeDependent returns true
func (v *SizeValidator) IsTypeDependent() bool { return true }

// MessagePlaceholders returns the ":value" placeholder.
func (v *SizeValidator) MessagePlaceholders(ctx *ContextV5) []string {
	return []string{
		":value", strconv.Itoa(v.Size),
	}
}

// Size validates the field under validation depending on its type.
//   - Strings must have a length of n characters (calculated based on the number of grapheme clusters)
//   - Arrays must have n elements
//   - Objects must have n keys
//   - Files must weight n KiB (for multi-files, all files must match this criteria). The number of KiB is rounded up (ceil).
func Size(size int) *SizeValidator {
	return &SizeValidator{Size: size}
}
