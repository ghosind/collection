package set

import (
	"errors"
	"math/rand"
	"sync"
	"testing"

	"github.com/ghosind/utils"
)

func TestConcurrentHashSet(t *testing.T) {
	testSet[int](t, NewConcurrentHashSet[int](), intData)
	testSet[string](t, NewConcurrentHashSet[string](), strData)
	testSet[testStruct](t, NewConcurrentHashSet[testStruct](), structData)
	testSet[*testStruct](t, NewConcurrentHashSet[*testStruct](), pointerData)
}

func TestConcurrentHashSetCloneAndEquals(t *testing.T) {
	set1 := NewConcurrentHashSet[int]()
	set1.Add(1)
	set1.Add(2)
	set1.Add(3)

	if set1.Equals(1) {
		t.Errorf("set1.Equals(1) return true, expect false")
	}

	set2 := NewConcurrentHashSet[string]()
	if set1.Equals(set2) {
		t.Errorf("set1.Equals(set2) return true, expect false")
	}

	set3 := NewConcurrentHashSet[int]()
	set3.Add(1)
	if set1.Equals(set3) {
		t.Errorf("set1.Equals(set3) return true, expect false")
	}

	set4 := set1.Clone()
	if !set1.Equals(set4) {
		t.Errorf("set1.Equals(set4) return false, expect true")
	}

	set5 := NewConcurrentHashSet[int]()
	set5.Add(1)
	set5.Add(2)
	set5.Add(4)
	if set1.Equals(set5) {
		t.Errorf("set1.Equals(set5) return true, expect false")
	}
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

func TestConcurrentHashSetWithConcurrent(t *testing.T) {
	set := NewConcurrentHashSet[int]()
	n := 100

	vals := rand.Perm(n)
	wg := sync.WaitGroup{}

	wg.Add(n)
	for _, val := range vals {
		go func(v int) {
			defer wg.Done()
			set.Add(v)
		}(val)
	}
	wg.Wait()

	if set.Size() != len(vals) {
		t.Errorf("set.Size() should be %d, but %d", len(vals), set.Size())
	}

	wg.Add(n)
	for _, val := range vals {
		go func(v int) {
			defer wg.Done()
			if !set.Contains(v) {
				t.Errorf("set should contains %d", v)
			}
		}(val)
	}
	wg.Wait()
}
