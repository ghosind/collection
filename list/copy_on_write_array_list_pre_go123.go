//go:build !go1.23

package list

// Iter returns a channel of all elements in this collection.
func (l *CopyOnWriteArrayList[T]) Iter() <-chan T {
	data := l.data
	ch := make(chan T)

	go func() {
		for _, e := range data {
			ch <- e
		}
		close(ch)
	}()

	return ch
}
