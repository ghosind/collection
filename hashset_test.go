package collection

import (
	"errors"
	"testing"

	"github.com/ghosind/utils"
)

func TestHashSet(t *testing.T) {
	testSet[int](t, NewHashSet[int](), intData)
	testSet[string](t, NewHashSet[string](), strData)
	testSet[testStruct](t, NewHashSet[testStruct](), structData)
	testSet[*testStruct](t, NewHashSet[*testStruct](), pointerData)
}

func TestHashSetCloneAndEquals(t *testing.T) {
	set1 := NewHashSet[int]()
	set1.Add(1)
	set1.Add(2)
	set1.Add(3)

	if set1.Equals(1) {
		t.Errorf("set1.Equals(1) return true, expect false")
	}

	set2 := NewHashSet[string]()
	if set1.Equals(set2) {
		t.Errorf("set1.Equals(set2) return true, expect false")
	}

	set3 := NewHashSet[int]()
	set3.Add(1)
	if set1.Equals(set3) {
		t.Errorf("set1.Equals(set3) return true, expect false")
	}

	set4 := set1.Clone()
	if !set1.Equals(set4) {
		t.Errorf("set1.Equals(set4) return false, expect true")
	}

	set5 := NewHashSet[int]()
	set5.Add(1)
	set5.Add(2)
	set5.Add(4)
	if set1.Equals(set5) {
		t.Errorf("set1.Equals(set5) return true, expect false")
	}
}

func TestHashSetForEach(t *testing.T) {
	set := NewHashSet[int]()

	testSetForEachAndIter(t, set)

	set.Add(5)
	if err := set.ForEach(func(e int) error {
		return utils.Conditional(e == 5, errors.New("some error"), nil)
	}); err == nil {
		t.Error("set.ForEach returns no error, expect \"some error\"")
	}
}
