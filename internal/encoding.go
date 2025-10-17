package internal

import (
	"fmt"
	"strconv"
)

// ValueString returns the string representation of a value.
func ValueString(v any) string {
	switch vt := v.(type) {
	case int:
		return strconv.Itoa(vt)
	case int8:
		return strconv.FormatInt(int64(vt), 10)
	case int16:
		return strconv.FormatInt(int64(vt), 10)
	case int32:
		return strconv.FormatInt(int64(vt), 10)
	case int64:
		return strconv.FormatInt(vt, 10)
	case uint:
		return strconv.FormatUint(uint64(vt), 10)
	case uint8:
		return strconv.FormatUint(uint64(vt), 10)
	case uint16:
		return strconv.FormatUint(uint64(vt), 10)
	case uint32:
		return strconv.FormatUint(uint64(vt), 10)
	case uint64:
		return strconv.FormatUint(vt, 10)
	case uintptr:
		return strconv.FormatUint(uint64(vt), 10)
	case float32:
		return strconv.FormatFloat(float64(vt), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(vt, 'f', -1, 64)
	case complex64:
		return strconv.FormatComplex(complex128(vt), 'f', -1, 64)
	case complex128:
		return strconv.FormatComplex(vt, 'f', -1, 128)
	case bool:
		return strconv.FormatBool(vt)
	case string:
		return vt
	default:
		if sv, ok := v.(fmt.Stringer); ok {
			return sv.String()
		}
		return fmt.Sprintf("%v", v)
	}
}
