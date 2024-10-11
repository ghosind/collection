//go:build !go1.23

package set

// Iter returns a channel of all elements in this set.
func (set *HashSet[T]) Iter() <-chan T {
	ch := make(chan T)

	go func() {
		for e := range *set {
			ch <- e
		}

		close(ch)
	}()

	return ch
}
