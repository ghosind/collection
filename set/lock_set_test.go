package set

import (
	"testing"

	"github.com/ghosind/collection"
	"github.com/ghosind/go-assert"
)

var lockSetConstructor = func(initData ...[]int) collection.Set[int] {
	if len(initData) > 0 && len(initData[0]) > 0 {
		return NewLockSetFrom(initData[0]...)
	}
	return NewLockSet[int]()
}

func TestLockSet(t *testing.T) {
	a := assert.New(t)

	testSet(a, lockSetConstructor)
}

func BenchmarkLockSet(b *testing.B) {
	benchmarkSet(b, lockSetConstructor, true)
}
