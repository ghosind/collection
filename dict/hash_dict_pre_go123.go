//go:build !go1.23

package dict

// KeysIter returns a channel iterator of all keys in this dictionary.
func (m *HashDict[K, V]) KeysIter() chan<- K {
	ch := make(chan K)
	go func() {
		for k := range *m {
			ch <- k
		}
		close(ch)
	}()
	return ch
}

// ValuesIter returns a channel iterator of all values in this dictionary.
func (m *HashDict[K, V]) ValuesIter() chan<- V {
	ch := make(chan V)
	go func() {
		for _, v := range *m {
			ch <- v
		}
		close(ch)
	}()
	return ch
}
