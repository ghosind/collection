//go:build go1.23

package list

import "github.com/ghosind/go-assert"

func testListIter(a *assert.Assertion, constructor listConstructor) {
	l := constructor()

	l.AddAll(testData...)
	a.EqualNow(len(testData), l.Size())
	res := make([]int, 0, l.Size())

	for v := range l.Iter() {
		res = append(res, v)
	}

	a.EqualNow(testData, res)

	for range l.Iter() {
		// yield should returns false
		break
	}
}
