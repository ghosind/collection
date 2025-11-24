package set

import (
	"testing"

	"github.com/ghosind/collection"
	"github.com/ghosind/go-assert"
)

var hashSetConstructor = func() collection.Set[int] {
	return NewHashSet[int]()
}

func TestHashSet(t *testing.T) {
	a := assert.New(t)

	testSet(a, hashSetConstructor)
}

func BenchmarkHashSet(b *testing.B) {
	benchmarkSet(b, hashSetConstructor, false)
}
