package set

import (
	"testing"

	"github.com/ghosind/collection"
	"github.com/ghosind/go-assert"
)

func TestHashSet(t *testing.T) {
	a := assert.New(t)
	constructor := func() collection.Set[int] {
		return NewHashSet[int]()
	}

	testSet(a, constructor)

	// TODO: Move to set_test.go
	testSetRemoveIf(a, constructor)
	testSetRetainAll(a, constructor)
}
