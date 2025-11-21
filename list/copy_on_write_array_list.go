package list

import (
	"bytes"
	"encoding/json"
	"sync"

	"github.com/ghosind/collection"
	"github.com/ghosind/collection/internal"
)

type CopyOnWriteArrayList[T any] struct {
	data []T
	mu   sync.RWMutex
}

func NewCopyOnWriteArrayList[T any]() *CopyOnWriteArrayList[T] {
	l := &CopyOnWriteArrayList[T]{
		data: make([]T, 0),
	}

	return l
}

// Add adds the specified element to this collection.
func (l *CopyOnWriteArrayList[T]) Add(e T) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	newData := make([]T, len(l.data)+1)
	copy(newData, l.data)
	newData[len(l.data)] = e
	l.data = newData

	return true
}

// AddAll adds all of the elements in the this collection.
func (l *CopyOnWriteArrayList[T]) AddAll(c ...T) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	newData := make([]T, len(l.data)+len(c))
	copy(newData, l.data)
	copy(newData[len(l.data):], c)
	l.data = newData

	return true
}

// AddAtIndex inserts the specified element to the specified position in this list.
func (l *CopyOnWriteArrayList[T]) AddAtIndex(i int, e T) {
	l.mu.Lock()
	defer l.mu.Unlock()

	internal.CheckIndex(i, len(l.data)+1)

	newData := make([]T, len(l.data)+1)
	copy(newData, l.data[:i])
	newData[i] = e
	if i < len(l.data) {
		copy(newData[i+1:], l.data[i:])
	}
	l.data = newData
}

// Clear removes all of the elements from this collection.
func (l *CopyOnWriteArrayList[T]) Clear() {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.data = make([]T, 0)
}

// Clone returns a copy of this list.
func (l *CopyOnWriteArrayList[T]) Clone() collection.List[T] {
	data := l.data
	clonedData := make([]T, len(data))
	copy(clonedData, data)

	return &CopyOnWriteArrayList[T]{
		data: clonedData,
		mu:   sync.RWMutex{},
	}
}

// Contains returns true if this collection contains the specified element.
func (l *CopyOnWriteArrayList[T]) Contains(e T) bool {
	data := l.data

	for _, v := range data {
		if internal.Equal(v, e) {
			return true
		}
	}

	return false
}

// ContainsAll returns true if this collection contains all of the elements in the specified
// collection.
func (l *CopyOnWriteArrayList[T]) ContainsAll(c ...T) bool {
	data := l.data

	for _, v := range c {
		found := false
		for _, dv := range data {
			if internal.Equal(v, dv) {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}

	return true
}

// Equals compares this collection with the object pass from parameter.
func (l *CopyOnWriteArrayList[T]) Equals(o any) bool {
	ol, ok := o.(*CopyOnWriteArrayList[T])
	if !ok {
		return false
	}

	ldata := l.data
	oldata := ol.data

	if len(ldata) != len(oldata) {
		return false
	}

	for i, v := range ldata {
		if !internal.Equal(v, oldata[i]) {
			return false
		}
	}

	return true
}

// ForEach performs the given handler for each elements in the collection until all elements
// have been processed or the handler returns an error.
func (l *CopyOnWriteArrayList[T]) ForEach(handler func(T) error) error {
	data := l.data

	for _, v := range data {
		if err := handler(v); err != nil {
			return err
		}
	}

	return nil
}

// Get returns the element at the specified position in this list.
func (l *CopyOnWriteArrayList[T]) Get(i int) T {
	data := l.data

	internal.CheckIndex(i, len(data))

	return data[i]
}

// IndexOf returns the index of the first occurrence of the specified element in this list, or -1
// if this list does not contain the element.
func (l *CopyOnWriteArrayList[T]) IndexOf(e T) int {
	data := l.data

	for i, v := range data {
		if internal.Equal(v, e) {
			return i
		}
	}

	return -1
}

// IsEmpty returns true if this collection contains no elements.
func (l *CopyOnWriteArrayList[T]) IsEmpty() bool {
	data := l.data
	return len(data) == 0
}

// LastIndexOf returns the index of the last occurrence of the specified element in this list, or
// -1 if this list does not contain the element.
func (l *CopyOnWriteArrayList[T]) LastIndexOf(e T) int {
	data := l.data

	for i := len(data) - 1; i >= 0; i-- {
		if internal.Equal(data[i], e) {
			return i
		}
	}

	return -1
}

// Remove removes the specified element from this collection.
func (l *CopyOnWriteArrayList[T]) Remove(e T) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	if len(l.data) == 0 {
		return false
	}

	newData := make([]T, 0, len(l.data))
	removed := false

	for _, v := range l.data {
		if !internal.Equal(v, e) {
			newData = append(newData, v)
		} else {
			removed = true
		}
	}

	if removed {
		l.data = newData
	}

	return removed
}

// RemoveAll removes all of the elements in the specified collection from this collection.
func (l *CopyOnWriteArrayList[T]) RemoveAll(c ...T) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	if len(c) == 0 {
		return false
	}

	newData := make([]T, 0, len(l.data))
	removed := false

	for _, v := range l.data {
		found := false
		for _, cv := range c {
			if internal.Equal(v, cv) {
				found = true
				break
			}
		}
		if !found {
			newData = append(newData, v)
		} else {
			removed = true
		}
	}

	if removed {
		l.data = newData
	}

	return removed
}

