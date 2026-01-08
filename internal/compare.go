package internal

import (
	"reflect"
	"sync"
)

var (
	emptyStruct   = struct{}{}
	stringMapPool = sync.Pool{
		New: func() any {
			m := make(map[string]struct{})
			return m
		},
	}
	intMapPool = sync.Pool{
		New: func() any {
			m := make(map[int]struct{})
			return m
		},
	}
	int8MapPool = sync.Pool{
		New: func() any {
			m := make(map[int8]struct{})
			return m
		},
	}
	int16MapPool = sync.Pool{
		New: func() any {
			m := make(map[int16]struct{})
			return m
		},
	}
	int32MapPool = sync.Pool{
		New: func() any {
			m := make(map[int32]struct{})
			return m
		},
	}
	int64MapPool = sync.Pool{
		New: func() any {
			m := make(map[int64]struct{})
			return m
		},
	}
	uintMapPool = sync.Pool{
		New: func() any {
			m := make(map[uint]struct{})
			return m
		},
	}
	uint8MapPool = sync.Pool{
		New: func() any {
			m := make(map[uint8]struct{})
			return m
		},
	}
	uint16MapPool = sync.Pool{
		New: func() any {
			m := make(map[uint16]struct{})
			return m
		},
	}
	uint32MapPool = sync.Pool{
		New: func() any {
			m := make(map[uint32]struct{})
			return m
		},
	}
	uint64MapPool = sync.Pool{
		New: func() any {
			m := make(map[uint64]struct{})
			return m
		},
	}
	uintptrMapPool = sync.Pool{
		New: func() any {
			m := make(map[uintptr]struct{})
			return m
		},
	}
	float32MapPool = sync.Pool{
		New: func() any {
			m := make(map[float32]struct{})
			return m
		},
	}
	float64MapPool = sync.Pool{
		New: func() any {
			m := make(map[float64]struct{})
			return m
		},
	}
	complex64MapPool = sync.Pool{
		New: func() any {
			m := make(map[complex64]struct{})
			return m
		},
	}
	complex128MapPool = sync.Pool{
		New: func() any {
			m := make(map[complex128]struct{})
			return m
		},
	}
	boolMapPool = sync.Pool{
		New: func() any {
			m := make(map[bool]struct{})
			return m
		},
	}
)

func ReleaseCacheMap(cache any) {
	if cache == nil {
		return
	}

	switch cached := cache.(type) {
	case map[string]struct{}:
		for k := range cached {
			delete(cached, k)
		}
		stringMapPool.Put(cached)
	case map[int]struct{}:
		for k := range cached {
			delete(cached, k)
		}
		intMapPool.Put(cached)
	case map[int8]struct{}:
		for k := range cached {
			delete(cached, k)
		}
		int8MapPool.Put(cached)
	case map[int16]struct{}:
		for k := range cached {
			delete(cached, k)
		}
		int16MapPool.Put(cached)
	case map[int32]struct{}:
		for k := range cached {
			delete(cached, k)
		}
		int32MapPool.Put(cached)
	case map[int64]struct{}:
		for k := range cached {
			delete(cached, k)
		}
		int64MapPool.Put(cached)
	case map[uint]struct{}:
		for k := range cached {
			delete(cached, k)
		}
		uintMapPool.Put(cached)
	case map[uint8]struct{}:
		for k := range cached {
			delete(cached, k)
		}
		uint8MapPool.Put(cached)
	case map[uint16]struct{}:
		for k := range cached {
			delete(cached, k)
		}
		uint16MapPool.Put(cached)
	case map[uint32]struct{}:
		for k := range cached {
			delete(cached, k)
		}
		uint32MapPool.Put(cached)
	case map[uint64]struct{}:
		for k := range cached {
			delete(cached, k)
		}
		uint64MapPool.Put(cached)
	case map[uintptr]struct{}:
		for k := range cached {
			delete(cached, k)
		}
		uintptrMapPool.Put(cached)
	case map[float32]struct{}:
		for k := range cached {
			delete(cached, k)
		}
		float32MapPool.Put(cached)
	case map[float64]struct{}:
		for k := range cached {
			delete(cached, k)
		}
		float64MapPool.Put(cached)
	case map[complex64]struct{}:
		for k := range cached {
			delete(cached, k)
		}
		complex64MapPool.Put(cached)

	case map[complex128]struct{}:
		for k := range cached {
			delete(cached, k)
		}
		complex128MapPool.Put(cached)
	case map[bool]struct{}:
		for k := range cached {
			delete(cached, k)
		}
		boolMapPool.Put(cached)
	}
}

