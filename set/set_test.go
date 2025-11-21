package set

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/ghosind/collection"
	"github.com/ghosind/go-assert"
)

type setConstructor func() collection.Set[int]

var testNums1 = []int{47, 11, 42, 13, 37, 23, 31, 29, 17, 19}
var testNums2 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func testSet(a *assert.Assertion, constructor setConstructor) {
	testSetAdd(a, constructor)
	testSetAddAll(a, constructor)
	testSetClear(a, constructor)
	testSetClone(a, constructor)
	testSetContains(a, constructor)
	testSetContainsAll(a, constructor)
	testSetEquals(a, constructor)
	testSetForEach(a, constructor)
	testSetIsEmpty(a, constructor)
	testSetIter(a, constructor)
	testSetRemove(a, constructor)
	testSetRemoveAll(a, constructor)
	testSetRemoveIf(a, constructor)
	testSetRetainAll(a, constructor)
	testSetSize(a, constructor)
	testSetString(a, constructor)
	testSetToSlice(a, constructor)
	testSetJSON(a, constructor)
}

func testSetAdd(a *assert.Assertion, constructor setConstructor) {
	set1 := constructor()
	last := 0
	for _, n := range testNums1 {
		set1.Add(n)
		a.TrueNow(set1.Contains(n))
		last = n
	}

	a.EqualNow(set1.Size(), len(testNums1))

	// Add should not add duplicates
	set1.Add(last)
	a.EqualNow(set1.Size(), len(testNums1))
}

func testSetAddAll(a *assert.Assertion, constructor setConstructor) {
	set1 := constructor()

	set1.AddAll(testNums1...)
	a.EqualNow(set1.Size(), len(testNums1))
	for _, n := range testNums1 {
		a.TrueNow(set1.Contains(n))
	}
}

func testSetClear(a *assert.Assertion, constructor setConstructor) {
	set1 := constructor()
	set1.AddAll(testNums1...)
	a.EqualNow(set1.Size(), len(testNums1))
	set1.Clear()
	a.EqualNow(set1.Size(), 0)
	a.NotTrueNow(set1.Contains(testNums1[len(testNums1)-1]))
}

func testSetClone(a *assert.Assertion, constructor setConstructor) {
	set1 := constructor()
	set1.AddAll(testNums1...)
	set2 := set1.Clone()
	a.NotEqualNow(set1, set2)
	a.TrueNow(set1.Equals(set2))
}

func testSetContains(a *assert.Assertion, constructor setConstructor) {
	set := constructor()
	set.AddAll(testNums1...)

	// Check that all elements are in the set
	for _, n := range testNums1 {
		a.TrueNow(set.Contains(n))
	}

	// Check that a non-existing element is not in the set
	a.NotTrueNow(set.Contains(0))
}

func testSetContainsAll(a *assert.Assertion, constructor setConstructor) {
	set := constructor()

	set.AddAll(testNums1...)
	a.TrueNow(set.ContainsAll(testNums1...))

	nums := append([]int{0}, testNums1...)
	a.NotTrueNow(set.ContainsAll(nums...))

	set.RemoveAll(nums...)
	a.NotTrueNow(set.ContainsAll(testNums1...))
}

func testSetEquals(a *assert.Assertion, constructor setConstructor) {
	set1 := constructor()
	a.NotTrueNow(set1.Equals(nil))

	set2 := constructor()
	set3 := constructor()

	set1.AddAll(testNums1...)
	set2.AddAll(testNums1...)
	a.TrueNow(set1.Equals(set2))

	set3.AddAll(testNums2...)
	a.NotTrueNow(set1.Equals(set3))

	set3.AddAll(testNums1...)
	a.NotTrueNow(set1.Equals(set3))
}

func testSetForEach(a *assert.Assertion, constructor setConstructor) {
	set1 := constructor()
	set1.AddAll(testNums1...)

	// Foreach should iterate over all elements
	set2 := constructor()
	err := set1.ForEach(func(e int) error {
		set2.Add(e)
		return nil
	})
	a.NilNow(err)
	a.TrueNow(set1.Equals(set2))

	// ForEach should exit early if the handler returns an error
	cnt := 0
	expectedErr := errors.New("expected error")
	err = set1.ForEach(func(e int) error {
		cnt++
		return expectedErr
	})
	a.EqualNow(err, expectedErr)
	a.EqualNow(cnt, 1)
}

