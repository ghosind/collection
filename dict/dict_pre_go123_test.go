//go:build !go1.23

package dict

import (
	"reflect"

	"github.com/ghosind/go-assert"
)

func testDictIter(a *assert.Assertion, constructor dictTestConstructor) {
	d := constructor()

	ty := reflect.TypeOf(d)
	_, ok := ty.MethodByName("Iter")
	a.NotTrueNow(ok)
}

func testDictKeysIter(a *assert.Assertion, constructor dictTestConstructor) {
	d := constructor()
	for k, v := range testDataEn {
		d.Put(k, v)
	}
	a.EqualNow(d.Size(), len(testDataEn))

	for k := range d.KeysIter() {
		_, ok := testDataEn[k]
		a.TrueNow(ok)
	}
}

func testDictValuesIter(a *assert.Assertion, constructor dictTestConstructor) {
	d := constructor()
	for k, v := range testDataEn {
		d.Put(k, v)
	}
	a.EqualNow(d.Size(), len(testDataEn))

	valueSet := make(map[string]struct{})
	for _, v := range testDataEn {
		valueSet[v] = struct{}{}
	}

	for v := range d.ValuesIter() {
		_, ok := valueSet[v]
		a.TrueNow(ok)
	}
}
