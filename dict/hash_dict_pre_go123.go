//go:build !go1.23

package dict

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
