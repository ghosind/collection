package dict

import (
	"errors"
	"math/rand"
	"testing"

	"github.com/ghosind/go-assert"
	"github.com/ghosind/utils"
)

func TestHashDict(t *testing.T) {
	a := assert.New(t)

	testDict(a, NewHashDict[int, int]())
}

func TestHashDictCloneAndEquals(t *testing.T) {
	a := assert.New(t)

	data := rand.Perm(10)
	m := NewHashDict[int, int]()

	testDictPut(a, m, data)

	newDict := m.Clone()
	a.TrueNow(m.Equals(newDict))

	newDict.ForEach(func(k, v int) error {
		newDict.Put(k, v+1)
		return nil
	})
	a.NotTrueNow(m.Equals(newDict))

	newDict.Clear()
	a.NotTrueNow(m.Equals(newDict))
	a.NotTrueNow(m.Equals(NewHashDict[string, int]()))
	a.NotTrueNow(m.Equals(1))
}

func TestHashDictForEach(t *testing.T) {
	a := assert.New(t)

	dict := NewHashDict[int, int]()

	dict.Put(1, 1)
	dict.Put(2, 2)
	dict.Put(3, 3)
	dict.Put(5, 5)
	err := dict.ForEach(func(k, v int) error {
		return utils.Conditional(k == 5, errors.New("some error"), nil)
	})
	a.NotNilNow(err)
}
