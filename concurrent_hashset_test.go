package collection

import (
	"errors"
	"testing"

	"github.com/ghosind/utils"
)

func TestConcurrentHashSet(t *testing.T) {
	testSet[int](t, NewConcurrentHashSet[int](), intData)
	testSet[string](t, NewConcurrentHashSet[string](), strData)
	testSet[testStruct](t, NewConcurrentHashSet[testStruct](), structData)
	testSet[*testStruct](t, NewConcurrentHashSet[*testStruct](), pointerData)
}

func TestConcurrentHashSetForEach(t *testing.T) {
	set := NewConcurrentHashSet[int]()

	testSetForEachAndIter(t, set)

	set.Add(5)
	if err := set.ForEach(func(e int) error {
		return utils.Conditional(e == 5, errors.New("some error"), nil)
	}); err == nil {
		t.Error("set.ForEach returns no error, expect \"some error\"")
	}
}
