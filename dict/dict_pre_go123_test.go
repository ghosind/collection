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
