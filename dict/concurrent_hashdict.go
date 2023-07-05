package dict

import (
	"reflect"
	"sync"

	"github.com/ghosind/collection"
	"github.com/ghosind/utils"
)

// ConcurrentHashDictionary is the thread-safe hash dictionary implementation.
type ConcurrentHashDictionary[K comparable, V any] struct {
	data *HashDictionary[K, V]

	mutex sync.RWMutex
}

// NewConcurrentHashDictionary creates a new ConcurrentHashDictionary.
func NewConcurrentHashDictionary[K comparable, V any]() *ConcurrentHashDictionary[K, V] {
	newDict := new(ConcurrentHashDictionary[K, V])
	newDict.data = NewHashDictionary[K, V]()

	return newDict
}

// Clear removes all key-value pairs in this dictionary.
func (m *ConcurrentHashDictionary[K, V]) Clear() {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.data = NewHashDictionary[K, V]()
}

// Clone returns a copy of this dictionary.
func (m *ConcurrentHashDictionary[K, V]) Clone() collection.Dictionary[K, V] {
	newDict := NewConcurrentHashDictionary[K, V]()

	for k, v := range *m.data {
		(*newDict.data)[k] = v
	}

	return newDict
}

// ContainsKey returns true if this dictionary contains a key-value pair with the specified key.
func (m *ConcurrentHashDictionary[K, V]) ContainsKey(k K) bool {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	_, ok := (*m.data)[k]

	return ok
}

// Equals compares this dictionary with the object pass from parameter.
func (m *ConcurrentHashDictionary[K, V]) Equals(o any) bool {
	if !utils.IsSameType(m, o) {
		return false
	}

	om := o.(*ConcurrentHashDictionary[K, V])
	m.mutex.RLock()
	om.mutex.RLock()
	defer m.mutex.RUnlock()
	defer om.mutex.RUnlock()

	for k, v := range *m.data {
		val, ok := (*om.data)[k]
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
func (m *ConcurrentHashDictionary[K, V]) ForEach(handler func(K, V) error) error {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	for k, v := range *m.data {
		if err := handler(k, v); err != nil {
			return err
		}
	}

	return nil
}

// Get returns the value which associated to the specified key.
func (m *ConcurrentHashDictionary[K, V]) Get(k K) (V, bool) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	v, ok := (*m.data)[k]
	return v, ok
}

// GetDefault returns the value associated with the specified key, and returns the default value if
// this dictionary contains no pair with the key.
func (m *ConcurrentHashDictionary[K, V]) GetDefault(k K, defaultVal V) V {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	v, ok := (*m.data)[k]
	if !ok {
		return defaultVal
	}

	return v
}

// IsEmpty returns true if this dictionary is empty.
func (m *ConcurrentHashDictionary[K, V]) IsEmpty() bool {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	return m.data.Size() == 0
}

// Keys returns a slice that contains all the keys in this dictionary.
func (m *ConcurrentHashDictionary[K, V]) Keys() []K {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	keys := make([]K, 0, len(*m.data))
	for k := range *m.data {
		keys = append(keys, k)
	}

	return keys
}

// Put associate the specified value with the specified key in this dictionary.
func (m *ConcurrentHashDictionary[K, V]) Put(k K, v V) V {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	old := (*m.data)[k]
	(*m.data)[k] = v

	return old
}

// Remove removes the key-value pair with the specified key.
func (m *ConcurrentHashDictionary[K, V]) Remove(k K) V {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	old := (*m.data)[k]
	delete((*m.data), k)

	return old
}

// Replace replaces the value for the specified key only if it is currently in this dictionary.
func (m *ConcurrentHashDictionary[K, V]) Replace(k K, v V) (V, bool) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	old, ok := (*m.data)[k]
	if !ok {
		return old, false // zero value
	}

	(*m.data)[k] = v

	return old, true
}

// Size returns the number of key-value pairs in this dictionary.
func (m *ConcurrentHashDictionary[K, V]) Size() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	return m.data.Size()
}

// Values returns a slice that contains all the values in this dictionary.
func (m *ConcurrentHashDictionary[K, V]) Values() []V {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	arr := make([]V, 0, len(*m.data))

	for _, v := range *m.data {
		arr = append(arr, v)
	}

	return arr
}