// RemoveAtIndex removes the element at the specified position in this list.
func (l *CopyOnWriteArrayList[T]) RemoveAtIndex(i int) T {
	l.mu.Lock()
	defer l.mu.Unlock()

	data := l.data
	internal.CheckIndex(i, len(data))

	old := data[i]
	newData := make([]T, 0, len(data)-1)

	newData = append(newData, data[:i]...)
	if i+1 < len(data) {
		newData = append(newData, data[i+1:]...)
	}

	l.data = newData
	return old
}

// RemoveIf removes all of the elements of this collection that satisfy the given predicate.
func (l *CopyOnWriteArrayList[T]) RemoveIf(f func(T) bool) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	if len(l.data) == 0 {
		return false
	}

	newData := make([]T, 0, len(l.data))
	removed := false

	for _, v := range l.data {
		if !f(v) {
			newData = append(newData, v)
		} else {
			removed = true
		}
	}

	if removed {
		l.data = newData
	}

	return removed
}

// RetainAll retains only the elements in this collection that are contained in the specified
// collection.
func (l *CopyOnWriteArrayList[T]) RetainAll(c ...T) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	if len(c) == 0 {
		if len(l.data) == 0 {
			return false
		}

		l.data = make([]T, 0)
		return true
	}

	newData := make([]T, 0, len(l.data))
	changed := false

	for _, v := range l.data {
		found := false
		for _, cv := range c {
			if internal.Equal(v, cv) {
				found = true
				break
			}
		}
		if found {
			newData = append(newData, v)
		} else {
			changed = true
		}
	}

	if changed {
		l.data = newData
	}

	return changed
}

// Set replaces the element at the specified position in this list with the specified element.
func (l *CopyOnWriteArrayList[T]) Set(i int, e T) T {
	l.mu.Lock()
	defer l.mu.Unlock()

	internal.CheckIndex(i, len(l.data)+1)

	var old T
	var newData []T

	if i == len(l.data) {
		newData = make([]T, len(l.data)+1)
	} else {
		old = l.data[i]
		newData = make([]T, len(l.data))
	}

	copy(newData, l.data)
	newData[i] = e
	l.data = newData

	return old
}

// Size returns the number of elements in this collection.
func (l *CopyOnWriteArrayList[T]) Size() int {
	data := l.data
	return len(data)
}

// String returns the string representation of this collection.
func (l *CopyOnWriteArrayList[T]) String() string {
	buf := bytes.NewBufferString("list[")
	first := true
	data := l.data
	for _, v := range data {
		if !first {
			buf.WriteString(" ")
		}
		first = false
		buf.WriteString(internal.ValueString(v))
	}
	buf.WriteString("]")
	return buf.String()
}

// ToSlice returns a slice containing all of the elements in this collection.
func (l *CopyOnWriteArrayList[T]) ToSlice() []T {
	slice := make([]T, len(l.data))
	copy(slice, l.data)

	return slice
}

// MarshalJSON marshals the copy-on-write list as a JSON array.
func (l *CopyOnWriteArrayList[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.ToSlice())
}

// UnmarshalJSON unmarshals a JSON array into the copy-on-write list.
func (l *CopyOnWriteArrayList[T]) UnmarshalJSON(b []byte) error {
	var items []T
	if err := json.Unmarshal(b, &items); err != nil {
		return err
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	l.data = make([]T, len(items))
	copy(l.data, items)
	return nil
}
