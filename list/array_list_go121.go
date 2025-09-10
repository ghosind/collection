//go:build go1.21

package list

// Clear removes all of the elements from this list.
func (l *ArrayList[T]) Clear() {
	clear(*l)
}
