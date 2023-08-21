package collection

// Set is a collection interface that contains no duplicate elements.
type Set[T comparable] interface {
	Collection[T]

	// Clone returns a copy of this set.
	Clone() Set[T]

	// ForEach performs the given handler for each elements in the collection until all elements
	// have been processed or the handler returns an error.
	ForEach(func(e T) error) error
}
