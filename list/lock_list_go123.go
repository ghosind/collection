//go:build go1.23

package list

import "iter"

// Iter returns an iterator over the elements in this list in proper sequence.
func (l *LockList[T]) Iter() iter.Seq[T] {
	l.mu.RLock()
	defer l.mu.RUnlock()

	return l.data.Iter()
}
