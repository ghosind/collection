//go:build go1.23

package dict

import "iter"

// Iter returns an iterator of all elements in this dictionary.
func (m *HashDict[K, V]) Iter() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range *m {
			if !yield(k, v) {
				break
			}
		}
	}
}
