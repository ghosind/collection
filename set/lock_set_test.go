package set

import (
	"testing"

	"github.com/ghosind/collection"
	"github.com/ghosind/go-assert"
)

var lockSetConstructor = func() collection.Set[int] {
	return NewLockSet[int]()
}

func TestLockSet(t *testing.T) {
	a := assert.New(t)

	testSet(a, lockSetConstructor)
}

func BenchmarkLockSet(b *testing.B) {
	benchmarkSet(b, lockSetConstructor, true)
}
