package set

import (
	"github.com/ghosind/collection"
	"github.com/ghosind/go-assert"
)

type testStruct struct {
	v int
}

var intData = []int{1, 2, 3, 4, 5, 6, 7}
var strData = []string{"a", "b", "c", "d", "e", "f", "g"}
var structData = []testStruct{{1}, {2}, {3}, {4}, {5}, {6}, {7}}
var pointerData = []*testStruct{{1}, {2}, {3}, {4}, {5}, {6}, {7}}

func testSetAdd[T comparable](a *assert.Assertion, set collection.Set[T], e T) {
	a.TrueNow(set.Add(e))

	a.NotTrueNow(set.Add(e))
}

func testSetAddAll[T comparable](a *assert.Assertion, set collection.Set[T], c ...T) {
	a.TrueNow(set.AddAll(c...))

	a.NotTrueNow(set.AddAll(c...))
}

func testSetContains[T comparable](a *assert.Assertion, set collection.Set[T], data []T) {
	a.TrueNow(set.Contains(data[0]))

	a.NotTrueNow(set.Contains(data[len(data)-1]))

	a.TrueNow(set.ContainsAll(data[0 : len(data)-1]...))

	a.NotTrueNow(set.ContainsAll(data...))
}

func testSetToSlice[T comparable](a *assert.Assertion, set collection.Set[T]) {
	slice := set.ToSlice()
	a.EqualNow(len(slice), set.Size())

	for _, e := range slice {
		a.TrueNow(set.Contains(e))
	}
}

func testSetRemove[T comparable](a *assert.Assertion, set collection.Set[T], data []T) {
	a.NotTrueNow(set.IsEmpty())
	a.EqualNow(set.Size(), len(data)-1)

	a.TrueNow(set.Remove(data[0]))
	a.EqualNow(set.Size(), len(data)-2)

	a.NotTrueNow(set.Remove(data[len(data)-1]))
	a.EqualNow(set.Size(), len(data)-2)
}

func testSetRemoveAll[T comparable](a *assert.Assertion, set collection.Set[T], data []T) {
	a.TrueNow(set.RemoveAll(data[0:2]...))
	a.EqualNow(set.Size(), len(data)-3)
}

func testSetClear[T comparable](a *assert.Assertion, set collection.Set[T]) {
	set.Clear()
	a.TrueNow(set.IsEmpty())
}

func testSet[T comparable](a *assert.Assertion, set collection.Set[T], data []T) {
	a.NotNilNow(set)

	testSetAdd(a, set, data[0])
	testSetAddAll(a, set, data[0:len(data)-1]...)
	testSetContains(a, set, data)
	testSetToSlice(a, set)
	testSetRemove(a, set, data)
	testSetRemoveAll(a, set, data)
	testSetClear(a, set)
}

func testSetForEachAndIter(a *assert.Assertion, set collection.Set[int]) {
	set.Add(1)
	set.Add(2)
	set.Add(3)

	testSetIter(a, set)

	records := map[int]int{}

	err := set.ForEach(func(e int) error {
		records[e]++
		return nil
	})
	a.NilNow(err)

	a.Equal(len(records), set.Size())
	for _, v := range records {
		a.EqualNow(v, 1)
	}
}
