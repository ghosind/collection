package list

import (
	"github.com/ghosind/collection"
	"github.com/ghosind/go-assert"
)

type listConstructor func() collection.List[int]

var testData []int = []int{1, 2, 3, 4, 5}

func testList(a *assert.Assertion, constructor listConstructor) {
	testListAdd(a, constructor)
	testListAddAll(a, constructor)
	testListAddAtIndex(a, constructor)
	testListClear(a, constructor)
	testListClone(a, constructor)
	testListContains(a, constructor)
	testListContainsAll(a, constructor)
	testListEquals(a, constructor)
	testListGet(a, constructor)
	testListIndexOf(a, constructor)
	testListIsEmpty(a, constructor)
	testListIter(a, constructor)
	testListLastIndexOf(a, constructor)
	testListRemove(a, constructor)
	testListRemoveAll(a, constructor)
	testListRemoveAtIndex(a, constructor)
	testListRemoveIf(a, constructor)
	testListRetainAll(a, constructor)
	testListSet(a, constructor)
	testListSize(a, constructor)
	testListToSlice(a, constructor)
}

func testListAdd(a *assert.Assertion, constructor listConstructor) {
	l := constructor()

	a.Equal(0, l.Size())
	for _, v := range testData {
		a.True(l.Add(v))
	}
	a.EqualNow(len(testData), l.Size())
	a.EqualNow(testData, l.ToSlice())
}

func testListAddAll(a *assert.Assertion, constructor listConstructor) {
	l := constructor()

	a.EqualNow(0, l.Size())
	a.TrueNow(l.AddAll(testData...))
	a.EqualNow(len(testData), l.Size())
	a.EqualNow(testData, l.ToSlice())
}

func testListAddAtIndex(a *assert.Assertion, constructor listConstructor) {
	l := constructor()

	l.AddAll(testData...)
	a.EqualNow(len(testData), l.Size())
	a.EqualNow(testData, l.ToSlice())
	l.AddAtIndex(0, 0)
	a.EqualNow(len(testData)+1, l.Size())
	a.EqualNow([]int{0, 1, 2, 3, 4, 5}, l.ToSlice())

	l.AddAtIndex(l.Size(), 6)
	a.EqualNow(len(testData)+2, l.Size())
	a.EqualNow([]int{0, 1, 2, 3, 4, 5, 6}, l.ToSlice())

	a.PanicOfNow(func() { l.AddAtIndex(-1, 100) }, collection.ErrOutOfBounds)
}

func testListClear(a *assert.Assertion, constructor listConstructor) {
	l := constructor()

	l.AddAll(testData...)
	a.EqualNow(len(testData), l.Size())
	l.Clear()
	a.EqualNow(0, l.Size())
	a.EqualNow([]int{}, l.ToSlice())
}

func testListClone(a *assert.Assertion, constructor listConstructor) {
	l := constructor()

	l.AddAll(testData...)
	clone := l.Clone()
	a.EqualNow(l.Size(), clone.Size())
	a.EqualNow(l.ToSlice(), clone.ToSlice())

	clone.Add(6)
	a.NotEqualNow(l.Size(), clone.Size())
	a.NotEqualNow(l.ToSlice(), clone.ToSlice())
}

func testListContains(a *assert.Assertion, constructor listConstructor) {
	l := constructor()

	l.AddAll(testData...)
	for _, v := range testData {
		a.TrueNow(l.Contains(v))
	}
	a.NotTrueNow(l.Contains(100))
}

func testListContainsAll(a *assert.Assertion, constructor listConstructor) {
	l := constructor()

	l.AddAll(testData...)
	a.TrueNow(l.ContainsAll(1, 2, 3))
	a.NotTrueNow(l.ContainsAll(1, 2, 100))
}

func testListEquals(a *assert.Assertion, constructor listConstructor) {
	l1 := constructor()
	l2 := constructor()

	a.TrueNow(l1.Equals(l2))

	l1.AddAll(testData...)
	a.NotTrueNow(l1.Equals(l2))

	l2.AddAll(testData...)
	a.TrueNow(l1.Equals(l2))

	l2.Add(6)
	a.NotTrueNow(l1.Equals(l2))

	l2.Clear()
	for _, v := range testData {
		l2.Add(v + 1)
	}
	a.NotTrueNow(l1.Equals(l2))

	a.NotTrueNow(l1.Equals(nil))
	a.NotTrueNow(l1.Equals("string"))
}

func testListGet(a *assert.Assertion, constructor listConstructor) {
	l := constructor()

	l.AddAll(testData...)
	for i, v := range testData {
		a.EqualNow(v, l.Get(i))
	}
	a.PanicOfNow(func() { l.Get(-1) }, collection.ErrOutOfBounds)
	a.PanicOfNow(func() { l.Get(l.Size()) }, collection.ErrOutOfBounds)
}

