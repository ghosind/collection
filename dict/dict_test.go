package dict

import (
	"math/rand"
	"testing"

	"github.com/ghosind/collection"
	"github.com/ghosind/collection/set"
)

func testDictionaryPut(t *testing.T, m collection.Dictionary[int, int], data []int) {
	for i := 0; i < len(data)/2; i++ {
		m.Put(i, data[i])
	}
	if m.Size() != len(data)/2 {
		t.Errorf("Dictionary.Size() returns %d, expect %d", m.Size(), len(data)/2)
	}

	for i, v := range data {
		old := m.Put(i, v)
		expect := 0
		if i < len(data)/2 {
			expect = data[i]
		}
		if old != expect {
			t.Errorf("Dictionary.Put returns %d, expect %d", old, expect)
		}
	}
	if m.Size() != len(data) {
		t.Errorf("Dictionary.Size() returns %d, expect %d", m.Size(), len(data))
	}
}

func testDictionaryGet(t *testing.T, m collection.Dictionary[int, int], data []int) {
	for i := 0; i < len(data)*2; i++ {
		v, ok := m.Get(i)
		if i < len(data) {
			if !ok {
				t.Errorf("Dictionary not contains key %d, expect contains", i)
			}
			if v != data[i] {
				t.Errorf("Dictionary.Get(%d) returns %d, expect %d", i, v, data[i])
			}
		} else {
			if ok {
				t.Errorf("Dictionary contains key %d, expect not", i)
			}
		}
	}

	for i := 0; i < len(data)*2; i++ {
		v := m.GetDefault(i, i+1)
		if i < len(data) {
			if v != data[i] {
				t.Errorf("Dictionary.GetDefault(%d) returns %d, expect %d", i, v, data[i])
			}
		} else {
			if v != i+1 {
				t.Errorf("Dictionary.GetDefault(%d) returns %d, expect %d", i, v, i+1)
			}
		}
	}
}

func testDictionaryContains(t *testing.T, m collection.Dictionary[int, int], data []int) {
	for i := 0; i < len(data)*2; i++ {
		isContains := m.ContainsKey(i)
		if i < len(data) {
			if !isContains {
				t.Errorf("Dictionary.Contains(%d) return false, expect true", i)
			}
		} else {
			if isContains {
				t.Errorf("Dictionary.Contains(%d) return true, expect false", i)
			}
		}
	}
}

func testDictionaryForEach(t *testing.T, m collection.Dictionary[int, int], data []int) {
	n := 0
	err := m.ForEach(func(k, v int) error {
		n++
		if data[k] != v {
			t.Errorf("")
		}

		return nil
	})
	if err != nil {
		t.Errorf("Dictionary.ForEach returns %v, expect nil", err)
	}

	if n != m.Size() {
		t.Errorf("Dictionary.ForEach run handler %d times, expect %d", n, m.Size())
	}
}

func testDictionaryRemove(t *testing.T, m collection.Dictionary[int, int], data []int) {
	for i := 0; i < len(data)*2; i += 2 {
		old := m.Remove(i)
		if i < len(data) {
			if old != data[i] {
				t.Errorf("Dictionary.Remove(%d) returns %d, expect %d", i, old, data[i]+1)
			}
		} else if old != 0 {
			t.Errorf("Dictionary.Remove(%d) returns %d, expect 0", i, old)
		}
	}
}

func testDictionaryClear(t *testing.T, m collection.Dictionary[int, int]) {
	if m.IsEmpty() {
		t.Error("Dictionary.IsEmpty() return true, expect false")
	}

	m.Clear()

	if !m.IsEmpty() {
		t.Error("Dictionary.IsEmpty() return false, expect true")
	}
}

func testDictionaryKeys(t *testing.T, m collection.Dictionary[int, int]) {
	keys := m.Keys()
	if len(keys) != m.Size() {
		t.Errorf("Dictionary.Keys() return an array contains %d element, expect %d", len(keys), m.Size())
	}

	for _, k := range keys {
		if !m.ContainsKey(k) {
			t.Errorf("key %d not in Dictionary", k)
		}
	}
}

func testDictionaryValues(t *testing.T, m collection.Dictionary[int, int]) {
	vals := m.Values()
	if len(vals) != m.Size() {
		t.Errorf("Dictionary.Values() return an array contains %d element, expect %d", len(vals), m.Size())
	}

	valSet := set.NewHashSet[int]()
	valSet.AddAll(vals...)

	m.ForEach(func(_, v int) error {
		if !valSet.Contains(v) {
			t.Errorf("Dictionary.Values() not contains %dd", v)
		}
		return nil
	})
}

func testDictionaryReplace(t *testing.T, m collection.Dictionary[int, int], data []int) {
	for i := 0; i < len(data)/2; i++ {
		m.Put(i, data[i])
	}

	for i := 0; i < len(data); i++ {
		old, ok := m.Replace(i, data[i]+1)
		if i < len(data)/2 {
			if !ok {
				t.Errorf("Dictionary.Replace(%d, ?) no old value found, expect return old value", i)
			} else if old != data[i] {
				t.Errorf("Dictionary.Replace(%d, ?) return old value %d, expect %d", i, old, data[i])
			}
		} else if ok {
			t.Errorf("Dictionary.Replace(%d, ?) found old value, expect no old value", i)
		}
	}

	for i := 0; i < len(data); i++ {
		v := m.GetDefault(i, 0)
		if i < len(data)/2 {
			if v != data[i]+1 {
				t.Errorf("Dictionary.Get(%d) return %d, expect %d", i, v, data[i]+1)
			}
		} else if v != 0 {
			t.Errorf("Dictionary.Get(%d) return %d, expect 0", i, v)
		}
	}
}

func testDictionary(t *testing.T, m collection.Dictionary[int, int]) {
	data := rand.Perm(10)

	if !m.IsEmpty() {
		t.Error("Dictionary.IsEmpty() return false, expect true")
	}

	testDictionaryPut(t, m, data)

	testDictionaryGet(t, m, data)

	testDictionaryContains(t, m, data)

	testDictionaryForEach(t, m, data)

	testDictionaryKeys(t, m)

	testDictionaryValues(t, m)

	testDictionaryRemove(t, m, data)

	testDictionaryClear(t, m)

	testDictionaryReplace(t, m, data)
}
