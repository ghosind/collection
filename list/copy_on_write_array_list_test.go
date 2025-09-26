package list

import (
	"testing"

	"github.com/ghosind/collection"
	"github.com/ghosind/go-assert"
)

func TestCopyOnWriteArrayList(t *testing.T) {
	a := assert.New(t)
	constructor := func() collection.List[int] {
		return NewCopyOnWriteArrayList[int]()
	}

	testList(a, constructor)
}
