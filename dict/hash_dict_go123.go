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

// KeysIter returns an iterator of all keys in this dictionary.
func (m *HashDict[K, V]) KeysIter() iter.Seq[K] {
	return func(yield func(K) bool) {
		for k := range *m {
			if !yield(k) {
				break
			}
		}
	}
}

// ValuesIter returns an iterator of all values in this dictionary.
func (m *HashDict[K, V]) ValuesIter() iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range *m {
			if !yield(v) {
				break
			}
		}
	}
}
