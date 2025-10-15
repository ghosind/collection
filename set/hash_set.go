package set

import (
	"bytes"
	"fmt"

	"github.com/ghosind/collection"
)

// HashSet is a set implementation that uses a Golang builtin map to store its elements.
type HashSet[T comparable] map[T]empty

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

	(*set)[e] = empty{}
	return true
}

// AddAll adds all of the specified elements to this set.
func (set *HashSet[T]) AddAll(c ...T) bool {
	isChanged := false

	for _, e := range c {
		_, found := (*set)[e]
		if !found {
			(*set)[e] = empty{}
			isChanged = true
		}
	}

	return isChanged
}

// Clone returns a copy of this set.
func (set *HashSet[T]) Clone() collection.Set[T] {
	newSet := new(HashSet[T])
	*newSet = make(map[T]empty, set.Size())

	for e := range *set {
		(*newSet)[e] = empty{}
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

// ForEach performs the given handler for each elements in the set until all elements have been
// processed or the handler returns an error.
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

// RemoveIf removes all of the elements of this set that satisfy the given predicate.
func (set *HashSet[T]) RemoveIf(filter func(T) bool) bool {
	isChanged := false

	for e := range *set {
		if filter(e) {
			delete(*set, e)
			isChanged = true
		}
	}

	return isChanged
}

// RetainAll retains only the elements in this set that are contained in the specified collection.
func (set *HashSet[T]) RetainAll(c ...T) bool {
	cSet := NewHashSet[T]()
	cSet.AddAll(c...)
	isChanged := false

	for e := range *set {
		if !cSet.Contains(e) {
			delete(*set, e)
			isChanged = true
		}
	}

	return isChanged
}

// Size returns the number of elements in this set.
func (set *HashSet[T]) Size() int {
	return len(*set)
}

// String returns the string representation of this set.
func (set *HashSet[T]) String() string {
	buf := bytes.NewBufferString("set[")
	first := true
	for e := range *set {
		if !first {
			buf.WriteString(" ")
		}
		first = false
		fmt.Fprintf(buf, "%v", e)
	}
	buf.WriteString("]")
	return buf.String()
}

// ToSlice returns a slice containing all of the elements in this set.
func (set *HashSet[T]) ToSlice() []T {
	slice := make([]T, 0, set.Size())

	for e := range *set {
		slice = append(slice, e)
	}

	return slice
}
