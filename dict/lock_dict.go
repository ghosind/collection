package dict

import (
	"bytes"
	"encoding/json"
	"sync"

	"github.com/ghosind/collection"
	"github.com/ghosind/collection/internal"
)

// LockDict is a thread-safe dictionary using a read-write mutex.
type LockDict[K comparable, V any] struct {
	data map[K]V
	mu   sync.RWMutex
}

// NewLockDict creates a new LockDict.
func NewLockDict[K comparable, V any]() *LockDict[K, V] {
	d := new(LockDict[K, V])
	d.data = make(map[K]V)

	return d
}

// NewLockDictFrom creates a new LockDict from the given map.
func NewLockDictFrom[K comparable, V any](m map[K]V) *LockDict[K, V] {
	d := new(LockDict[K, V])
	d.data = make(map[K]V, len(m))

	for k, v := range m {
		d.data[k] = v
	}

	return d
}

// Clone returns a copy of this dictionary.
func (m *LockDict[K, V]) Clone() collection.Dict[K, V] {
	m.mu.RLock()
	defer m.mu.RUnlock()

	newDict := new(LockDict[K, V])
	newDict.data = make(map[K]V, len(m.data))

	for k, v := range m.data {
		newDict.data[k] = v
	}

	return newDict
}

// ContainsKey returns true if this dictionary contains a key-value pair with the specified key.
func (m *LockDict[K, V]) ContainsKey(k K) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()

	_, ok := m.data[k]

	return ok
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

	if len(m.data) != len(om.data) {
		return false
	}

	for k, v := range m.data {
		val, ok := om.data[k]
		if !ok {
			return false
		}

		if !internal.Equal(v, val) {
			return false
		}
	}

	return true
}

// ForEach performs the given handler for each key-value pairs in the dictionary until all pairs
// have been processed or the handler returns an error.
func (m *LockDict[K, V]) ForEach(handler func(K, V) error) error {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for k, v := range m.data {
		if err := handler(k, v); err != nil {
			return err
		}
	}

	return nil
}

// Get returns the value which associated to the specified key.
func (m *LockDict[K, V]) Get(k K) (V, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	v, ok := m.data[k]
	return v, ok
}

// GetDefault returns the value associated with the specified key, and returns the default value if
// this dictionary contains no pair with the key.
func (m *LockDict[K, V]) GetDefault(k K, defaultVal V) V {
	m.mu.RLock()
	defer m.mu.RUnlock()

	v, ok := m.data[k]
	if !ok {
		return defaultVal
	}

	return v
}

// IsEmpty returns true if this dictionary is empty.
func (m *LockDict[K, V]) IsEmpty() bool {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return len(m.data) == 0
}

// Keys returns a slice that contains all the keys in this dictionary.
func (m *LockDict[K, V]) Keys() []K {
	m.mu.RLock()
	defer m.mu.RUnlock()

	keys := make([]K, 0, len(m.data))
	for k := range m.data {
		keys = append(keys, k)
	}

	return keys
}

// Put associate the specified value with the specified key in this dictionary.
func (m *LockDict[K, V]) Put(k K, v V) V {
	m.mu.Lock()
	defer m.mu.Unlock()

	old := m.data[k]
	m.data[k] = v

	return old
}

// Remove removes the key-value pair with the specified key.
func (m *LockDict[K, V]) Remove(k K) V {
	m.mu.Lock()
	defer m.mu.Unlock()

	old := m.data[k]
	delete(m.data, k)

	return old
}

// Replace replaces the value for the specified key only if it is currently in this dictionary.
func (m *LockDict[K, V]) Replace(k K, v V) (V, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	old, ok := m.data[k]
	if !ok {
		return old, false // zero value
	}

	m.data[k] = v

	return old, true
}

// Size returns the number of key-value pairs in this dictionary.
func (m *LockDict[K, V]) Size() int {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return len(m.data)
}

// String returns the string representation of this dictionary.
func (m *LockDict[K, V]) String() string {
	m.mu.RLock()
	defer m.mu.RUnlock()

	buf := bytes.NewBufferString("dict[")
	count := 0
	for k, v := range m.data {
		if count > 0 {
			buf.WriteString(" ")
		}
		buf.WriteString(internal.ValueString(k))
		buf.WriteString(": ")
		buf.WriteString(internal.ValueString(v))
		count++
	}
	buf.WriteString("]")
	return buf.String()
}

// Values returns a slice that contains all the values in this dictionary.
func (m *LockDict[K, V]) Values() []V {
	m.mu.RLock()
	defer m.mu.RUnlock()

	arr := make([]V, 0, len(m.data))

	for _, v := range m.data {
		arr = append(arr, v)
	}

	return arr
}

// MarshalJSON marshals the HashDict as a JSON object (map).
func (m *LockDict[K, V]) MarshalJSON() ([]byte, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return json.Marshal(m.data)
}

// UnmarshalJSON unmarshals a JSON object into the HashDict.
func (m *LockDict[K, V]) UnmarshalJSON(b []byte) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	var tmp map[K]V
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	m.data = tmp
	return nil
}
