//go:build go1.23

package list

import "iter"

// Iter returns a channel of all elements in this collection.
func (l *LinkedList[T]) Iter() iter.Seq[T] {
	return func(yield func(T) bool) {
		for node := l.head; node != nil; node = node.Next {
			if !yield(node.Value) {
				break
			}
		}
	}
}
