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

	// RemoveFirst removes the first occurrence of the specified element from this list, if it is present.
	// Returns true if the element was removed.
	RemoveFirst(e T) bool

	// RemoveFirstN removes the first n occurrences of the specified element from this list.
	// Returns the number of elements removed.
	RemoveFirstN(e T, n int) int

	// RemoveLast removes the last occurrence of the specified element from this list, if it is present.
	// Returns true if the element was removed.
	RemoveLast(e T) bool

	// RemoveLastN removes the last n occurrences of the specified element from this list.
	// Returns the number of elements removed.
	RemoveLastN(e T, n int) int

	// Set replaces the element at the specified position in this list with the specified element.
	Set(i int, e T) T

	// Trim removes the first n elements from this list. Returns the number of elements removed.
	Trim(n int) int

	// TrimLast removes the last n elements from this list. Returns the number of elements removed.
	TrimLast(n int) int
}

// List is an ordered collection.
type List[T any] interface {
	SequencedCollection[T]

	// Clone returns a copy of this list.
	Clone() List[T]

	// SubList returns a view of the portion of this list between the specified fromIndex, inclusive,
	// and toIndex, exclusive.
	SubList(fromIndex, toIndex int) List[T]
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
