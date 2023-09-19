package set

import (
	"github.com/ghosind/collection"
)

// HashSet is a set implementation that uses a Golang builtin map to store its elements.
type HashSet[T comparable] map[T]struct{}

// NewHashSet creates a new HashSet.
func NewHashSet[T comparable]() *HashSet[T] {
	set := make(HashSet[T])

	return &set
}

// Add adds the specified element to this set.
func (set *HashSet[T]) Add(e T) bool {
	_, found := (*set)[e]
	if found {
		return false
	}

	(*set)[e] = struct{}{}
	return true
}

// AddAll adds all of the specified elements to this set.
func (set *HashSet[T]) AddAll(c ...T) bool {
	isChanged := false

	for _, e := range c {
		_, found := (*set)[e]
		if !found {
			(*set)[e] = struct{}{}
			isChanged = true
		}
	}

	return isChanged
}

// Clear removes all of the elements from this set.
func (set *HashSet[T]) Clear() {
	*set = make(map[T]struct{})
}

// Clone returns a copy of this set.
func (set *HashSet[T]) Clone() collection.Set[T] {
	newSet := new(HashSet[T])
	*newSet = make(map[T]struct{}, set.Size())

	for e := range *set {
		(*newSet)[e] = struct{}{}
	}

	return newSet
}

// Contains returns true if this set contains the specified element.
func (set *HashSet[T]) Contains(e T) bool {
	_, found := (*set)[e]

	return found
}

// ContainsAll returns true if this set contains all of the specified elements.
func (set *HashSet[T]) ContainsAll(c ...T) bool {
	for _, e := range c {
		_, found := (*set)[e]
		if !found {
			return false
		}
	}

	return true
}

// Equals compares set with the object pass from parameter.
func (set *HashSet[T]) Equals(o any) bool {
	s, ok := o.(*HashSet[T])
	if !ok {
		return false
	}

	if s.Size() != set.Size() {
		return false
	}

	for k := range *set {
		_, ok := (*s)[k]
		if !ok {
			return false
		}
	}

	return true
}

// ForEach performs the given handler for each elements in the set until all elements have been processed or the handler returns an error.
func (set *HashSet[T]) ForEach(handler func(e T) error) error {
	for e := range *set {
		if err := handler(e); err != nil {
			return err
		}
	}

	return nil
}

// IsEmpty returns true if this set contains no elements.
func (set *HashSet[T]) IsEmpty() bool {
	return set.Size() == 0
}

// Iter returns a channel of all elements in this set.
func (set *HashSet[T]) Iter() <-chan T {
	ch := make(chan T)

	go func() {
		for e := range *set {
			ch <- e
		}

		close(ch)
	}()

	return ch
}

// Remove removes the specified element from this set.
func (set *HashSet[T]) Remove(e T) bool {
	_, found := (*set)[e]
	if !found {
		return false
	}

	delete(*set, e)
	return true
}

// RemoveAll removes all of the specified elements from this set.
func (set *HashSet[T]) RemoveAll(c ...T) bool {
	isChanged := false

	for _, e := range c {
		_, found := (*set)[e]
		if found {
			isChanged = true
			delete(*set, e)
		}
	}

	return isChanged
}

// Size returns the number of elements in this set.
func (set *HashSet[T]) Size() int {
	return len(*set)
}

// ToSlice returns a slice containing all of the elements in this set.
func (set *HashSet[T]) ToSlice() []T {
	slice := make([]T, 0, set.Size())

	for e := range *set {
		slice = append(slice, e)
	}

	return slice
}
