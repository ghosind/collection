package collection

import "testing"

func TestHashSet(t *testing.T) {
	testSet[int](t, NewHashSet[int](), intData)
	testSet[string](t, NewHashSet[string](), strData)
	testSet[testStruct](t, NewHashSet[testStruct](), structData)
	testSet[*testStruct](t, NewHashSet[*testStruct](), pointerData)
}