func MakeSliceCacheMap[T any](s []T) any {
	if len(s) == 0 {
		return nil
	}

	switch any(s[0]).(type) {
	case string:
		cache := stringMapPool.Get().(map[string]struct{})
		for _, e := range s {
			cache[any(e).(string)] = emptyStruct
		}
		return cache
	case int:
		cache := intMapPool.Get().(map[int]struct{})
		for _, e := range s {
			cache[any(e).(int)] = emptyStruct
		}
		return cache
	case int8:
		cache := int8MapPool.Get().(map[int8]struct{})
		for _, e := range s {
			cache[any(e).(int8)] = emptyStruct
		}
		return cache
	case int16:
		cache := int16MapPool.Get().(map[int16]struct{})
		for _, e := range s {
			cache[any(e).(int16)] = emptyStruct
		}
		return cache
	case int32:
		cache := int32MapPool.Get().(map[int32]struct{})
		for _, e := range s {
			cache[any(e).(int32)] = emptyStruct
		}
		return cache
	case int64:
		cache := int64MapPool.Get().(map[int64]struct{})
		for _, e := range s {
			cache[any(e).(int64)] = emptyStruct
		}
		return cache
	case uint:
		cache := uintMapPool.Get().(map[uint]struct{})
		for _, e := range s {
			cache[any(e).(uint)] = emptyStruct
		}
		return cache
	case uint8:
		cache := uint8MapPool.Get().(map[uint8]struct{})
		for _, e := range s {
			cache[any(e).(uint8)] = emptyStruct
		}
		return cache
	case uint16:
		cache := uint16MapPool.Get().(map[uint16]struct{})
		for _, e := range s {
			cache[any(e).(uint16)] = emptyStruct
		}
		return cache
	case uint32:
		cache := uint32MapPool.Get().(map[uint32]struct{})
		for _, e := range s {
			cache[any(e).(uint32)] = emptyStruct
		}
		return cache
	case uint64:
		cache := uint64MapPool.Get().(map[uint64]struct{})
		for _, e := range s {
			cache[any(e).(uint64)] = emptyStruct
		}
		return cache
	case uintptr:
		cache := uintptrMapPool.Get().(map[uintptr]struct{})
		for _, e := range s {
			cache[any(e).(uintptr)] = emptyStruct
		}
		return cache
	case float32:
		cache := float32MapPool.Get().(map[float32]struct{})
		for _, e := range s {
			cache[any(e).(float32)] = emptyStruct
		}
		return cache
	case float64:
		cache := float64MapPool.Get().(map[float64]struct{})
		for _, e := range s {
			cache[any(e).(float64)] = emptyStruct
		}
		return cache
	case complex64:
		cache := complex64MapPool.Get().(map[complex64]struct{})
		for _, e := range s {
			cache[any(e).(complex64)] = emptyStruct
		}
		return cache
	case complex128:
		cache := complex128MapPool.Get().(map[complex128]struct{})
		for _, e := range s {
			cache[any(e).(complex128)] = emptyStruct
		}
		return cache
	case bool:
		cache := boolMapPool.Get().(map[bool]struct{})
		for _, e := range s {
			cache[any(e).(bool)] = emptyStruct
		}
		return cache
	default:
		return nil
	}
}

func InSlice[T any](e T, s []T, cache any) bool {
	if cache == nil {
		goto fallback
	}

	switch va := any(e).(type) {
	case string:
		cached, _ := cache.(map[string]struct{})
		if cached != nil {
			_, found := cached[va]
			return found
		}
	case int:
		cached, _ := cache.(map[int]struct{})
		if cached != nil {
			_, found := cached[va]
			return found
		}
	case int8:
		cached, _ := cache.(map[int8]struct{})
		if cached != nil {
			_, found := cached[va]
			return found
		}
	case int16:
		cached, _ := cache.(map[int16]struct{})
		if cached != nil {
			_, found := cached[va]
			return found
		}
	case int32:
		cached, _ := cache.(map[int32]struct{})
		if cached != nil {
			_, found := cached[va]
			return found
		}
	case int64:
		cached, _ := cache.(map[int64]struct{})
		if cached != nil {
			_, found := cached[va]
			return found
		}
	case uint:
		cached, _ := cache.(map[uint]struct{})
		if cached != nil {
			_, found := cached[va]
			return found
		}
	case uint8:
		cached, _ := cache.(map[uint8]struct{})
		if cached != nil {
			_, found := cached[va]
			return found
		}
	case uint16:
		cached, _ := cache.(map[uint16]struct{})
		if cached != nil {
			_, found := cached[va]
			return found
		}
	case uint32:
		cached, _ := cache.(map[uint32]struct{})
		if cached != nil {
			_, found := cached[va]
			return found
		}
	case uint64:
		cached, _ := cache.(map[uint64]struct{})
		if cached != nil {
			_, found := cached[va]
			return found
		}
	case uintptr:
		cached, _ := cache.(map[uintptr]struct{})
		if cached != nil {
			_, found := cached[va]
			return found
		}
	case float32:
		cached, _ := cache.(map[float32]struct{})
		if cached != nil {
			_, found := cached[va]
			return found
		}
	case float64:
		cached, _ := cache.(map[float64]struct{})
		if cached != nil {
			_, found := cached[va]
			return found
		}
	case complex64:
		cached, _ := cache.(map[complex64]struct{})
		if cached != nil {
			_, found := cached[va]
			return found
		}
	case complex128:
		cached, _ := cache.(map[complex128]struct{})
		if cached != nil {
			_, found := cached[va]
			return found
		}
	case bool:
		cached, _ := cache.(map[bool]struct{})
		if cached != nil {
			_, found := cached[va]
			return found
		}
	default:
		// types without a cache fall through to the linear search below
	}

fallback:
	for _, v := range s {
		if Equal(v, e) {
			return true
		}
	}

	return false
}

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
