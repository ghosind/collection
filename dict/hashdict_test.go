package dict

import (
	"testing"

	"github.com/ghosind/collection"
	"github.com/ghosind/go-assert"
)

func TestHashDictionary(t *testing.T) {
	a := assert.New(t)
	constructor := func() collection.Dict[string, string] {
		return NewHashDictionary[string, string]()
	}

	testDict(a, constructor)
}
