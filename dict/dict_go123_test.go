//go:build go1.23

package dict

import (
	"github.com/ghosind/go-assert"
)

func testDictIter(a *assert.Assertion, constructor dictConstructor) {
	d := constructor(testDataEn)
	a.EqualNow(d.Size(), len(testDataEn))

	for k, v := range d.Iter() {
		a.EqualNow(v, testDataEn[k])
	}

	for range d.Iter() {
		// yield should returns false
		break
	}
}

func testDictKeysIter(a *assert.Assertion, constructor dictConstructor) {
	d := constructor(testDataEn)
	a.EqualNow(d.Size(), len(testDataEn))

	for k := range d.KeysIter() {
		_, ok := testDataEn[k]
		a.TrueNow(ok)
	}

	for range d.KeysIter() {
		// yield should returns false
		break
	}
}

func testDictValuesIter(a *assert.Assertion, constructor dictConstructor) {
	d := constructor(testDataEn)
	a.EqualNow(d.Size(), len(testDataEn))

	valueSet := make(map[string]struct{})
	for _, v := range testDataEn {
		valueSet[v] = struct{}{}
	}

	for v := range d.ValuesIter() {
		_, ok := valueSet[v]
		a.TrueNow(ok)
	}

	for range d.ValuesIter() {
		// yield should returns false
		break
	}
}
