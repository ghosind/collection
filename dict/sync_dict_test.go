package dict

import (
	"errors"
	"math/rand"
	"testing"

	"github.com/ghosind/go-assert"
	"github.com/ghosind/utils"
)

func TestSyncDict(t *testing.T) {
	a := assert.New(t)

	testDict(a, NewSyncDict[int, int]())
}

func TestSyncDictCloneAndEquals(t *testing.T) {
	a := assert.New(t)

	data := rand.Perm(10)

	map1 := NewSyncDict[int, int]()
	testDictPut(a, map1, data)

	a.NotTrueNow(map1.Equals(1))

	map2 := NewSyncDict[string, int]()
	a.NotTrueNow(map1.Equals(map2))

	map3 := NewSyncDict[int, int]()
	map3.Put(1, 1)
	a.NotTrueNow(map1.Equals(map3))

	map4 := map1.Clone()
	a.TrueNow(map1.Equals(map4))

	map5 := NewSyncDict[int, int]()
	for i := 0; i < 10; i++ {
		map5.Put(i, i+1)
	}
	a.NotTrueNow(map1.Equals(map5))
}

func TestSyncDictForEach(t *testing.T) {
	a := assert.New(t)

	dict := NewSyncDict[int, int]()

	dict.Put(1, 1)
	dict.Put(2, 2)
	dict.Put(3, 3)
	dict.Put(5, 5)
	err := dict.ForEach(func(k, v int) error {
		return utils.Conditional(k == 5, errors.New("some error"), nil)
	})
	a.NotNilNow(err)
}
