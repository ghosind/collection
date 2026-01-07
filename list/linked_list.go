package list

import (
	"bytes"
	"encoding/json"
	"sync"

	"github.com/ghosind/collection"
	"github.com/ghosind/collection/internal"
)

// LinkedListNode represents a node in the linked list.
type LinkedListNode[T any] struct {
	Value T
	Next  *LinkedListNode[T]
	Prev  *LinkedListNode[T]
}

// LinkedList represents a doubly linked list.
type LinkedList[T any] struct {
	head *LinkedListNode[T]
	tail *LinkedListNode[T]
	size int
	pool sync.Pool
}

// NewLinkedList creates and returns a new empty linked list.
func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{
		head: nil,
		tail: nil,
		size: 0,
		pool: sync.Pool{New: func() any {
			return &LinkedListNode[T]{}
		}},
	}
}

// NewLinkedListFrom creates and returns a new linked list containing the elements of the
// provided collection.
func NewLinkedListFrom[T any](c ...T) *LinkedList[T] {
	l := NewLinkedList[T]()
	for _, e := range c {
		l.Add(e)
	}
	return l
}

// Add adds the specified element to this collection.
func (l *LinkedList[T]) Add(e T) bool {
	newNode := l.getNode(e)
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.Next = newNode
		newNode.Prev = l.tail
		l.tail = newNode
	}
	l.size++
	return true
}

// AddAll adds all of the elements in the this collection.
func (l *LinkedList[T]) AddAll(c ...T) bool {
	for _, e := range c {
		l.Add(e)
	}
	return len(c) > 0
}

// AddAtIndex inserts the specified element to the specified position in this list.
func (l *LinkedList[T]) AddAtIndex(i int, e T) {
	internal.CheckIndex(i, l.size+1)

	newNode := l.getNode(e)

	if l.size == 0 {
		l.head = newNode
		l.tail = newNode
		l.size++
		return
	}

	switch i {
	case l.size: // append to the end
		if l.tail != nil {
			l.tail.Next = newNode
			newNode.Prev = l.tail
			l.tail = newNode
		}
	case 0: // insert at head
		newNode.Next = l.head
		if l.head != nil {
			l.head.Prev = newNode
		}
		l.head = newNode
	default: // insert in the middle
		current := l.head
		for j := 0; j < i; j++ {
			current = current.Next
		}
		newNode.Next = current
		newNode.Prev = current.Prev
		if current.Prev != nil {
			current.Prev.Next = newNode
		}
		current.Prev = newNode
	}
	l.size++
}

// Clear removes all of the elements from this collection.
func (l *LinkedList[T]) Clear() {
	for node := l.head; node != nil; {
		next := node.Next
		l.pool.Put(node)
		node = next
	}

	l.head = nil
	l.tail = nil
	l.size = 0
}

// Clone returns a copy of this list.
func (l *LinkedList[T]) Clone() collection.List[T] {
	clone := NewLinkedList[T]()
	for node := l.head; node != nil; node = node.Next {
		clone.Add(node.Value)
	}
	return clone
}

// Contains returns true if this collection contains the specified element.
func (l *LinkedList[T]) Contains(e T) bool {
	return l.IndexOf(e) >= 0
}

// ContainsAll returns true if this collection contains all of the elements in the specified
// collection.
func (l *LinkedList[T]) ContainsAll(c ...T) bool {
	for _, e := range c {
		if !l.Contains(e) {
			return false
		}
	}
	return true
}

// Equals compares this collection with the object pass from parameter.
func (l *LinkedList[T]) Equals(o any) bool {
	other, ok := o.(*LinkedList[T])
	if !ok {
		return false
	}
	if l.size != other.size {
		return false
	}
	node1 := l.head
	node2 := other.head
	for node1 != nil && node2 != nil {
		if !internal.Equal(node1.Value, node2.Value) {
			return false
		}
		node1 = node1.Next
		node2 = node2.Next
	}
	return true
}

// ForEach performs the given handler for each elements in the collection until all elements
// have been processed or the handler returns an error.
func (l *LinkedList[T]) ForEach(handler func(e T) error) error {
	for node := l.head; node != nil; node = node.Next {
		if err := handler(node.Value); err != nil {
			return err
		}
	}
	return nil
}

