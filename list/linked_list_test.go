package list

import (
	"testing"

	"github.com/ghosind/collection"
	"github.com/ghosind/go-assert"
)

func TestLinkedList(t *testing.T) {
	a := assert.New(t)
	constructor := func() collection.List[int] {
		return NewLinkedList[int]()
	}

	testList(a, constructor)
}
