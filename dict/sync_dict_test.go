package dict

import (
	"testing"

	"github.com/ghosind/collection"
	"github.com/ghosind/go-assert"
)

func TestSyncDict(t *testing.T) {
	a := assert.New(t)
	constructor := func() collection.Dict[string, string] {
		return NewSyncDict[string, string]()
	}

	testDict(a, constructor)
}
