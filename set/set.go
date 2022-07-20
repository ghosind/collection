package set

import "github.com/ghosind/collection"

// Set is a collection interface that contains no duplicate elements.
type Set[T comparable] interface {
	collection.Collection[T]
}
