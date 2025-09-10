//go:build go1.23

package list

import "iter"

// Iter returns an iterator over the elements in this list in proper sequence.
func (l *ArrayList[T]) Iter() iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, e := range *l {
			if !yield(e) {
				break
			}
		}
	}
}
