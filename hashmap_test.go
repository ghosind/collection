package collection

import (
	"math/rand"
	"testing"
)

func TestHashMap(t *testing.T) {
	testMap(t, NewHashMap[int, int]())
}

func TestHashMapCloneAndEquals(t *testing.T) {
	data := rand.Perm(10)
	m := NewHashMap[int, int]()

	testMapPut(t, m, data)

	newMap := m.Clone()
	if !m.Equals(newMap) {
		t.Errorf("HashMap.Equals(newMap) returns false, expect true")
	}

	newMap.ForEach(func(k, v int) error {
		newMap.Put(k, v+1)
		return nil
	})
	if m.Equals(newMap) {
		t.Errorf("HashMap.Equals(1) returns true, expect false")
	}

	newMap.Clear()
	if m.Equals(newMap) {
		t.Errorf("HashMap.Equals(newMap) returns true, expect false")
	}

	if m.Equals(NewHashMap[string, int]()) {
		t.Errorf("HashMap.Equals(NewHashMap[string, int]()) returns true, expect false")
	}

	if m.Equals(1) {
		t.Errorf("HashMap.Equals(1) returns true, expect false")
	}
}
