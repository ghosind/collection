//go:build go1.23

package dict

import "iter"

// Iter returns an iterator of all elements in this dictionary.
func (m *LockDict[K, V]) Iter() iter.Seq2[K, V] {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return func(yield func(K, V) bool) {
		for k, v := range m.data {
			if !yield(k, v) {
				break
			}
		}
	}
}

// KeysIter returns an iterator of all keys in this dictionary.
func (m *LockDict[K, V]) KeysIter() iter.Seq[K] {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return func(yield func(K) bool) {
		for k := range m.data {
			if !yield(k) {
				break
			}
		}
	}
}

// ValuesIter returns an iterator of all values in this dictionary.
func (m *LockDict[K, V]) ValuesIter() iter.Seq[V] {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return func(yield func(V) bool) {
		for _, v := range m.data {
			if !yield(v) {
				break
			}
		}
	}
}
