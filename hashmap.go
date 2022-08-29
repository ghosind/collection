package collection

import (
	"reflect"

	"github.com/ghosind/utils"
)

// HashMap is a Golang builtin map wrapper.
type HashMap[K comparable, V any] map[K]V

// NewHashMap creates a new HashMap.
func NewHashMap[K comparable, V any]() *HashMap[K, V] {
	m := new(HashMap[K, V])
	*m = make(map[K]V)

	return m
}

// Clear removes all key-value pairs in this map.
func (m *HashMap[K, V]) Clear() {
	*m = make(HashMap[K, V])
}

// ContainsKey returns true if this map contains a key-value pair with the specified key.
func (m *HashMap[K, V]) ContainsKey(k K) bool {
	_, ok := (*m)[k]

	return ok
}

// Equals compares this map with the object pass from parameter.
func (m *HashMap[K, V]) Equals(o any) bool {
	if !utils.IsSameType(m, o) {
		return false
	}

	om := o.(*HashMap[K, V])
	for k, v := range *m {
		val, ok := (*om)[k]
		if !ok {
			return false
		}

		if !reflect.DeepEqual(v, val) {
			return false
		}
	}

	return true
}

// ForEach performs the given handler for each key-value pairs in the map until all pairs have
// been processed or the handler returns an error.
func (m *HashMap[K, V]) ForEach(handler func(K, V) error) error {
	for k, v := range *m {
		if err := handler(k, v); err != nil {
			return err
		}
	}

	return nil
}

// Get returns the value which associated to the specified key.
func (m *HashMap[K, V]) Get(k K) (V, bool) {
	v, ok := (*m)[k]
	return v, ok
}

// IsEmpty returns true if this map is empty.
func (m *HashMap[K, V]) IsEmpty() bool {
	return m.Size() == 0
}

// Put associate the specified value with the specified key in this map.
func (m *HashMap[K, V]) Put(k K, v V) V {
	old := (*m)[k]
	(*m)[k] = v

	return old
}

// Remove removes the key-value pair with the specified key.
func (m *HashMap[K, V]) Remove(k K) V {
	old := (*m)[k]
	delete((*m), k)

	return old
}

// Size returns the number of key-value pairs in this map.
func (m *HashMap[K, V]) Size() int {
	return len(*m)
}
