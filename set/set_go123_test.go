//go:build go1.23

package set

import (
	"github.com/ghosind/go-assert"
)

func testSetIter(a *assert.Assertion, constructor setTestConstructor) {
	set1 := constructor()
	set1.AddAll(testNums1...)

	set2 := constructor()

	for e := range set1.Iter() {
		set2.Add(e)
	}

	a.TrueNow(set1.Equals(set2))
}
