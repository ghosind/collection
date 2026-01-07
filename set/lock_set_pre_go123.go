//go:build !go1.23

package set

// Iter returns a channel of all elements in this set.
func (s *LockSet[T]) Iter() <-chan T {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.data.Iter()
}
