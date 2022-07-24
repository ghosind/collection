package collection

import "testing"

type testStruct struct {
	v int
}

var intData = []int{1, 2, 3, 4, 5, 6, 7}
var strData = []string{"a", "b", "c", "d", "e", "f", "g"}
var structData = []testStruct{{1}, {2}, {3}, {4}, {5}, {6}, {7}}
var pointerData = []*testStruct{{1}, {2}, {3}, {4}, {5}, {6}, {7}}

func testSetAdd[T comparable](t *testing.T, set Set[T], e T) {
	if ok := set.Add(e); !ok {
		t.Errorf("set.Add(%v) returned false", e)
	}

	if ok := set.Add(e); ok {
		t.Errorf("set.Add(%v) returned true", e)
	}
}

func testSetAddAll[T comparable](t *testing.T, set Set[T], c ...T) {
	if ok := set.AddAll(c...); !ok {
		t.Errorf("set.AddAll(%v) returned false", c)
	}

	if ok := set.AddAll(c...); ok {
		t.Errorf("set.AddAll(%v) returned true", c)
	}
}

func testSetContains[T comparable](t *testing.T, set Set[T], data []T) {
	if found := set.Contains(data[0]); !found {
		t.Errorf("set.Contains(%v) returned false", data[0])
	}

	if found := set.Contains(data[len(data)-1]); found {
		t.Errorf("set.Contains(%v) returned true", data[len(data)-1])
	}

	if found := set.ContainsAll(data[0 : len(data)-1]...); !found {
		t.Errorf("set.ContainsAll(%v) returned false", data[0:len(data)-1])
	}

	if found := set.ContainsAll(data...); found {
		t.Errorf("set.ContainsAll(%v) returned true", data)
	}
}

func testSetToSlice[T comparable](t *testing.T, set Set[T]) {
	slice := set.ToSlice()
	for _, e := range slice {
		if !set.Contains(e) {
			t.Errorf("set.ToSlice() returned %v, but set.Contains(%v) returned false", slice, e)
		}
	}
}

func testSetRemove[T comparable](t *testing.T, set Set[T], data []T) {
	if isEmpty := set.IsEmpty(); isEmpty {
		t.Errorf("set.IsEmpty() returned true")
	}
	if size := set.Size(); size != len(data)-1 {
		t.Errorf("set.Size() returned %v, expected %v", size, len(data)-1)
	}

	if ok := set.Remove(data[0]); !ok {
		t.Errorf("set.Remove(%v) returned false", data[0])
	}
	if size := set.Size(); size != len(data)-2 {
		t.Errorf("set.Size() returned %v, expected %v", size, len(data)-2)
	}

	if ok := set.Remove(data[len(data)-1]); ok {
		t.Errorf("set.Remove(%v) returned true", data[len(data)-1])
	}
	if size := set.Size(); size != len(data)-2 {
		t.Errorf("set.Size() returned %v, expected %v", size, len(data)-2)
	}
}

func testSetRemoveAll[T comparable](t *testing.T, set Set[T], data []T) {
	if ok := set.RemoveAll(data[0:2]...); !ok {
		t.Errorf("set.RemoveAll(%v) returned false", data[0:2])
	}
	if size := set.Size(); size != len(data)-3 {
		t.Errorf("set.Size() returned %v, expected %d", size, len(data)-3)
	}
}

func testSetClear[T comparable](t *testing.T, set Set[T]) {
	set.Clear()
	if isEmpty := set.IsEmpty(); !isEmpty {
		t.Errorf("set.IsEmpty() returned false")
	}
}

func testSet[T comparable](t *testing.T, set Set[T], data []T) {
	if set == nil {
		t.Errorf("set is nil")
	}

	testSetAdd(t, set, data[0])
	testSetAddAll(t, set, data[0:len(data)-1]...)
	testSetContains(t, set, data)
	testSetToSlice(t, set)
	testSetRemove(t, set, data)
	testSetRemoveAll(t, set, data)
	testSetClear(t, set)
}
