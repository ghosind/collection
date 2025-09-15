package collection

// Set is a collection interface that contains no duplicate elements.
type Set[T comparable] interface {
	Collection[T]

	// Clone returns a copy of this set.
	Clone() Set[T]
}
