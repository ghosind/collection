//go:build !go1.21

package dict

// Clear removes all key-value pairs in this dictionary.
func (m *HashDict[K, V]) Clear() {
	*m = make(HashDict[K, V])
}
