//go:build go1.23

package dict

import "iter"

// Iter returns an iterator of all elements in this dictionary.
func (m *LockDict[K, V]) Iter() iter.Seq2[K, V] {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.data.Iter()
}

// KeysIter returns an iterator of all keys in this dictionary.
func (m *LockDict[K, V]) KeysIter() iter.Seq[K] {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.data.KeysIter()
}

// ValuesIter returns an iterator of all values in this dictionary.
func (m *LockDict[K, V]) ValuesIter() iter.Seq[V] {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.data.ValuesIter()
}
