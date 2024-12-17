//go:build go1.23

package dict

import "iter"

// Iter returns an iterator of all elements in this dictionary.
func (d *SyncDict[K, V]) Iter() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		read := d.loadPresentReadOnly()

		for k, e := range read.M {
			v, ok := e.Load(d.zero)
			if !ok {
				continue
			}
			if !yield(k, v) {
				break
			}
		}
	}
}
