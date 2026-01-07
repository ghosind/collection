package stack

import (
	"testing"

	"github.com/ghosind/collection"
	"github.com/ghosind/go-assert"
)

func TestStack(t *testing.T) {
	a := assert.New(t)
	stack := NewStack[int]()

	a.EqualNow(0, stack.Size())
	a.TrueNow(stack.IsEmpty())

	stack.Push(10)
	a.EqualNow(1, stack.Size())
	a.EqualNow(10, stack.Peek())

	stack.Push(20)
	a.EqualNow(2, stack.Size())
	a.EqualNow(20, stack.Peek())

	a.EqualNow(stack.String(), "stack[10 20]")

	stack2 := NewStackFrom([]int{10, 20})
	a.TrueNow(stack2.Equals(stack))

	v := stack.Pop()
	a.EqualNow(20, v)
	a.EqualNow(1, stack.Size())
	a.EqualNow(10, stack.Peek())

	v2 := stack.Pop()
	a.EqualNow(10, v2)
	a.TrueNow(stack.IsEmpty())

	// Pop/Peek on empty should panic with ErrOutOfBounds
	a.PanicOfNow(func() { stack.Pop() }, collection.ErrOutOfBounds)
	a.PanicOfNow(func() { stack.Peek() }, collection.ErrOutOfBounds)
}

func TestStack_Clone(t *testing.T) {
	a := assert.New(t)
	stack := NewStack[int]()
	stack.Push(10)
	stack.Push(20)

	clone := stack.Clone()
	a.TrueNow(clone.Equals(stack))
	a.EqualNow(2, clone.Size())
	a.EqualNow(20, clone.Peek())

	v := clone.Pop()
	a.EqualNow(20, v)
	a.EqualNow(1, clone.Size())
	a.EqualNow(10, clone.Peek())

	a.NotTrueNow(clone.Equals(stack))

	// Original stack should remain unchanged
	a.EqualNow(2, stack.Size())
	a.EqualNow(20, stack.Peek())
}

func TestStack_Equals(t *testing.T) {
	a := assert.New(t)
	stack1 := NewStackFrom([]int{1, 2, 3})
	stack2 := NewStackFrom([]int{1, 2, 3})
	stack3 := NewStackFrom([]int{1, 2, 4})
	stack4 := NewStackFrom([]int{1, 2, 3, 4})

	a.TrueNow(stack1.Equals(stack2))
	a.NotTrueNow(stack1.Equals(stack3))
	a.NotTrueNow(stack1.Equals(stack4))
	a.NotTrueNow(stack1.Equals("not a stack"))
}
