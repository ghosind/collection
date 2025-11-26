package list

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
