//go:build go1.23

package dict

import "iter"

// Iter returns an iterator of all elements in this dictionary.
func (m *ConcurrentHashDictionary[K, V]) Iter() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		m.mutex.RLock()
		defer m.mutex.RUnlock()

		for k, v := range *m.data {
			if !yield(k, v) {
				break
			}
		}
	}
}
