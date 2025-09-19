package collection

// List is an ordered collection.
type List[T any] interface {
	Collection[T]

	// AddAtIndex inserts the specified element to the specified position in this list.
	AddAtIndex(i int, e T)

	// Clone returns a copy of this list.
	Clone() List[T]

	// Get returns the element at the specified position inn this list.
	Get(i int) T

	// IndexOf returns the index of the first occurrence of the specified element in this list, or -1
	// if this list does not contain the element.
	IndexOf(e T) int

	// LastIndexOf returns the index of the last occurrence of the specified element in this list, or
	// -1 if this list does not contain the element.
	LastIndexOf(e T) int

	// RemoveAtIndex removes the element at the specified position in this list.
	RemoveAtIndex(i int) T

	// Set replaces the element at the specified position in this list with the specified element.
	Set(i int, e T) T
}
