package set

import (
	"sync"

	"github.com/ghosind/collection"
	"github.com/ghosind/utils"
)

// ConcurrentHashSet is a thread-safe set implementation that uses a Golang builtin map to store
// its elements.
type ConcurrentHashSet[T comparable] struct {
	data *HashSet[T]

	mutex sync.RWMutex
}

// NewConcurrentHashSet creates and returns am empty ConcurrentHashSet with the specified type.
func NewConcurrentHashSet[T comparable]() *ConcurrentHashSet[T] {
	set := new(ConcurrentHashSet[T])
	set.data = NewHashSet[T]()

	return set
}

// Add adds the specified element to this set.
func (set *ConcurrentHashSet[T]) Add(e T) bool {
	set.mutex.Lock()
	defer set.mutex.Unlock()

	return set.data.Add(e)
}

// AddAll adds all of the specified elements to this set.
func (set *ConcurrentHashSet[T]) AddAll(c ...T) bool {
	set.mutex.Lock()
	defer set.mutex.Unlock()

	return set.data.AddAll(c...)
}

// Clear removes all of the elements from this set.
func (set *ConcurrentHashSet[T]) Clear() {
	set.mutex.Lock()
	defer set.mutex.Unlock()

	set.data.Clear()
}

// Clone returns a copy of this set.
func (set *ConcurrentHashSet[T]) Clone() collection.Set[T] {
	set.mutex.RLock()
	defer set.mutex.RUnlock()

	newSet := new(ConcurrentHashSet[T])
	newSet.data = set.data.Clone().(*HashSet[T])

	return newSet
}

// Contains returns true if this set contains the specified element.
func (set *ConcurrentHashSet[T]) Contains(e T) bool {
	set.mutex.RLock()
	defer set.mutex.RUnlock()

	return set.data.Contains(e)
}

// ContainsAll returns true if this set contains all of the specified elements.
func (set *ConcurrentHashSet[T]) ContainsAll(c ...T) bool {
	set.mutex.RLock()
	defer set.mutex.RUnlock()

	return set.data.ContainsAll(c...)
}

// Equals compares set with the object pass from parameter.
func (set *ConcurrentHashSet[T]) Equals(o any) bool {
	if !utils.IsSameType(set, o) {
		return false
	}

	s := o.(*ConcurrentHashSet[T])
	set.mutex.RLock()
	s.mutex.RLock()
	defer set.mutex.RUnlock()
	defer s.mutex.RUnlock()

	if len(*set.data) != len(*s.data) {
		return false
	}

	for k := range *set.data {
		_, ok := (*s.data)[k]
		if !ok {
			return false
		}
	}

	return true
}

// ForEach performs the given handler for each elements in the set until all elements have been processed or the handler returns an error.
func (set *ConcurrentHashSet[T]) ForEach(handler func(e T) error) error {
	set.mutex.RLock()
	defer set.mutex.RUnlock()

	for e := range *set.data {
		if err := handler(e); err != nil {
			return err
		}
	}

	return nil
}

// IsEmpty returns true if this set contains no elements.
func (set *ConcurrentHashSet[T]) IsEmpty() bool {
	return set.Size() == 0
}

// Iter returns a channel of all elements in this set.
func (set *ConcurrentHashSet[T]) Iter() <-chan T {
	ch := make(chan T)

	go func() {
		set.mutex.RLock()
		defer set.mutex.RUnlock()

		for e := range *set.data {
			ch <- e
		}

		close(ch)
	}()

	return ch
}

// Remove removes the specified element from this set.
func (set *ConcurrentHashSet[T]) Remove(e T) bool {
	set.mutex.Lock()
	defer set.mutex.Unlock()

	_, found := (*set.data)[e]
	if !found {
		return false
	}

	delete(*set.data, e)
	return true
}

// RemoveAll removes all of the specified elements from this set.
func (set *ConcurrentHashSet[T]) RemoveAll(c ...T) bool {
	set.mutex.Lock()
	defer set.mutex.Unlock()

	isChanged := false

	for _, e := range c {
		_, found := (*set.data)[e]
		if found {
			isChanged = true
			delete(*set.data, e)
		}
	}

	return isChanged
}

// Size returns the number of elements in this set.
func (set *ConcurrentHashSet[T]) Size() int {
	set.mutex.RLock()
	defer set.mutex.RUnlock()

	return len(*set.data)
}

// ToSlice returns a slice containing all of the elements in this set.
func (set *ConcurrentHashSet[T]) ToSlice() []T {
	set.mutex.RLock()
	defer set.mutex.RUnlock()

	slice := make([]T, 0, len(*set.data))

	for e := range *set.data {
		slice = append(slice, e)
	}

	return slice
}
