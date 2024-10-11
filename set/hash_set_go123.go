//go:build go1.23

package set

import "iter"

// Iter returns a channel of all elements in this set.
func (set *HashSet[T]) Iter() iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := range *set {
			if !yield(e) {
				break
			}
		}
	}
}
