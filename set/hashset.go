package set

// HashSet is a set implementation that uses a Golang builtin map to store its elements.
type HashSet[T comparable] struct {
	Set[T]

	data map[T]struct{}
}

// NewHashSet creates a new HashSet.
func NewHashSet[T comparable]() *HashSet[T] {
	set := new(HashSet[T])

	set.data = make(map[T]struct{})

	return set
}

// Add adds the specified element to this set.
func (set *HashSet[T]) Add(e T) bool {
	_, found := set.data[e]
	if found {
		return false
	}

	set.data[e] = struct{}{}
	return true
}

// AddAll adds all of the specified elements to this set.
func (set *HashSet[T]) AddAll(c ...T) bool {
	isChanged := false

	for _, e := range c {
		_, found := set.data[e]
		if !found {
			set.data[e] = struct{}{}
			isChanged = true
		}
	}

	return isChanged
}

// Clear removes all of the elements from this set.
func (set *HashSet[T]) Clear() {
	set.data = make(map[T]struct{})
}

// Contains returns true if this set contains the specified element.
func (set *HashSet[T]) Contains(e T) bool {
	_, found := set.data[e]

	return found
}

// ContainsAll returns true if this set contains all of the specified elements.
func (set *HashSet[T]) ContainsAll(c ...T) bool {
	for _, e := range c {
		_, found := set.data[e]
		if !found {
			return false
		}
	}

	return true
}

// IsEmpty returns true if this set contains no elements.
func (set *HashSet[T]) IsEmpty() bool {
	return set.Size() == 0
}

// Remove removes the specified element from this set.
func (set *HashSet[T]) Remove(e T) bool {
	_, found := set.data[e]
	if !found {
		return false
	}

	delete(set.data, e)
	return true
}

// RemoveAll removes all of the specified elements from this set.
func (set *HashSet[T]) RemoveAll(c ...T) bool {
	isChanged := false

	for _, e := range c {
		_, found := set.data[e]
		if found {
			isChanged = true
			delete(set.data, e)
		}
	}

	return isChanged
}

// Size returns the number of elements in this set.
func (set *HashSet[T]) Size() int {
	return len(set.data)
}

// ToSlice returns a slice containing all of the elements in this set.
func (set *HashSet[T]) ToSlice() []T {
	slice := make([]T, 0, set.Size())

	for e := range set.data {
		slice = append(slice, e)
	}

	return slice
}
