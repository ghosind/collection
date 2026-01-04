package collection

// SequencedCollection is a collection that maintains the order of elements.
type SequencedCollection[T any] interface {
	Collection[T]

	// AddAtIndex inserts the specified element to the specified position in this list.
	AddAtIndex(i int, e T)

	// Get returns the element at the specified position in this list.
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

// List is an ordered collection.
type List[T any] interface {
	SequencedCollection[T]

	// Clone returns a copy of this list.
	Clone() List[T]
}

// Stack is a collection that follows the LIFO (last-in, first-out) principle.
type Stack[T any] interface {
	SequencedCollection[T]

	// Clone returns a shallow copy of this stack.
	Clone() Stack[T]

	// Peek returns the element at the top of this stack without removing it.
	Peek() T

	// Pop removes and returns the element at the top of this stack.
	Pop() T

	// Push adds the specified element to the top of this stack.
	Push(e T)
}