func testListIndexOf(a *assert.Assertion, constructor listConstructor) {
	l := constructor()

	l.AddAll([]int{1, 2, 3, 2, 1}...)
	a.EqualNow(0, l.IndexOf(1))
	a.EqualNow(1, l.IndexOf(2))
	a.EqualNow(2, l.IndexOf(3))
	a.EqualNow(-1, l.IndexOf(100))
}

func testListIsEmpty(a *assert.Assertion, constructor listConstructor) {
	l := constructor()

	a.TrueNow(l.IsEmpty())
	l.AddAll(testData...)
	a.NotTrueNow(l.IsEmpty())
	l.Clear()
	a.TrueNow(l.IsEmpty())
}

func testListLastIndexOf(a *assert.Assertion, constructor listConstructor) {
	l := constructor()

	l.AddAll([]int{1, 2, 3, 2, 1}...)
	a.EqualNow(4, l.LastIndexOf(1))
	a.EqualNow(3, l.LastIndexOf(2))
	a.EqualNow(2, l.LastIndexOf(3))
	a.EqualNow(-1, l.LastIndexOf(100))
}

func testListRemove(a *assert.Assertion, constructor listConstructor) {
	l := constructor()

	l.AddAll(testData...)
	l.Add(3)
	a.EqualNow(len(testData)+1, l.Size())
	a.TrueNow(l.Remove(3))
	a.EqualNow(len(testData)-1, l.Size())
	a.EqualNow([]int{1, 2, 4, 5}, l.ToSlice())
	a.NotTrueNow(l.Remove(100))
	a.EqualNow(len(testData)-1, l.Size())
	a.EqualNow([]int{1, 2, 4, 5}, l.ToSlice())
}

func testListRemoveAll(a *assert.Assertion, constructor listConstructor) {
	l := constructor()

	l.AddAll(testData...)
	l.Add(3)
	a.EqualNow(len(testData)+1, l.Size())
	a.TrueNow(l.RemoveAll(3, 4, 5, 100))
	a.EqualNow(len(testData)-3, l.Size())
	a.EqualNow([]int{1, 2}, l.ToSlice())
	a.NotTrueNow(l.RemoveAll(100, 200))
	a.EqualNow(len(testData)-3, l.Size())
	a.EqualNow([]int{1, 2}, l.ToSlice())
}

func testListRemoveAtIndex(a *assert.Assertion, constructor listConstructor) {
	l := constructor()

	l.AddAll(testData...)
	a.EqualNow(len(testData), l.Size())
	v := l.RemoveAtIndex(0)
	a.EqualNow(1, v)
	a.EqualNow(len(testData)-1, l.Size())
	a.EqualNow([]int{2, 3, 4, 5}, l.ToSlice())
}

func testListRemoveIf(a *assert.Assertion, constructor listConstructor) {
	l := constructor()

	l.AddAll(testData...)
	a.EqualNow(len(testData), l.Size())
	a.TrueNow(l.RemoveIf(func(i int) bool { return i%2 == 0 }))
	a.EqualNow(3, l.Size())
	a.EqualNow([]int{1, 3, 5}, l.ToSlice())
	a.NotTrueNow(l.RemoveIf(func(i int) bool { return i > 10 }))
	a.EqualNow(3, l.Size())
	a.EqualNow([]int{1, 3, 5}, l.ToSlice())
}

func testListRetainAll(a *assert.Assertion, constructor listConstructor) {
	l := constructor()

	l.AddAll(testData...)
	a.EqualNow(len(testData), l.Size())
	a.TrueNow(l.RetainAll(2, 3, 4, 100))
	a.EqualNow(3, l.Size())
	a.EqualNow([]int{2, 3, 4}, l.ToSlice())
	a.TrueNow(l.RetainAll(100, 200))
	a.EqualNow(0, l.Size())
	a.EqualNow([]int{}, l.ToSlice())
}

func testListSet(a *assert.Assertion, constructor listConstructor) {
	l := constructor()

	l.AddAll(testData...)
	a.EqualNow(len(testData), l.Size())
	old := l.Set(0, 100)
	a.EqualNow(1, old)
	a.EqualNow(100, l.Get(0))
	a.EqualNow(len(testData), l.Size())

	old = l.Set(l.Size(), 200)
	a.EqualNow(0, old)
	a.EqualNow(200, l.Get(l.Size()-1))
	a.EqualNow(len(testData)+1, l.Size())

	a.PanicOfNow(func() { l.Set(-1, 300) }, collection.ErrOutOfBounds)
}

func testListSize(a *assert.Assertion, constructor listConstructor) {
	l := constructor()

	a.EqualNow(0, l.Size())
	l.AddAll(testData...)
	a.EqualNow(len(testData), l.Size())
}

func testListToSlice(a *assert.Assertion, constructor listConstructor) {
	l := constructor()

	a.EqualNow([]int{}, l.ToSlice())
	l.AddAll(testData...)
	a.EqualNow(testData, l.ToSlice())
}
