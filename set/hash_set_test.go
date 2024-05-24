package set

import (
	"errors"
	"testing"

	"github.com/ghosind/go-assert"
	"github.com/ghosind/utils"
)

func TestHashSet(t *testing.T) {
	a := assert.New(t)

	testSet[int](a, NewHashSet[int](), intData)
	testSet[string](a, NewHashSet[string](), strData)
	testSet[testStruct](a, NewHashSet[testStruct](), structData)
	testSet[*testStruct](a, NewHashSet[*testStruct](), pointerData)
}

func TestHashSetCloneAndEquals(t *testing.T) {
	a := assert.New(t)

	set1 := NewHashSet[int]()
	set1.Add(1)
	set1.Add(2)
	set1.Add(3)

	a.NotTrueNow(set1.Equals(1))

	set2 := NewHashSet[string]()
	a.NotTrueNow(set1.Equals(set2))

	set3 := NewHashSet[int]()
	set3.Add(1)
	a.NotTrueNow(set1.Equals(set3))

	set4 := set1.Clone()
	a.TrueNow(set1.Equals(set4))

	set5 := NewHashSet[int]()
	set5.Add(1)
	set5.Add(2)
	set5.Add(4)
	a.NotTrueNow(set1.Equals(set5))
}

func TestHashSetForEach(t *testing.T) {
	a := assert.New(t)

	set := NewHashSet[int]()

	testSetForEachAndIter(a, set)

	set.Add(5)
	err := set.ForEach(func(e int) error {
		return utils.Conditional(e == 5, errors.New("some error"), nil)
	})
	a.NotNilNow(err)
}
