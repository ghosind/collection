//go:build !go1.23

package dict

import (
	"github.com/ghosind/collection"
	"github.com/ghosind/go-assert"
)

func testDictIter(a *assert.Assertion, m collection.Dict[int, int], data []int) {
	// Do nothing
}
