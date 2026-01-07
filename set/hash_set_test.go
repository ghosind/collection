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

func BenchmarkHashSet_Add(b *testing.B) {
	benchmarkSet_Add(b, hashSetConstructor, false)
}

func BenchmarkHashSet_Contains(b *testing.B) {
	benchmarkSet_Contains(b, hashSetConstructor, false)
}

var lockHashSetConstructor = func(initData ...[]int) collection.Set[int] {
	if len(initData) > 0 && len(initData[0]) > 0 {
		return NewLockSet[int](NewHashSetFrom(initData[0]...))
	}
	return NewLockSet[int](NewHashSet[int]())
}

func TestLockHashSet(t *testing.T) {
	a := assert.New(t)

	testSet(a, lockHashSetConstructor)
}

func BenchmarkLockHashSet_Add(b *testing.B) {
	benchmarkSet_Add(b, lockHashSetConstructor, false)
}

func BenchmarkLockHashSet_Contains(b *testing.B) {
	benchmarkSet_Contains(b, lockHashSetConstructor, false)
}
