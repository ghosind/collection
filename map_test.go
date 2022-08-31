package collection

import (
	"math/rand"
	"testing"
)

func testMapPut(t *testing.T, m Map[int, int], data []int) {
	for i := 0; i < len(data)/2; i++ {
		m.Put(i, data[i])
	}
	if m.Size() != len(data)/2 {
		t.Errorf("HashMap.Size() returns %d, expect %d", m.Size(), len(data)/2)
	}

	for i, v := range data {
		old := m.Put(i, v)
		expect := 0
		if i < len(data)/2 {
			expect = data[i]
		}
		if old != expect {
			t.Errorf("HashMap.Put returns %d, expect %d", old, expect)
		}
	}
	if m.Size() != len(data) {
		t.Errorf("HashMap.Size() returns %d, expect %d", m.Size(), len(data))
	}
}

func testMapGet(t *testing.T, m Map[int, int], data []int) {
	for i := 0; i < len(data)*2; i++ {
		v, ok := m.Get(i)
		if i < len(data) {
			if !ok {
				t.Errorf("HashMap not contains key %d, expect contains", i)
			}
			if v != data[i] {
				t.Errorf("HashMap.Get(%d) returns %d, expect %d", i, v, data[i])
			}
		} else {
			if ok {
				t.Errorf("HashMap contains key %d, expect not", i)
			}
		}
	}
}

func testMapContains(t *testing.T, m Map[int, int], data []int) {
	for i := 0; i < len(data)*2; i++ {
		isContains := m.ContainsKey(i)
		if i < len(data) {
			if !isContains {
				t.Errorf("HashMap.Contains(%d) return false, expect true", i)
			}
		} else {
			if isContains {
				t.Errorf("HashMap.Contains(%d) return true, expect false", i)
			}
		}
	}
}

func testMapForEach(t *testing.T, m Map[int, int], data []int) {
	n := 0
	err := m.ForEach(func(k, v int) error {
		n++
		if data[k] != v {
			t.Errorf("")
		}

		return nil
	})
	if err != nil {
		t.Errorf("HashMap.ForEach returns %v, expect nil", err)
	}

	if n != m.Size() {
		t.Errorf("HashMap.ForEach run handler %d times, expect %d", n, m.Size())
	}
}

func testMapRemove(t *testing.T, m Map[int, int], data []int) {
	for i := 0; i < len(data)*2; i += 2 {
		old := m.Remove(i)
		if i < len(data) {
			if old != data[i] {
				t.Errorf("HashMap.Remove(%d) returns %d, expect %d", i, old, data[i])
			}
		} else if old != 0 {
			t.Errorf("HashMap.Remove(%d) returns %d, expect 0", i, old)
		}
	}
}

func testMapClear(t *testing.T, m Map[int, int]) {
	if m.IsEmpty() {
		t.Error("HashMap.IsEmpty() return true, expect false")
	}

	m.Clear()

	if !m.IsEmpty() {
		t.Error("HashMap.IsEmpty() return false, expect true")
	}
}

func testMap(t *testing.T, m Map[int, int]) {
	data := rand.Perm(10)

	if !m.IsEmpty() {
		t.Error("HashMap.IsEmpty() return false, expect true")
	}

	testMapPut(t, m, data)

	testMapGet(t, m, data)

	testMapContains(t, m, data)

	testMapForEach(t, m, data)

	testMapRemove(t, m, data)

	testMapClear(t, m)
}
