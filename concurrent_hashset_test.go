package collection

import "testing"

func TestConcurrentHashSet(t *testing.T) {
	testSet[int](t, NewConcurrentHashSet[int](), intData)
	testSet[string](t, NewConcurrentHashSet[string](), strData)
	testSet[testStruct](t, NewConcurrentHashSet[testStruct](), structData)
	testSet[*testStruct](t, NewConcurrentHashSet[*testStruct](), pointerData)
}
