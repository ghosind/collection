package collection

import "testing"

func TestHashMap(t *testing.T) {
	testMap(t, NewHashMap[int, int]())
}
