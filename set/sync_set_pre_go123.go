//go:build !go1.23

package set

// Iter returns a channel of all elements in this set.
func (set *SyncSet[T]) Iter() <-chan T {
	read := set.loadPresentReadOnly()

	ch := make(chan T)
	go func() {
		for k, e := range read.M {
			_, ok := e.Load(emptyZero)
			if ok {
				ch <- k
			}
		}
		close(ch)
	}()

	return ch
}
