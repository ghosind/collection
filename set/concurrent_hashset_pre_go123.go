//go:build !go1.23

package set

// Iter returns a channel of all elements in this set.
func (set *ConcurrentHashSet[T]) Iter() <-chan T {
	ch := make(chan T)

	go func() {
		set.mutex.RLock()
		defer set.mutex.RUnlock()

		for e := range *set.data {
			ch <- e
		}

		close(ch)
	}()

	return ch
}
