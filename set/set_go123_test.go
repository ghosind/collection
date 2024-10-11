//go:build go1.23

package set

import (
	"github.com/ghosind/collection"
	"github.com/ghosind/go-assert"
)

func testSetIter(a *assert.Assertion, set collection.Set[int]) {
	records := map[int]int{}

	for e := range set.Iter() {
		records[e]++
	}

	a.EqualNow(len(records), set.Size())
	for _, v := range records {
		a.EqualNow(v, 1)
	}
}
