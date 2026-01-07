package set

import (
	"sync"

	"github.com/ghosind/collection"
)

// LockSet is a thread-safe set that wraps another set with read-write locks.
type LockSet[T comparable] struct {
	data collection.Set[T]
	mu   sync.RWMutex
}

// NewHashSet creates a new HashSet.
func NewLockSet[T comparable](data collection.Set[T]) *LockSet[T] {
	s := new(LockSet[T])
	s.data = data

	return s
}

// Add adds the specified element to this set.
func (s *LockSet[T]) Add(e T) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.data.Add(e)
}

// AddAll adds all of the specified elements to this set.
func (s *LockSet[T]) AddAll(c ...T) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.data.AddAll(c...)
}

// Clear removes all of the elements from this set.
func (s *LockSet[T]) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data.Clear()
}

// Clone returns a copy of this set.
func (s *LockSet[T]) Clone() collection.Set[T] {
	s.mu.RLock()
	defer s.mu.RUnlock()

	cloned := s.data.Clone()

	return NewLockSet[T](cloned)
}

// Contains returns true if this set contains the specified element.
func (s *LockSet[T]) Contains(e T) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.data.Contains(e)
}

// ContainsAll returns true if this set contains all of the specified elements.
func (s *LockSet[T]) ContainsAll(c ...T) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.data.ContainsAll(c...)
}

// Equals compares set with the object pass from parameter.
func (s *LockSet[T]) Equals(o any) bool {
	os, ok := o.(*LockSet[T])
	if !ok {
		return false
	}

	s.mu.RLock()
	defer s.mu.RUnlock()
	os.mu.RLock()
	defer os.mu.RUnlock()

	return s.data.Equals(os.data)
}

// ForEach performs the given handler for each elements in the set until all elements have been
// processed or the handler returns an error.
func (s *LockSet[T]) ForEach(handler func(e T) error) error {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.data.ForEach(handler)
}

// IsEmpty returns true if this set contains no elements.
func (s *LockSet[T]) IsEmpty() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.data.IsEmpty()
}

// Remove removes the specified element from this set.
func (s *LockSet[T]) Remove(e T) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.data.Remove(e)
}

// RemoveAll removes all of the specified elements from this set.
func (s *LockSet[T]) RemoveAll(c ...T) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.data.RemoveAll(c...)
}

// RemoveIf removes all of the elements of this set that satisfy the given predicate.
func (s *LockSet[T]) RemoveIf(filter func(T) bool) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.data.RemoveIf(filter)
}

// RetainAll retains only the elements in this set that are contained in the specified collection.
func (s *LockSet[T]) RetainAll(c ...T) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.data.RetainAll(c...)
}

// Size returns the number of elements in this set.
func (s *LockSet[T]) Size() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.data.Size()
}

// String returns the string representation of this set.
func (s *LockSet[T]) String() string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.data.String()
}

// ToSlice returns a slice containing all of the elements in this set.
func (s *LockSet[T]) ToSlice() []T {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.data.ToSlice()
}

// MarshalJSON marshals the set as a JSON array.
func (s *LockSet[T]) MarshalJSON() ([]byte, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.data.MarshalJSON()
}

// UnmarshalJSON unmarshals a JSON array into the set.
func (s *LockSet[T]) UnmarshalJSON(b []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.data.UnmarshalJSON(b)
}
