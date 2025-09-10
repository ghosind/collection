package list

import (
	"testing"

	"github.com/ghosind/collection"
	"github.com/ghosind/go-assert"
)

func TestArrayList(t *testing.T) {
	a := assert.New(t)
	constructor := func() collection.List[int] {
		return NewArrayList[int]()
	}

	testList(a, constructor)
}
