package list

import (
	"sync"

	"github.com/ghosind/collection"
)

// LockList is a thread-safe list that wraps another list with read-write locks.
type LockList[T any] struct {
	data collection.List[T]
	mu   sync.RWMutex
}

// NewLockList creates a new LockList.
func NewLockList[T any](data collection.List[T]) *LockList[T] {
	l := new(LockList[T])
	l.data = data

	return l
}

// Add adds the specified element to this collection.
func (l *LockList[T]) Add(e T) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	return l.data.Add(e)
}

// AddAll adds all of the elements in the this collection.
func (l *LockList[T]) AddAll(c ...T) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	return l.data.AddAll(c...)
}

// AddAtIndex inserts the specified element to the specified position in this list.
func (l *LockList[T]) AddAtIndex(index int, e T) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.data.AddAtIndex(index, e)
}

// Clear removes all of the elements from this collection.
func (l *LockList[T]) Clear() {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.data.Clear()
}

// Clone returns a copy of this list.
func (l *LockList[T]) Clone() collection.List[T] {
	l.mu.RLock()
	defer l.mu.RUnlock()

	cloned := l.data.Clone()

	return NewLockList[T](cloned)
}

// Contains returns true if this collection contains the specified element.
func (l *LockList[T]) Contains(e T) bool {
	l.mu.RLock()
	defer l.mu.RUnlock()

	return l.data.Contains(e)
}

// ContainsAll returns true if this collection contains all of the elements in the specified
// collection.
func (l *LockList[T]) ContainsAll(c ...T) bool {
	l.mu.RLock()
	defer l.mu.RUnlock()

	return l.data.ContainsAll(c...)
}

// Equals compares this collection with the object pass from parameter.
func (l *LockList[T]) Equals(o any) bool {
	lo, ok := o.(*LockList[T])
	if !ok {
		return false
	}

	l.mu.RLock()
	defer l.mu.RUnlock()
	lo.mu.RLock()
	defer lo.mu.RUnlock()

	return l.data.Equals(lo.data)
}

// ForEach performs the given handler for each elements in the collection until all elements
// have been processed or the handler returns an error.
func (l *LockList[T]) ForEach(handler func(e T) error) error {
	l.mu.RLock()
	defer l.mu.RUnlock()

	return l.data.ForEach(handler)
}

// Get returns the element at the specified position in this list.
func (l *LockList[T]) Get(index int) T {
	l.mu.RLock()
	defer l.mu.RUnlock()

	return l.data.Get(index)
}

// IndexOf returns the index of the first occurrence of the specified element in this list, or -1
// if this list does not contain the element.
func (l *LockList[T]) IndexOf(e T) int {
	l.mu.RLock()
	defer l.mu.RUnlock()

	return l.data.IndexOf(e)
}

// IsEmpty returns true if this collection contains no elements.
func (l *LockList[T]) IsEmpty() bool {
	l.mu.RLock()
	defer l.mu.RUnlock()

	return l.data.IsEmpty()
}

// LastIndexOf returns the index of the last occurrence of the specified element in this list, or
// -1 if this list does not contain the element.
func (l *LockList[T]) LastIndexOf(e T) int {
	l.mu.RLock()
	defer l.mu.RUnlock()

	return l.data.LastIndexOf(e)
}

// Remove removes the specified element from this collection.
func (l *LockList[T]) Remove(e T) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	return l.data.Remove(e)
}

// RemoveAll removes all of the elements in the specified collection from this collection.
func (l *LockList[T]) RemoveAll(c ...T) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	return l.data.RemoveAll(c...)
}

// RemoveAtIndex removes the element at the specified position in this list.
func (l *LockList[T]) RemoveAtIndex(index int) T {
	l.mu.Lock()
	defer l.mu.Unlock()

	return l.data.RemoveAtIndex(index)
}

// RemoveFirst removes the first occurrence of the specified element from this list, if it is present.
// Returns true if the element was removed.
func (l *LockList[T]) RemoveFirst(e T) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	return l.data.RemoveFirst(e)
}

// RemoveFirstN removes the first n occurrences of the specified element from this list.
// Returns the number of elements removed.
func (l *LockList[T]) RemoveFirstN(e T, n int) int {
	l.mu.Lock()
	defer l.mu.Unlock()

	return l.data.RemoveFirstN(e, n)
}

// RemoveIf removes all of the elements of this collection that satisfy the given predicate.
func (l *LockList[T]) RemoveIf(filter func(T) bool) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	return l.data.RemoveIf(filter)
}

// RemoveLast removes the last occurrence of the specified element from this list, if it is present.
// Returns true if the element was removed.
func (l *LockList[T]) RemoveLast(e T) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	return l.data.RemoveLast(e)
}

// RemoveLastN removes the last n occurrences of the specified element from this list.
// Returns the number of elements removed.
func (l *LockList[T]) RemoveLastN(e T, n int) int {
	l.mu.Lock()
	defer l.mu.Unlock()

	return l.data.RemoveLastN(e, n)
}

// RetainAll retains only the elements in this collection that are contained in the specified
// collection.
func (l *LockList[T]) RetainAll(c ...T) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	return l.data.RetainAll(c...)
}

// Set replaces the element at the specified position in this list with the specified element.
func (l *LockList[T]) Set(index int, e T) T {
	l.mu.Lock()
	defer l.mu.Unlock()

	return l.data.Set(index, e)
}

// Size returns the number of elements in this collection.
func (l *LockList[T]) Size() int {
	l.mu.RLock()
	defer l.mu.RUnlock()

	return l.data.Size()
}

// String returns the string representation of this collection.
func (l *LockList[T]) String() string {
	l.mu.RLock()
	defer l.mu.RUnlock()

	return l.data.String()
}

// SubList returns a view of the portion of this list between the specified fromIndex, inclusive,
// and toIndex, exclusive.
func (l *LockList[T]) SubList(fromIndex, toIndex int) collection.List[T] {
	l.mu.RLock()
	defer l.mu.RUnlock()

	sub := l.data.SubList(fromIndex, toIndex)

	return NewLockList[T](sub)
}

// ToSlice returns a slice containing all of the elements in this collection.
func (l *LockList[T]) ToSlice() []T {
	l.mu.RLock()
	defer l.mu.RUnlock()

	return l.data.ToSlice()
}

// Trim removes the first n elements from this list. Returns the number of elements removed.
func (l *LockList[T]) Trim(n int) int {
	l.mu.Lock()
	defer l.mu.Unlock()

	return l.data.Trim(n)
}

// TrimLast removes the last n elements from this list. Returns the number of elements removed.
func (l *LockList[T]) TrimLast(n int) int {
	l.mu.Lock()
	defer l.mu.Unlock()

	return l.data.TrimLast(n)
}

// MarshalJSON marshals the linked list as a JSON array.
func (l *LockList[T]) MarshalJSON() ([]byte, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	return l.data.MarshalJSON()
}

// UnmarshalJSON unmarshals a JSON array into the linked list.
func (l *LockList[T]) UnmarshalJSON(b []byte) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	return l.data.UnmarshalJSON(b)
}
