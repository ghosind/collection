//go:build go1.23

package set

import "iter"

// Iter returns a channel of all elements in this set.
func (s *LockSet[T]) Iter() iter.Seq[T] {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.data.Iter()
}
