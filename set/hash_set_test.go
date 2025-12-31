package set

import (
	"testing"

	"github.com/ghosind/collection"
	"github.com/ghosind/go-assert"
)

var hashSetConstructor = func(initData ...[]int) collection.Set[int] {
	if len(initData) > 0 && len(initData[0]) > 0 {
		return NewHashSetFrom(initData[0]...)
	}
	return NewHashSet[int]()
}

func TestHashSet(t *testing.T) {
	a := assert.New(t)

	testSet(a, hashSetConstructor)
}

func BenchmarkHashSet(b *testing.B) {
	benchmarkSet(b, hashSetConstructor, false)
}
