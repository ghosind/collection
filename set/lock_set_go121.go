//go:build go1.21

package set

// Clear removes all of the elements from this set.
func (s *LockSet[T]) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()

	clear(s.data)
}