func testSetIsEmpty(a *assert.Assertion, constructor setConstructor) {
	set := constructor()
	a.TrueNow(set.IsEmpty())

	set.AddAll(testNums1...)
	a.NotTrueNow(set.IsEmpty())

	set.Clear()
	a.TrueNow(set.IsEmpty())
}

func testSetRemove(a *assert.Assertion, constructor setConstructor) {
	set := constructor()
	set.AddAll(testNums1...)
	a.EqualNow(set.Size(), len(testNums1))

	// Remove non-existing element should not change the size
	ret := set.Remove(0)
	a.NotTrueNow(ret)
	a.EqualNow(set.Size(), len(testNums1))
	a.NotTrueNow(set.Contains(0))

	for _, n := range testNums1 {
		ret := set.Remove(n)
		a.TrueNow(ret)
		a.NotTrueNow(set.Contains(n))
	}

	a.EqualNow(set.Size(), 0)
}

func testSetRemoveAll(a *assert.Assertion, constructor setConstructor) {
	set := constructor()
	set.AddAll(testNums1...)
	a.EqualNow(set.Size(), len(testNums1))

	set.RemoveAll(testNums1...)
	a.EqualNow(set.Size(), 0)

	for _, n := range testNums1 {
		a.NotTrueNow(set.Contains(n))
	}
}

func testSetRemoveIf(a *assert.Assertion, constructor setConstructor) {
	set := constructor()
	set.AddAll(testNums1...)
	a.EqualNow(set.Size(), len(testNums1))

	// Remove all even numbers
	ret := set.RemoveIf(func(e int) bool {
		return e%2 == 0
	})
	a.EqualNow(ret, true)

	for _, n := range testNums1 {
		if n%2 == 0 {
			a.NotTrueNow(set.Contains(n))
		} else {
			a.TrueNow(set.Contains(n))
		}
	}
}

func testSetRetainAll(a *assert.Assertion, constructor setConstructor) {
	set := constructor()
	set.AddAll(testNums1...)
	a.EqualNow(set.Size(), len(testNums1))

	evenNums := make([]int, 0)
	for _, n := range testNums1 {
		if n%2 == 0 {
			evenNums = append(evenNums, n)
		}
	}

	// Retain all even numbers
	set.RetainAll(evenNums...)
	a.EqualNow(set.Size(), len(evenNums))

	for _, n := range testNums1 {
		if n%2 == 0 {
			a.TrueNow(set.Contains(n))
		} else {
			a.NotTrueNow(set.Contains(n))
		}
	}
}

func testSetSize(a *assert.Assertion, constructor setConstructor) {
	set := constructor()
	a.EqualNow(set.Size(), 0)

	set.AddAll(testNums1...)
	a.EqualNow(set.Size(), len(testNums1))

	set.Clear()
	a.EqualNow(set.Size(), 0)
}

func testSetString(a *assert.Assertion, constructor setConstructor) {
	set := constructor()
	set.AddAll(testNums1...)
	str := set.String()

	a.HasPrefixString(str, "set[")
	a.HasSuffixString(str, "]")

	for _, n := range testNums1 {
		a.ContainsString(str, strconv.FormatInt(int64(n), 10))
	}
}

func testSetToSlice(a *assert.Assertion, constructor setConstructor) {
	set := constructor()
	set.AddAll(testNums1...)
	slice := set.ToSlice()
	a.EqualNow(len(slice), len(testNums1))

	for _, n := range slice {
		a.TrueNow(set.Contains(n))
	}
}

func testSetJSON(a *assert.Assertion, constructor setConstructor) {
	s1 := constructor()
	s1.AddAll(testNums1...)

	b, err := s1.MarshalJSON()
	a.NilNow(err)

	s2 := constructor()
	err = s2.UnmarshalJSON(b)
	a.NilNow(err)
	a.TrueNow(s1.Equals(s2))

	s2.Clear()

	b, err = json.Marshal(s1)
	a.NilNow(err)

	err = json.Unmarshal(b, s2)
	a.NilNow(err)
	a.TrueNow(s1.Equals(s2))
}
