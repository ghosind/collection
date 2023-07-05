package dict

import (
	"math/rand"
	"testing"
)

func TestHashDictionary(t *testing.T) {
	testDictionary(t, NewHashDictionary[int, int]())
}

func TestHashDictionaryCloneAndEquals(t *testing.T) {
	data := rand.Perm(10)
	m := NewHashDictionary[int, int]()

	testDictionaryPut(t, m, data)

	newDictionary := m.Clone()
	if !m.Equals(newDictionary) {
		t.Errorf("HashDictionary.Equals(newDictionary) returns false, expect true")
	}

	newDictionary.ForEach(func(k, v int) error {
		newDictionary.Put(k, v+1)
		return nil
	})
	if m.Equals(newDictionary) {
		t.Errorf("HashDictionary.Equals(1) returns true, expect false")
	}

	newDictionary.Clear()
	if m.Equals(newDictionary) {
		t.Errorf("HashDictionary.Equals(newDictionary) returns true, expect false")
	}

	if m.Equals(NewHashDictionary[string, int]()) {
		t.Errorf("HashDictionary.Equals(NewHashDictionary[string, int]()) returns true, expect false")
	}

	if m.Equals(1) {
		t.Errorf("HashDictionary.Equals(1) returns true, expect false")
	}
}
