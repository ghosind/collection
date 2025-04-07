package dict

import (
	"testing"

	"github.com/ghosind/collection"
	"github.com/ghosind/go-assert"
)

func TestHashDict(t *testing.T) {
	a := assert.New(t)
	constructor := func() collection.Dict[string, string] {
		return NewHashDict[string, string]()
	}

	testDict(a, constructor)
}
