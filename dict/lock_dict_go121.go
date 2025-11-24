//go:build go1.21

package dict

// Clear removes all key-value pairs in this dictionary.
func (m *LockDict[K, V]) Clear() {
	m.mu.Lock()
	defer m.mu.Unlock()

	clear(m.data)
}
