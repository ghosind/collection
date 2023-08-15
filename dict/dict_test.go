package dict

import (
	"math/rand"

	"github.com/ghosind/collection"
	"github.com/ghosind/collection/set"
	"github.com/ghosind/go-assert"
)

func testDictionaryPut(a *assert.Assertion, m collection.Dictionary[int, int], data []int) {
	for i := 0; i < len(data)/2; i++ {
		m.Put(i, data[i])
	}
	a.EqualNow(m.Size(), len(data)/2)

	for i, v := range data {
		old := m.Put(i, v)
		expect := 0
		if i < len(data)/2 {
			expect = data[i]
		}
		a.EqualNow(old, expect)
	}
	a.EqualNow(m.Size(), len(data))
}

func testDictionaryGet(a *assert.Assertion, m collection.Dictionary[int, int], data []int) {
	for i := 0; i < len(data)*2; i++ {
		v, ok := m.Get(i)
		if i < len(data) {
			a.TrueNow(ok)

			a.EqualNow(v, data[i])
		} else {
			a.NotTrueNow(ok)
		}
	}

	for i := 0; i < len(data)*2; i++ {
		v := m.GetDefault(i, i+1)
		if i < len(data) {
			a.EqualNow(v, data[i])
		} else {
			a.EqualNow(v, i+1)
		}
	}
}

func testDictionaryContains(a *assert.Assertion, m collection.Dictionary[int, int], data []int) {
	for i := 0; i < len(data)*2; i++ {
		isContains := m.ContainsKey(i)
		if i < len(data) {
			a.TrueNow(isContains)
		} else {
			a.NotTrueNow(isContains)
		}
	}
}

func testDictionaryForEach(a *assert.Assertion, m collection.Dictionary[int, int], data []int) {
	n := 0
	err := m.ForEach(func(k, v int) error {
		n++
		a.EqualNow(v, data[k])

		return nil
	})
	a.NilNow(err)

	a.EqualNow(n, m.Size())
}

func testDictionaryRemove(a *assert.Assertion, m collection.Dictionary[int, int], data []int) {
	for i := 0; i < len(data)*2; i += 2 {
		old := m.Remove(i)
		if i < len(data) {
			a.EqualNow(old, data[i])
		} else {
			a.EqualNow(old, 0)
		}
	}
}

func testDictionaryClear(a *assert.Assertion, m collection.Dictionary[int, int]) {
	a.NotTrueNow(m.IsEmpty())

	m.Clear()

	a.TrueNow(m.IsEmpty())
}

func testDictionaryKeys(a *assert.Assertion, m collection.Dictionary[int, int]) {
	keys := m.Keys()
	a.EqualNow(len(keys), m.Size())

	for _, k := range keys {
		a.TrueNow(m.ContainsKey(k))
	}
}

func testDictionaryValues(a *assert.Assertion, m collection.Dictionary[int, int]) {
	vals := m.Values()
	a.EqualNow(len(vals), m.Size())

	valSet := set.NewHashSet[int]()
	valSet.AddAll(vals...)

	m.ForEach(func(_, v int) error {
		a.TrueNow(valSet.Contains(v))
		return nil
	})
}

func testDictionaryReplace(a *assert.Assertion, m collection.Dictionary[int, int], data []int) {
	for i := 0; i < len(data)/2; i++ {
		m.Put(i, data[i])
	}

	for i := 0; i < len(data); i++ {
		old, ok := m.Replace(i, data[i]+1)
		if i < len(data)/2 {
			a.TrueNow(ok)

			a.EqualNow(old, data[i])
		} else {
			a.NotTrueNow(ok)
		}
	}

	for i := 0; i < len(data); i++ {
		v := m.GetDefault(i, 0)
		if i < len(data)/2 {
			a.EqualNow(v, data[i]+1)
		} else {
			a.EqualNow(v, 0)
		}
	}
}

func testDictionary(a *assert.Assertion, m collection.Dictionary[int, int]) {
	data := rand.Perm(10)

	a.TrueNow(m.IsEmpty(), "Dictionary.IsEmpty() return false, expect true")

	testDictionaryPut(a, m, data)

	testDictionaryGet(a, m, data)

	testDictionaryContains(a, m, data)

	testDictionaryForEach(a, m, data)

	testDictionaryKeys(a, m)

	testDictionaryValues(a, m)

	testDictionaryRemove(a, m, data)

	testDictionaryClear(a, m)

	testDictionaryReplace(a, m, data)
}
