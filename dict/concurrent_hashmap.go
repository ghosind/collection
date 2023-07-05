package dict

import (
	"reflect"
	"sync"

	"github.com/ghosind/collection"
	"github.com/ghosind/utils"
)

// ConcurrentHashMap is the thread-safe map implementation.
type ConcurrentHashMap[K comparable, V any] struct {
	data *HashMap[K, V]

	mutex sync.RWMutex
}

// NewConcurrentHashMap creates a new ConcurrentHashMap.
func NewConcurrentHashMap[K comparable, V any]() *ConcurrentHashMap[K, V] {
	newMap := new(ConcurrentHashMap[K, V])
	newMap.data = NewHashMap[K, V]()

	return newMap
}

// Clear removes all key-value pairs in this map.
func (m *ConcurrentHashMap[K, V]) Clear() {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.data = NewHashMap[K, V]()
}

// Clone returns a copy of this map.
func (m *ConcurrentHashMap[K, V]) Clone() collection.Map[K, V] {
	newMap := NewConcurrentHashMap[K, V]()

	for k, v := range *m.data {
		(*newMap.data)[k] = v
	}

	return newMap
}

// ContainsKey returns true if this map contains a key-value pair with the specified key.
func (m *ConcurrentHashMap[K, V]) ContainsKey(k K) bool {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	_, ok := (*m.data)[k]

	return ok
}

// Equals compares this map with the object pass from parameter.
func (m *ConcurrentHashMap[K, V]) Equals(o any) bool {
	if !utils.IsSameType(m, o) {
		return false
	}

	om := o.(*ConcurrentHashMap[K, V])
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

// ForEach performs the given handler for each key-value pairs in the map until all pairs have
// been processed or the handler returns an error.
func (m *ConcurrentHashMap[K, V]) ForEach(handler func(K, V) error) error {
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
func (m *ConcurrentHashMap[K, V]) Get(k K) (V, bool) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	v, ok := (*m.data)[k]
	return v, ok
}

// GetDefault returns the value associated with the specified key, and returns the default value if
// this map contains no pair with the key.
func (m *ConcurrentHashMap[K, V]) GetDefault(k K, defaultVal V) V {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	v, ok := (*m.data)[k]
	if !ok {
		return defaultVal
	}

	return v
}

// IsEmpty returns true if this map is empty.
func (m *ConcurrentHashMap[K, V]) IsEmpty() bool {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	return m.data.Size() == 0
}

// Keys returns a slice that contains all the keys in this map.
func (m *ConcurrentHashMap[K, V]) Keys() []K {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	keys := make([]K, 0, len(*m.data))
	for k := range *m.data {
		keys = append(keys, k)
	}

	return keys
}

// Put associate the specified value with the specified key in this map.
func (m *ConcurrentHashMap[K, V]) Put(k K, v V) V {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	old := (*m.data)[k]
	(*m.data)[k] = v

	return old
}

// Remove removes the key-value pair with the specified key.
func (m *ConcurrentHashMap[K, V]) Remove(k K) V {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	old := (*m.data)[k]
	delete((*m.data), k)

	return old
}

// Replace replaces the value for the specified key only if it is currently in this map.
func (m *ConcurrentHashMap[K, V]) Replace(k K, v V) (V, bool) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	old, ok := (*m.data)[k]
	if !ok {
		return old, false // zero value
	}

	(*m.data)[k] = v

	return old, true
}

// Size returns the number of key-value pairs in this map.
func (m *ConcurrentHashMap[K, V]) Size() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	return m.data.Size()
}

// Values returns a slice that contains all the values in this map.
func (m *ConcurrentHashMap[K, V]) Values() []V {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	arr := make([]V, 0, len(*m.data))

	for _, v := range *m.data {
		arr = append(arr, v)
	}

	return arr
}
