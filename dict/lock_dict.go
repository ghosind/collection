package dict

import (
	"sync"

	"github.com/ghosind/collection"
)

// LockDict is a thread-safe dictionary that wraps another dictionary with read-write locks.
type LockDict[K comparable, V any] struct {
	data collection.Dict[K, V]
	mu   sync.RWMutex
}

// NewLockDict creates a new LockDict.
func NewLockDict[K comparable, V any](data collection.Dict[K, V]) *LockDict[K, V] {
	d := new(LockDict[K, V])
	d.data = data

	return d
}

// Clear removes all key-value pairs in this dictionary.
func (m *LockDict[K, V]) Clear() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.data.Clear()
}

// Clone returns a copy of this dictionary.
func (m *LockDict[K, V]) Clone() collection.Dict[K, V] {
	m.mu.RLock()
	defer m.mu.RUnlock()

	cloned := m.data.Clone()

	return NewLockDict[K, V](cloned)
}

// ContainsKey returns true if this dictionary contains a key-value pair with the specified key.
func (m *LockDict[K, V]) ContainsKey(k K) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.data.ContainsKey(k)
}

// Equals compares this dictionary with the object pass from parameter.
func (m *LockDict[K, V]) Equals(o any) bool {
	om, ok := o.(*LockDict[K, V])
	if !ok {
		return false
	}

	m.mu.RLock()
	defer m.mu.RUnlock()

	om.mu.RLock()
	defer om.mu.RUnlock()

	return m.data.Equals(om.data)
}

// ForEach performs the given handler for each key-value pairs in the dictionary until all pairs
// have been processed or the handler returns an error.
func (m *LockDict[K, V]) ForEach(handler func(K, V) error) error {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.data.ForEach(handler)
}

// Get returns the value which associated to the specified key.
func (m *LockDict[K, V]) Get(k K) (V, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.data.Get(k)
}

// GetDefault returns the value associated with the specified key, and returns the default value if
// this dictionary contains no pair with the key.
func (m *LockDict[K, V]) GetDefault(k K, defaultVal V) V {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.data.GetDefault(k, defaultVal)
}

// IsEmpty returns true if this dictionary is empty.
func (m *LockDict[K, V]) IsEmpty() bool {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.data.IsEmpty()
}

// Keys returns a slice that contains all the keys in this dictionary.
func (m *LockDict[K, V]) Keys() []K {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.data.Keys()
}

// Put associate the specified value with the specified key in this dictionary.
func (m *LockDict[K, V]) Put(k K, v V) V {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.data.Put(k, v)
}

// Remove removes the key-value pair with the specified key.
func (m *LockDict[K, V]) Remove(k K) V {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.data.Remove(k)
}

// Replace replaces the value for the specified key only if it is currently in this dictionary.
func (m *LockDict[K, V]) Replace(k K, v V) (V, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.data.Replace(k, v)
}

// Size returns the number of key-value pairs in this dictionary.
func (m *LockDict[K, V]) Size() int {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.data.Size()
}

// String returns the string representation of this dictionary.
func (m *LockDict[K, V]) String() string {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.data.String()
}

// Values returns a slice that contains all the values in this dictionary.
func (m *LockDict[K, V]) Values() []V {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.data.Values()
}

// MarshalJSON marshals the HashDict as a JSON object (map).
func (m *LockDict[K, V]) MarshalJSON() ([]byte, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.data.MarshalJSON()
}

// UnmarshalJSON unmarshals a JSON object into the HashDict.
func (m *LockDict[K, V]) UnmarshalJSON(b []byte) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.data.UnmarshalJSON(b)
}