// Get returns the element at the specified position in this list.
func (l *LinkedList[T]) Get(i int) T {
	internal.CheckIndex(i, l.size)

	current := l.head
	for j := 0; j < i; j++ {
		current = current.Next
	}
	return current.Value
}

// IndexOf returns the index of the first occurrence of the specified element in this list, or -1
// if this list does not contain the element.
func (l *LinkedList[T]) IndexOf(e T) int {
	current := l.head
	for i := 0; i < l.size; i++ {
		if internal.Equal(current.Value, e) {
			return i
		}
		current = current.Next
	}
	return -1
}

// IsEmpty returns true if this collection contains no elements.
func (l *LinkedList[T]) IsEmpty() bool {
	return l.size == 0
}

// LastIndexOf returns the index of the last occurrence of the specified element in this list, or
// -1 if this list does not contain the element.
func (l *LinkedList[T]) LastIndexOf(e T) int {
	current := l.tail
	for i := l.size - 1; i >= 0; i-- {
		if internal.Equal(current.Value, e) {
			return i
		}
		current = current.Prev
	}
	return -1
}

// Remove removes the specified element from this collection.
func (l *LinkedList[T]) Remove(e T) bool {
	if l.size == 0 {
		return false
	}

	found := false
	current := l.head
	for current != nil {
		if internal.Equal(current.Value, e) {
			next := current.Next
			l.removeNode(current)
			found = true
			current = next
		} else {
			current = current.Next
		}
	}
	return found
}

// RemoveAll removes all of the elements in the specified collection from this collection.
func (l *LinkedList[T]) RemoveAll(c ...T) bool {
	if len(c) == 0 {
		return false
	}

	found := false
	for _, e := range c {
		if l.Remove(e) {
			found = true
		}
	}
	return found
}

// RemoveAtIndex removes the element at the specified position in this list.
func (l *LinkedList[T]) RemoveAtIndex(i int) T {
	internal.CheckIndex(i, l.size)
	current := l.head
	for j := 0; j < i; j++ {
		current = current.Next
	}

	val := current.Value
	l.removeNode(current)

	return val
}

// RemoveFirst removes the first occurrence of the specified element from this list, if it is present.
// Returns true if the element was removed.
func (l *LinkedList[T]) RemoveFirst(e T) bool {
	if l.size == 0 {
		return false
	}

	current := l.head
	for current != nil {
		if internal.Equal(current.Value, e) {
			l.removeNode(current)
			return true
		}
		current = current.Next
	}
	return false
}

// RemoveFirstN removes the first n occurrences of the specified element from this list.
// Returns the number of elements removed.
func (l *LinkedList[T]) RemoveFirstN(e T, n int) int {
	if l.size == 0 || n <= 0 {
		return 0
	}

	removedCount := 0
	current := l.head
	for current != nil && removedCount < n {
		if internal.Equal(current.Value, e) {
			next := current.Next
			l.removeNode(current)
			current = next
			removedCount++
		} else {
			current = current.Next
		}
	}
	return removedCount
}

// RemoveIf removes all of the elements of this collection that satisfy the given predicate.
func (l *LinkedList[T]) RemoveIf(f func(T) bool) bool {
	if l.size == 0 {
		return false
	}

	found := false
	current := l.head
	for current != nil {
		if f(current.Value) {
			next := current.Next
			l.removeNode(current)
			current = next
			found = true
		} else {
			current = current.Next
		}
	}
	return found
}

// RemoveLast removes the last occurrence of the specified element from this list, if it is present.
// Returns true if the element was removed.
func (l *LinkedList[T]) RemoveLast(e T) bool {
	if l.size == 0 {
		return false
	}

	current := l.tail
	for current != nil {
		if internal.Equal(current.Value, e) {
			l.removeNode(current)
			return true
		}
		current = current.Prev
	}
	return false
}

// RemoveLastN removes the last n occurrences of the specified element from this list.
// Returns the number of elements removed.
func (l *LinkedList[T]) RemoveLastN(e T, n int) int {
	if l.size == 0 || n <= 0 {
		return 0
	}

	removedCount := 0
	current := l.tail
	for current != nil && removedCount < n {
		if internal.Equal(current.Value, e) {
			prev := current.Prev
			l.removeNode(current)
			current = prev
			removedCount++
		} else {
			current = current.Prev
		}
	}
	return removedCount
}

