package set

import (
	"bytes"
	"encoding/json"
	"sync"

	"github.com/ghosind/collection"
	"github.com/ghosind/collection/internal"
)

// LockSet is a thread-safe set implementation that uses a Golang builtin map to store its
// elements.
type LockSet[T comparable] struct {
	data map[T]empty
	mu   sync.RWMutex
}

// NewHashSet creates a new HashSet.
func NewLockSet[T comparable]() *LockSet[T] {
	s := new(LockSet[T])
	s.data = make(map[T]empty)

	return s
}

// NewHashSetFrom creates and returns a new HashSet containing the elements of the
// provided collection.
func NewLockSetFrom[T comparable](c ...T) *LockSet[T] {
	s := new(LockSet[T])
	s.data = make(map[T]empty, len(c))
	for _, e := range c {
		s.data[e] = empty{}
	}

	return s
}

// Add adds the specified element to this set.
func (s *LockSet[T]) Add(e T) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, found := s.data[e]
	if found {
		return false
	}

	s.data[e] = empty{}
	return true
}

// AddAll adds all of the specified elements to this set.
func (s *LockSet[T]) AddAll(c ...T) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	isChanged := false

	for _, e := range c {
		_, found := s.data[e]
		if !found {
			s.data[e] = empty{}
			isChanged = true
		}
	}

	return isChanged
}

// Clone returns a copy of this set.
func (s *LockSet[T]) Clone() collection.Set[T] {
	s.mu.RLock()
	defer s.mu.RUnlock()

	newSet := new(LockSet[T])
	newSet.data = make(map[T]empty, s.Size())

	for e := range s.data {
		newSet.data[e] = empty{}
	}

	return newSet
}

// Contains returns true if this set contains the specified element.
func (s *LockSet[T]) Contains(e T) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	_, found := s.data[e]

	return found
}

// ContainsAll returns true if this set contains all of the specified elements.
func (s *LockSet[T]) ContainsAll(c ...T) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, e := range c {
		_, found := s.data[e]
		if !found {
			return false
		}
	}

	return true
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

	if s.Size() != os.Size() {
		return false
	}

	for k := range s.data {
		_, ok := os.data[k]
		if !ok {
			return false
		}
	}

	return true
}

// ForEach performs the given handler for each elements in the set until all elements have been
// processed or the handler returns an error.
func (s *LockSet[T]) ForEach(handler func(e T) error) error {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for e := range s.data {
		if err := handler(e); err != nil {
			return err
		}
	}

	return nil
}

// IsEmpty returns true if this set contains no elements.
func (s *LockSet[T]) IsEmpty() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return len(s.data) == 0
}

// Remove removes the specified element from this set.
func (s *LockSet[T]) Remove(e T) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, found := s.data[e]
	if !found {
		return false
	}

	delete(s.data, e)
	return true
}

// RemoveAll removes all of the specified elements from this set.
func (s *LockSet[T]) RemoveAll(c ...T) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	isChanged := false

	for _, e := range c {
		_, found := s.data[e]
		if found {
			isChanged = true
			delete(s.data, e)
		}
	}

	return isChanged
}

// RemoveIf removes all of the elements of this set that satisfy the given predicate.
func (s *LockSet[T]) RemoveIf(filter func(T) bool) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	isChanged := false

	for e := range s.data {
		if filter(e) {
			delete(s.data, e)
			isChanged = true
		}
	}

	return isChanged
}

// RetainAll retains only the elements in this set that are contained in the specified collection.
func (s *LockSet[T]) RetainAll(c ...T) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	cSet := NewHashSet[T]()
	cSet.AddAll(c...)
	isChanged := false

	for e := range s.data {
		if !cSet.Contains(e) {
			delete(s.data, e)
			isChanged = true
		}
	}

	return isChanged
}

// Size returns the number of elements in this set.
func (s *LockSet[T]) Size() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return len(s.data)
}

// String returns the string representation of this set.
func (s *LockSet[T]) String() string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	buf := bytes.NewBufferString("set[")
	first := true
	for e := range s.data {
		if !first {
			buf.WriteString(" ")
		}
		first = false
		buf.WriteString(internal.ValueString(e))
	}
	buf.WriteString("]")
	return buf.String()
}

// ToSlice returns a slice containing all of the elements in this set.
func (s *LockSet[T]) ToSlice() []T {
	s.mu.RLock()
	defer s.mu.RUnlock()

	slice := make([]T, 0, s.Size())

	for e := range s.data {
		slice = append(slice, e)
	}

	return slice
}

// MarshalJSON marshals the set as a JSON array.
func (s *LockSet[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.ToSlice())
}

// UnmarshalJSON unmarshals a JSON array into the set.
func (s *LockSet[T]) UnmarshalJSON(b []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	var items []T
	if err := json.Unmarshal(b, &items); err != nil {
		return err
	}
	s.data = make(HashSet[T], len(items))
	for _, v := range items {
		s.data[v] = empty{}
	}
	return nil
}
