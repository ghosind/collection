package dict

import (
	"reflect"

	"github.com/ghosind/collection"
)

// HashDict is a Golang builtin map wrapper.
type HashDict[K comparable, V any] map[K]V

// NewHashDict creates a new HashDict.
func NewHashDict[K comparable, V any]() *HashDict[K, V] {
	d := make(HashDict[K, V])

	return (&d)
}

// Clone returns a copy of this dictionary.
func (m *HashDict[K, V]) Clone() collection.Dict[K, V] {
	newDict := make(HashDict[K, V], len(*m))

	for k, v := range *m {
		newDict[k] = v
	}

	return &newDict
}

// ContainsKey returns true if this dictionary contains a key-value pair with the specified key.
func (m *HashDict[K, V]) ContainsKey(k K) bool {
	_, ok := (*m)[k]

	return ok
}

// Equals compares this dictionary with the object pass from parameter.
func (m *HashDict[K, V]) Equals(o any) bool {
	om, ok := o.(*HashDict[K, V])
	if !ok {
		return false
	}

	if m.Size() != om.Size() {
		return false
	}

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

// ForEach performs the given handler for each key-value pairs in the dictionary until all pairs
// have been processed or the handler returns an error.
func (m *HashDict[K, V]) ForEach(handler func(K, V) error) error {
	for k, v := range *m {
		if err := handler(k, v); err != nil {
			return err
		}
	}

	return nil
}

// Get returns the value which associated to the specified key.
func (m *HashDict[K, V]) Get(k K) (V, bool) {
	v, ok := (*m)[k]
	return v, ok
}

// GetDefault returns the value associated with the specified key, and returns the default value if
// this dictionary contains no pair with the key.
func (m *HashDict[K, V]) GetDefault(k K, defaultVal V) V {
	v, ok := (*m)[k]
	if !ok {
		return defaultVal
	}

	return v
}

// IsEmpty returns true if this dictionary is empty.
func (m *HashDict[K, V]) IsEmpty() bool {
	return m.Size() == 0
}

// Keys returns a slice that contains all the keys in this dictionary.
func (m *HashDict[K, V]) Keys() []K {
	keys := make([]K, 0, len(*m))
	for k := range *m {
		keys = append(keys, k)
	}

	return keys
}

// Put associate the specified value with the specified key in this dictionary.
func (m *HashDict[K, V]) Put(k K, v V) V {
	old := (*m)[k]
	(*m)[k] = v

	return old
}

// Remove removes the key-value pair with the specified key.
func (m *HashDict[K, V]) Remove(k K) V {
	old := (*m)[k]
	delete((*m), k)

	return old
}

// Replace replaces the value for the specified key only if it is currently in this dictionary.
func (m *HashDict[K, V]) Replace(k K, v V) (V, bool) {
	old, ok := (*m)[k]
	if !ok {
		return old, false // zero value
	}

	(*m)[k] = v

	return old, true
}

// Size returns the number of key-value pairs in this dictionary.
func (m *HashDict[K, V]) Size() int {
	return len(*m)
}

// Values returns a slice that contains all the values in this dictionary.
func (m *HashDict[K, V]) Values() []V {
	arr := make([]V, 0, len(*m))

	for _, v := range *m {
		arr = append(arr, v)
	}

	return arr
}
