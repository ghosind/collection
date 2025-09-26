package list

import "github.com/ghosind/collection"

// clearListForRetainAll clears the list and returns true if the list was not empty.
// This is specifically designed for the case when RetainAll is called with no elements (empty collection),
// which should remove all elements from the list. It is non thread-safe.
func clearListForRetainAll[T any](l collection.List[T]) bool {
	if l.IsEmpty() {
		return false
	}

	l.Clear()
	return true
}