// RetainAll retains only the elements in this collection that are contained in the specified
// collection.
func (l *LinkedList[T]) RetainAll(c ...T) bool {
	if len(c) == 0 {
		return clearListForRetainAll[T](l)
	}

	found := false
	current := l.head
	for current != nil {
		shouldRetain := false
		for _, e := range c {
			if internal.Equal(current.Value, e) {
				shouldRetain = true
				break
			}
		}
		if !shouldRetain {
			next := current.Next
			l.removeNode(current)
			current = next
			found = true
		} else {
			current = current.Next
		}
	}
	return found
}

// Set replaces the element at the specified position in this list with the specified element.
func (l *LinkedList[T]) Set(i int, e T) T {
	internal.CheckIndex(i, l.size+1)

	if i == l.size { // append to the end
		l.Add(e)
		return *new(T) // return zero value
	}

	current := l.head
	for j := 0; j < i; j++ {
		current = current.Next
	}
	oldValue := current.Value
	current.Value = e
	return oldValue
}

// Size returns the number of elements in this collection.
func (l *LinkedList[T]) Size() int {
	return l.size
}

// String returns the string representation of this collection.
func (l *LinkedList[T]) String() string {
	buf := bytes.NewBufferString("list[")
	for node := l.head; node != nil; node = node.Next {
		buf.WriteString(internal.ValueString(node.Value))
		if node.Next != nil {
			buf.WriteString(" ")
		}
	}
	buf.WriteString("]")
	return buf.String()
}

// SubList returns a view of the portion of this list between the specified fromIndex, inclusive,
// and toIndex, exclusive.
func (l *LinkedList[T]) SubList(fromIndex, toIndex int) collection.List[T] {
	internal.CheckIndex(fromIndex, l.size)
	internal.CheckIndex(toIndex, l.size)

	subList := NewLinkedList[T]()
	current := l.head
	for i := 0; i < fromIndex; i++ {
		current = current.Next
	}
	for i := fromIndex; i < toIndex; i++ {
		subList.Add(current.Value)
		current = current.Next
	}
	return subList
}

// ToSlice returns a slice containing all of the elements in this collection.
func (l *LinkedList[T]) ToSlice() []T {
	slice := make([]T, 0, l.size)
	for node := l.head; node != nil; node = node.Next {
		slice = append(slice, node.Value)
	}
	return slice
}

// Trim removes the first n elements from this list. Returns the number of elements removed.
func (l *LinkedList[T]) Trim(n int) int {
	removedCount := 0
	for removedCount < n && l.size > 0 {
		l.removeNode(l.head)
		removedCount++
	}
	return removedCount
}

// TrimLast removes the last n elements from this list. Returns the number of elements removed.
func (l *LinkedList[T]) TrimLast(n int) int {
	removedCount := 0
	for removedCount < n && l.size > 0 {
		l.removeNode(l.tail)
		removedCount++
	}
	return removedCount
}

// MarshalJSON marshals the linked list as a JSON array.
func (l *LinkedList[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.ToSlice())
}

// UnmarshalJSON unmarshals a JSON array into the linked list.
func (l *LinkedList[T]) UnmarshalJSON(b []byte) error {
	var items []T
	if err := json.Unmarshal(b, &items); err != nil {
		return err
	}
	// Ensure pool is initialized for zero-value LinkedList
	if l.pool.New == nil {
		l.pool = sync.Pool{New: func() any {
			return &LinkedListNode[T]{}
		}}
	}
	l.Clear()
	if len(items) > 0 {
		l.AddAll(items...)
	}
	return nil
}

func (l *LinkedList[T]) removeNode(node *LinkedListNode[T]) {
	if node.Prev != nil {
		node.Prev.Next = node.Next
	} else {
		l.head = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	} else {
		l.tail = node.Prev
	}
	l.pool.Put(node)
	l.size--
}

// getNode retrieves an empty node from the pool.
func (l *LinkedList[T]) getNode(e T) *LinkedListNode[T] {
	node := l.pool.Get().(*LinkedListNode[T])
	node.Value = e
	node.Next = nil
	node.Prev = nil
	return node
}
