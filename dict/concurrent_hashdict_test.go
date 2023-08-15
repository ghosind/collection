package dict

import (
	"errors"
	"math/rand"
	"testing"

	"github.com/ghosind/go-assert"
	"github.com/ghosind/utils"
)

func TestConcurrentHashDictionary(t *testing.T) {
	a := assert.New(t)

	testDictionary(a, NewConcurrentHashDictionary[int, int]())
}

func TestConcurrentHashDictionaryCloneAndEquals(t *testing.T) {
	a := assert.New(t)

	data := rand.Perm(10)

	map1 := NewConcurrentHashDictionary[int, int]()
	testDictionaryPut(a, map1, data)

	a.NotTrueNow(map1.Equals(1))

	map2 := NewConcurrentHashDictionary[string, int]()
	a.NotTrueNow(map1.Equals(map2))

	map3 := NewConcurrentHashDictionary[int, int]()
	map3.Put(1, 1)
	a.NotTrueNow(map1.Equals(map3))

	map4 := map1.Clone()
	a.TrueNow(map1.Equals(map4))

	map5 := NewConcurrentHashDictionary[int, int]()
	for i := 0; i < 10; i++ {
		map5.Put(i, i+1)
	}
	a.NotTrueNow(map1.Equals(map5))
}

func TestConcurrentHashDictionaryForEach(t *testing.T) {
	a := assert.New(t)

	dict := NewConcurrentHashDictionary[int, int]()

	dict.Put(1, 1)
	dict.Put(2, 2)
	dict.Put(3, 3)
	dict.Put(5, 5)
	err := dict.ForEach(func(k, v int) error {
		return utils.Conditional(k == 5, errors.New("some error"), nil)
	})
	a.NotNilNow(err)
}
