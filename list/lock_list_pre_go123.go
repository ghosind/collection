//go:build !go1.23

package list

// Iter returns a channel that can be used to iterate over the elements in this list in proper
// sequence.
func (l *LockList[T]) Iter() <-chan T {
	l.mu.RLock()
	defer l.mu.RUnlock()

	return l.data.Iter()
}
