package list

// Stack represents a stack data structure based on ArrayList.
type Stack[T any] struct {
	ArrayList[T]
}

// NewStack creates and returns a new Stack.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

// topIndex returns the index of the top element in the stack.
func (s *Stack[T]) topIndex() int {
	return s.Size() - 1
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
