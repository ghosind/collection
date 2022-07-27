package collection

import "testing"

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
}
