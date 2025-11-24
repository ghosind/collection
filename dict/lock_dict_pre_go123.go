//go:build !go1.23

package dict

// KeysIter returns a channel iterator of all keys in this dictionary.
func (m *LockDict[K, V]) KeysIter() chan<- K {
	m.mu.RLock()
	defer m.mu.RUnlock()

	ch := make(chan K)
	go func() {
		for k := range m.data {
			ch <- k
		}
		close(ch)
	}()

	return ch
}

// ValuesIter returns a channel iterator of all values in this dictionary.
func (m *LockDict[K, V]) ValuesIter() chan<- V {
	m.mu.RLock()
	defer m.mu.RUnlock()

	ch := make(chan V)
	go func() {
		for _, v := range m.data {
			ch <- v
		}
		close(ch)
	}()

	return ch
}
