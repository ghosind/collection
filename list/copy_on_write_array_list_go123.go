//go:build go1.23

package list

import "iter"

// Iter returns an iterator of all elements in this collection.
func (l *CopyOnWriteArrayList[T]) Iter() iter.Seq[T] {
	data := l.data

	return func(yield func(T) bool) {
		for _, e := range data {
			if !yield(e) {
				break
			}
		}
	}
}
