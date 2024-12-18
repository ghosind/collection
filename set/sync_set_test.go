package set

import (
	"errors"
	"testing"

	"github.com/ghosind/go-assert"
	"github.com/ghosind/utils"
)

func TestSyncSet(t *testing.T) {
	a := assert.New(t)

	testSet[int](a, NewSyncSet[int](), intData)
	testSet[string](a, NewSyncSet[string](), strData)
	testSet[testStruct](a, NewSyncSet[testStruct](), structData)
	testSet[*testStruct](a, NewSyncSet[*testStruct](), pointerData)
}

func TestSyncSetCloneAndEquals(t *testing.T) {
	a := assert.New(t)

	set1 := NewSyncSet[int]()
	set1.Add(1)
	set1.Add(2)
	set1.Add(3)

	a.NotTrueNow(set1.Equals(1))

	set2 := NewSyncSet[string]()
	a.NotTrueNow(set1.Equals(set2))

	set3 := NewSyncSet[int]()
	set3.Add(1)
	a.NotTrueNow(set1.Equals(set3))

	set4 := set1.Clone()
	a.TrueNow(set1.Equals(set4))

	set5 := NewSyncSet[int]()
	set5.Add(1)
	set5.Add(2)
	set5.Add(4)
	a.NotTrueNow(set1.Equals(set5))
}

func TestSyncSetForEach(t *testing.T) {
	a := assert.New(t)

	set := NewSyncSet[int]()

	testSetForEachAndIter(a, set)

	set.Add(5)
	err := set.ForEach(func(e int) error {
		return utils.Conditional(e == 5, errors.New("some error"), nil)
	})
	a.NotNilNow(err)
}
