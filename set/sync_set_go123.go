//go:build go1.23

package set

import "iter"

// Iter returns a channel of all elements in this set.
func (set *SyncSet[T]) Iter() iter.Seq[T] {
	read := set.loadPresentReadOnly()

	return func(yield func(T) bool) {
		for k, e := range read.M {
			_, ok := e.Load(emptyZero)
			if ok {
				if !yield(k) {
					break
				}
			}
		}
	}
}
