package list

import (
	"testing"

	"github.com/ghosind/collection"
	"github.com/ghosind/go-assert"
)

func TestLinkedList(t *testing.T) {
	a := assert.New(t)
	constructor := func(initData ...[]int) collection.List[int] {
		if len(initData) > 0 && len(initData[0]) > 0 {
			return NewLinkedListFrom(initData[0]...)
		}
		return NewLinkedList[int]()
	}

	testList(a, constructor)
}
