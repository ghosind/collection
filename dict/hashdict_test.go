package dict

import (
	"errors"
	"math/rand"
	"testing"

	"github.com/ghosind/go-assert"
	"github.com/ghosind/utils"
)

func TestHashDictionary(t *testing.T) {
	a := assert.New(t)

	testDict(a, NewHashDictionary[int, int]())
}

func TestHashDictionaryCloneAndEquals(t *testing.T) {
	a := assert.New(t)

	data := rand.Perm(10)
	m := NewHashDictionary[int, int]()

	testDictPut(a, m, data)

	newDictionary := m.Clone()
	a.TrueNow(m.Equals(newDictionary))

	newDictionary.ForEach(func(k, v int) error {
		newDictionary.Put(k, v+1)
		return nil
	})
	a.NotTrueNow(m.Equals(newDictionary))

	newDictionary.Clear()
	a.NotTrueNow(m.Equals(newDictionary))
	a.NotTrueNow(m.Equals(NewHashDictionary[string, int]()))
	a.NotTrueNow(m.Equals(1))
}

func TestHashDictionaryForEach(t *testing.T) {
	a := assert.New(t)

	dict := NewHashDictionary[int, int]()

	dict.Put(1, 1)
	dict.Put(2, 2)
	dict.Put(3, 3)
	dict.Put(5, 5)
	err := dict.ForEach(func(k, v int) error {
		return utils.Conditional(k == 5, errors.New("some error"), nil)
	})
	a.NotNilNow(err)
}
