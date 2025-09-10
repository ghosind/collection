//go:build !go1.23

package list

// Iter returns a channel that can be used to iterate over the elements in this list in proper
// sequence.
func (l *ArrayList[T]) Iter() <-chan T {
	ch := make(chan T)

	go func() {
		for _, e := range *l {
			ch <- e
		}
		close(ch)
	}()

	return ch
}
