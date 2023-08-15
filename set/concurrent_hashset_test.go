package set

import (
	"errors"
	"math/rand"
	"sync"
	"testing"

	"github.com/ghosind/go-assert"
	"github.com/ghosind/utils"
)

func TestConcurrentHashSet(t *testing.T) {
	a := assert.New(t)

	testSet[int](a, NewConcurrentHashSet[int](), intData)
	testSet[string](a, NewConcurrentHashSet[string](), strData)
	testSet[testStruct](a, NewConcurrentHashSet[testStruct](), structData)
	testSet[*testStruct](a, NewConcurrentHashSet[*testStruct](), pointerData)
}

func TestConcurrentHashSetCloneAndEquals(t *testing.T) {
	a := assert.New(t)

	set1 := NewConcurrentHashSet[int]()
	set1.Add(1)
	set1.Add(2)
	set1.Add(3)

	a.NotTrueNow(set1.Equals(1))

	set2 := NewConcurrentHashSet[string]()
	a.NotTrueNow(set1.Equals(set2))

	set3 := NewConcurrentHashSet[int]()
	set3.Add(1)
	a.NotTrueNow(set1.Equals(set3))

	set4 := set1.Clone()
	a.TrueNow(set1.Equals(set4))

	set5 := NewConcurrentHashSet[int]()
	set5.Add(1)
	set5.Add(2)
	set5.Add(4)
	a.NotTrueNow(set1.Equals(set5))
}

func TestConcurrentHashSetForEach(t *testing.T) {
	a := assert.New(t)

	set := NewConcurrentHashSet[int]()

	testSetForEachAndIter(a, set)

	set.Add(5)
	err := set.ForEach(func(e int) error {
		return utils.Conditional(e == 5, errors.New("some error"), nil)
	})
	a.NotNilNow(err)
}

func TestConcurrentHashSetWithConcurrent(t *testing.T) {
	a := assert.New(t)

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

	a.EqualNow(set.Size(), len(vals))

	wg.Add(n)
	for _, val := range vals {
		go func(v int) {
			defer wg.Done()
			a.TrueNow(set.Contains(v))
		}(val)
	}
	wg.Wait()
}
