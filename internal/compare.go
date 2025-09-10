package internal

import "reflect"

// Equal reports whether a and b are equal.
func Equal[T any](a, b T) bool {
	switch va := any(a).(type) {
	case string:
		vb := any(b).(string)
		return va == vb
	case int:
		vb := any(b).(int)
		return va == vb
	case int8:
		vb := any(b).(int8)
		return va == vb
	case int16:
		vb := any(b).(int16)
		return va == vb
	case int32:
		vb := any(b).(int32)
		return va == vb
	case int64:
		vb := any(b).(int64)
		return va == vb
	case uint:
		vb := any(b).(uint)
		return va == vb
	case uint8:
		vb := any(b).(uint8)
		return va == vb
	case uint16:
		vb := any(b).(uint16)
		return va == vb
	case uint32:
		vb := any(b).(uint32)
		return va == vb
	case uint64:
		vb := any(b).(uint64)
		return va == vb
	case uintptr:
		vb := any(b).(uintptr)
		return va == vb
	case float32:
		vb := any(b).(float32)
		return va == vb
	case float64:
		vb := any(b).(float64)
		return va == vb
	case complex64:
		vb := any(b).(complex64)
		return va == vb
	case complex128:
		vb := any(b).(complex128)
		return va == vb
	case bool:
		vb := any(b).(bool)
		return va == vb
	default:
		return reflect.DeepEqual(a, b)
	}
}
