//go:build go1.23

package dict

import (
	"github.com/ghosind/go-assert"
)

func testDictIter(a *assert.Assertion, constructor dictTestConstructor) {
	d := constructor()
	for k, v := range testDataEn {
		d.Put(k, v)
	}
	a.EqualNow(d.Size(), len(testDataEn))

	for k, v := range d.Iter() {
		a.EqualNow(v, testDataEn[k])
	}
}
