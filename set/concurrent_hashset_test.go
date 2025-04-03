package set

import (
	"math/rand"
	"sync"
	"testing"

	"github.com/ghosind/collection"
	"github.com/ghosind/go-assert"
)

func TestConcurrentHashSet(t *testing.T) {
	a := assert.New(t)
	constructor := func() collection.Set[int] {
		return NewConcurrentHashSet[int]()
	}

	testSet(a, constructor)
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
