//go:build go1.21

package set

// Clear removes all of the elements from this set.
func (set *HashSet[T]) Clear() {
	clear(*set)
}
