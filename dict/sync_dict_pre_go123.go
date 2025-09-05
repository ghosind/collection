//go:build !go1.23

package dict

func (m *SyncDict[K, V]) KeysIter() chan<- K {
	ch := make(chan K)
	go func() {
		read := m.loadPresentReadOnly()

		for k, e := range read.M {
			_, ok := e.Load(d.zero)
			if !ok {
				continue
			}
			ch <- k
		}
		close(ch)
	}()
	return ch
}

func (m *SyncDict[K, V]) ValuesIter() chan<- V {
	ch := make(chan V)
	go func() {
		read := m.loadPresentReadOnly()

		for _, e := range read.M {
			v, ok := e.Load(m.zero)
			if !ok {
				continue
			}
			ch <- v
		}
		close(ch)
	}()
	return ch
}
