package list

import (
	"github.com/ghosind/collection"
	"github.com/ghosind/collection/internal"
)

type ArrayList[T any] []T

// NewArrayList creates and returns a new empty list.
func NewArrayList[T any]() *ArrayList[T] {
	l := make(ArrayList[T], 0)

	return &l
}

// Add adds the specified element to the end of this list.
func (l *ArrayList[T]) Add(e T) bool {
	*l = append(*l, e)

	return true
}

// AddAll adds all of the elements to the end of this list.
func (l *ArrayList[T]) AddAll(c ...T) bool {
	*l = append(*l, c...)

	return true
}

// AddAtIndex inserts the specified element at the specified position in this list.
func (l *ArrayList[T]) AddAtIndex(i int, e T) {
	internal.CheckIndex(i, l.Size()+1)

	if i == l.Size() {
		*l = append(*l, e)
		return
	}

	*l = append(*l, e)
	copy((*l)[i+1:], (*l)[i:])
	(*l)[i] = e
}

// Clear removes all of the elements from this list.
func (l *ArrayList[T]) Clear() {
	*l = (*l)[:0]
}

// Clone returns a copy of this list.
func (l *ArrayList[T]) Clone() collection.List[T] {
	clone := make(ArrayList[T], 0, len(*l))
	clone.AddAll(*l...)
	return &clone
}

// Contains returns true if this list contains the specified element.
func (l *ArrayList[T]) Contains(e T) bool {
	return l.IndexOf(e) >= 0
}

// ContainsAll returns true if this list contains all of the elements.
func (l *ArrayList[T]) ContainsAll(c ...T) bool {
	for _, e := range c {
		if !l.Contains(e) {
			return false
		}
	}

	return true
}

// Equals returns true if this list is equal to the specified list.
func (l *ArrayList[T]) Equals(o any) bool {
	ol, ok := o.(*ArrayList[T])
	if !ok {
		return false
	}

	if l.Size() != ol.Size() {
		return false
	}

	for i, v := range *l {
		if !internal.Equal(v, (*ol)[i]) {
			return false
		}
	}

	return true
}

// ForEach performs the given handler for each element in this list until all elements have been
// processed or the handler returns an error.
func (l *ArrayList[T]) ForEach(handler func(e T) error) error {
	for _, v := range *l {
		if err := handler(v); err != nil {
			return err
		}
	}

	return nil
}

// Get returns the element at the specified position in this list.
func (l *ArrayList[T]) Get(i int) T {
	internal.CheckIndex(i, l.Size())
	return (*l)[i]
}

// IndexOf returns the index of the first occurrence of the specified element in this list,
// or -1 if this list does not contain the element.
func (l *ArrayList[T]) IndexOf(e T) int {
	for i, v := range *l {
		if internal.Equal(v, e) {
			return i
		}
	}

	return -1
}

// IsEmpty returns true if this list contains no elements.
func (l *ArrayList[T]) IsEmpty() bool {
	return l.Size() == 0
}

// LastIndexOf returns the index of the last occurrence of the specified element in this list,
// or -1 if this list does not contain the element.
func (l *ArrayList[T]) LastIndexOf(e T) int {
	for i := len(*l) - 1; i >= 0; i-- {
		if internal.Equal((*l)[i], e) {
			return i
		}
	}

	return -1
}

// Remove removes all occurrences of the specified element from this list, if it is present.
// Returns true if this list contained the specified element.
func (l *ArrayList[T]) Remove(e T) bool {
	i := l.IndexOf(e)
	if i == -1 {
		return false
	}

	for j := i; j < l.Size(); j++ {
		if !internal.Equal(e, (*l)[j]) {
			(*l)[i] = (*l)[j]
			i++
		}
	}

	*l = (*l)[:i]
	return true
}

// RemoveAll removes all occurrences of the specified elements from this list, if they are present.
// Returns true if this list contained any of the specified elements.
func (l *ArrayList[T]) RemoveAll(c ...T) bool {
	if len(c) == 0 {
		return false
	}

	found := false

	i := 0
	for j := 0; j < l.Size(); j++ {
		shouldRemove := false
		for _, e := range c {
			if internal.Equal(e, (*l)[j]) {
				shouldRemove = true
				found = true
				break
			}
		}

		if !shouldRemove {
			(*l)[i] = (*l)[j]
			i++
		}
	}

	*l = (*l)[:i]

	return found
}

// RemoveAtIndex removes the element at the specified position in this list.
// Returns the element that was removed from the list.
func (l *ArrayList[T]) RemoveAtIndex(i int) T {
	internal.CheckIndex(i, l.Size())

	old := (*l)[i]
	*l = append((*l)[:i], (*l)[i+1:]...)

	return old
}

// RemoveIf removes all of the elements of this list that satisfy the given predicate.
// Returns true if any elements were removed.
func (l *ArrayList[T]) RemoveIf(f func(T) bool) bool {
	if len(*l) == 0 {
		return false
	}

	found := false
	i := 0

	for j := 0; j < l.Size(); j++ {
		if !f((*l)[j]) {
			(*l)[i] = (*l)[j]
			i++
		} else {
			found = true
		}
	}

	*l = (*l)[:i]

	return found
}

// RetainAll retains only the elements in this list that are contained in the specified elements.
// In other words, removes from this list all of its elements that are not contained in the
// specified elements. Returns true if this list changed as a result of the call.
func (l *ArrayList[T]) RetainAll(c ...T) bool {
	if len(c) == 0 {
		return clearListForRetainAll[T](l)
	}

	found := false

	i := 0
	for j := 0; j < l.Size(); j++ {
		shouldRetain := false
		for _, e := range c {
			if internal.Equal(e, (*l)[j]) {
				shouldRetain = true
				break
			}
		}

		if shouldRetain {
			(*l)[i] = (*l)[j]
			i++
		} else {
			found = true
		}
	}

	*l = (*l)[:i]

	return found
}

// Set replaces the element at the specified position in this list with the specified element.
// Returns the element previously at the specified position. If the index is equal to the size of
// this list, the element is appended to the end of this list and a zero value is returned.
func (l *ArrayList[T]) Set(i int, e T) T {
	internal.CheckIndex(i, l.Size()+1)

	if i == l.Size() {
		*l = append(*l, e)
		var zero T
		return zero
	}

	old := (*l)[i]
	(*l)[i] = e

	return old
}

// Size returns the number of elements in this list.
func (l *ArrayList[T]) Size() int {
	return len(*l)
}

// ToSlice returns a slice containing all of the elements in this list in proper sequence.
func (l *ArrayList[T]) ToSlice() []T {
	arr := make([]T, len(*l))
	copy(arr, *l)

	return arr
}
