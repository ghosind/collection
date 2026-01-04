package list

import (
	"bytes"

	"github.com/ghosind/collection"
	"github.com/ghosind/collection/internal"
)

// Stack represents a stack data structure based on ArrayList.
type Stack[T any] struct {
	ArrayList[T]
}

// NewStack creates and returns a new Stack.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

// NewStackFrom creates and returns a new Stack containing the elements of the
// provided collection.
func NewStackFrom[T any](c []T) *Stack[T] {
	stack := NewStack[T]()
	for _, e := range c {
		stack.Push(e)
	}
	return stack
}

// topIndex returns the index of the top element in the stack.
func (s *Stack[T]) topIndex() int {
	return s.Size() - 1
}

// Clone returns a shallow copy of this stack.
func (s *Stack[T]) Clone() collection.Stack[T] {
	clone := NewStack[T]()
	clone.ArrayList = *(s.ArrayList.Clone().(*ArrayList[T]))
	return clone
}

// Equals checks whether this stack is equal to another stack.
func (s *Stack[T]) Equals(o any) bool {
	other, ok := o.(*Stack[T])
	if !ok {
		return false
	}
	return s.ArrayList.Equals(&other.ArrayList)
}

// Peek returns the element at the top of this stack without removing it.
func (s *Stack[T]) Peek() T {
	return s.Get(s.topIndex())
}

// Pop removes and returns the element at the top of this stack.
func (s *Stack[T]) Pop() T {
	return s.RemoveAtIndex(s.topIndex())
}

// Push adds an element to the top of this stack.
func (s *Stack[T]) Push(e T) {
	s.Add(e)
}

// String returns the string representation of this stack.
func (s *Stack[T]) String() string {
	buf := bytes.NewBufferString("stack[")
	for i := 0; i < s.Size(); i++ {
		if i > 0 {
			buf.WriteString(" ")
		}
		buf.WriteString(internal.ValueString(s.Get(i)))
	}
	buf.WriteString("]")

	return buf.String()
}
