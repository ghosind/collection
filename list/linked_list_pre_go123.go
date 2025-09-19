//go:build !go1.23

package list

// Iter returns a channel of all elements in this collection.
func (l *LinkedList[T]) Iter() <-chan T {
	ch := make(chan T)
	go func() {
		for node := l.head; node != nil; node = node.Next {
			ch <- node.Value
		}
		close(ch)
	}()
	return ch
}
