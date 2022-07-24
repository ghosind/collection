package collection

// Set is a collection interface that contains no duplicate elements.
type Set[T comparable] interface {
	Collection[T]
}
