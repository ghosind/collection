package set

import (
	"testing"

	"github.com/ghosind/collection"
	"github.com/ghosind/go-assert"
)

var syncSetConstructor = func() collection.Set[int] {
	return NewSyncSet[int]()
}

func TestSyncSet(t *testing.T) {
	a := assert.New(t)

	testSet(a, syncSetConstructor)
}

func BenchmarkSyncSet(b *testing.B) {
	benchmarkSet(b, syncSetConstructor, true)
}
