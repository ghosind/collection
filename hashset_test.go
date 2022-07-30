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
	set.Add(1)
	set.Add(2)
	set.Add(3)

	records := map[int]int{}

	if err := set.ForEach(func(e int) error {
		records[e]++
		return nil
	}); err != nil {
		t.Errorf("set.ForEach returns error (%v), expect no error", err)
	}

	if len(records) != set.Size() {
		t.Errorf("len(records) is %d, expect %d", len(records), set.Size())
	}
	for k, v := range records {
		if v != 1 {
			t.Errorf("records[%d] is %d, expect 1", k, v)
		}
	}

	set.Add(5)
	if err := set.ForEach(func(e int) error {
		return utils.Conditional(e == 5, errors.New("some error"), nil)
	}); err == nil {
		t.Error("set.ForEach returns no error, expect \"some error\"")
	}
}
