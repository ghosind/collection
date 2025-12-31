//go:build go1.23

package set

import (
	"github.com/ghosind/go-assert"
)

func testSetIter(a *assert.Assertion, constructor setConstructor) {
	set1 := constructor(testNums1)
	set2 := constructor()

	for e := range set1.Iter() {
		set2.Add(e)
	}

	a.TrueNow(set1.Equals(set2))

	for range set1.Iter() {
		// yield should returns false
		break
	}
}
