package set

import (
	"testing"

	"github.com/ghosind/collection"
	"github.com/ghosind/go-assert"
)

var syncSetConstructor = func(initData ...[]int) collection.Set[int] {
	if len(initData) > 0 && len(initData[0]) > 0 {
		return NewSyncSetFrom(initData[0]...)
	}
	return NewSyncSet[int]()
}

func TestSyncSet(t *testing.T) {
	a := assert.New(t)

	testSet(a, syncSetConstructor)
}

func BenchmarkSyncSet_Add(b *testing.B) {
	benchmarkSet_Add(b, syncSetConstructor, true)
}

func BenchmarkSyncSet_Contains(b *testing.B) {
	benchmarkSet_Contains(b, syncSetConstructor, true)
}
