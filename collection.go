package collection

// Collection is the root interface for this collections framework hierarchy.
type Collection[T any] interface {
	// Add adds the specified element to this collection.
	Add(e T) bool

	// AddAll adds all of the elements in the this collection.
	AddAll(c ...T) bool

	// Clear removes all of the elements from this collection.
	Clear()

	// Contains returns true if this collection contains the specified element.
	Contains(e T) bool

	// ContainsAll returns true if this collection contains all of the elements in the specified collection.
	ContainsAll(c ...T) bool

	// ForEach performs the given handler for each elements in the collection until all elements have been processed or the handler returns an error.
	ForEach(func(e T) error) error

	// IsEmpty returns true if this collection contains no elements.
	IsEmpty() bool

	// Remove removes the specified element from this collection.
	Remove(e T) bool

	// RemoveAll removes all of the elements in the specified collection from this collection.
	RemoveAll(c ...T) bool

	// Size returns the number of elements in this collection.
	Size() int

	// ToSlice returns a slice containing all of the elements in this collection.
	ToSlice() []T
}
