package internal

import (
	"testing"

	"github.com/ghosind/go-assert"
)

func TestInSlice(t *testing.T) {
	a := assert.New(t)

	testInSlice(a, []int{1, 2, 3, 4, 5}, 3, true)
	testInSlice(a, []int{1, 2, 3, 4, 5}, 6, false)

	testInSlice(a, []string{"a", "b", "c"}, "b", true)
	testInSlice(a, []string{"a", "b", "c"}, "d", false)

	testInSlice(a, []int8{1, 2, 3}, int8(2), true)
	testInSlice(a, []int8{1, 2, 3}, int8(4), false)

	testInSlice(a, []int16{10, 20, 30}, int16(20), true)
	testInSlice(a, []int16{10, 20, 30}, int16(40), false)

	testInSlice(a, []int32{100, 200, 300}, int32(200), true)
	testInSlice(a, []int32{100, 200, 300}, int32(400), false)

	testInSlice(a, []int64{1000, 2000, 3000}, int64(2000), true)
	testInSlice(a, []int64{1000, 2000, 3000}, int64(4000), false)

	testInSlice(a, []uint{1, 2, 3}, uint(2), true)
	testInSlice(a, []uint{1, 2, 3}, uint(4), false)

	testInSlice(a, []uint8{10, 20, 30}, uint8(20), true)
	testInSlice(a, []uint8{10, 20, 30}, uint8(40), false)

	testInSlice(a, []uint16{100, 200, 300}, uint16(200), true)
	testInSlice(a, []uint16{100, 200, 300}, uint16(400), false)

	testInSlice(a, []uint32{1000, 2000, 3000}, uint32(2000), true)
	testInSlice(a, []uint32{1000, 2000, 3000}, uint32(4000), false)

	testInSlice(a, []uint64{10000, 20000, 30000}, uint64(20000), true)
	testInSlice(a, []uint64{10000, 20000, 30000}, uint64(40000), false)

	testInSlice(a, []uintptr{1, 2, 3}, uintptr(2), true)
	testInSlice(a, []uintptr{1, 2, 3}, uintptr(4), false)

	testInSlice(a, []float32{1.1, 2.2, 3.3}, float32(2.2), true)
	testInSlice(a, []float32{1.1, 2.2, 3.3}, float32(4.4), false)

	testInSlice(a, []float64{1.11, 2.22, 3.33}, float64(2.22), true)
	testInSlice(a, []float64{1.11, 2.22, 3.33}, float64(4.44), false)

	testInSlice(a, []complex64{1 + 2i, 3 + 4i}, complex64(3+4i), true)
	testInSlice(a, []complex64{1 + 2i, 3 + 4i}, complex64(5+6i), false)

	testInSlice(a, []complex128{1 + 2i, 3 + 4i}, complex128(3+4i), true)
	testInSlice(a, []complex128{1 + 2i, 3 + 4i}, complex128(5+6i), false)

	testInSlice(a, []bool{true, false}, false, true)
	testInSlice(a, []bool{true, false}, true, true)
	testInSlice(a, []bool{true}, false, false)

	type myStruct struct {
		A int
		B string
	}

	s1 := myStruct{A: 1, B: "a"}
	s2 := myStruct{A: 2, B: "b"}
	s3 := myStruct{A: 3, B: "c"}

	testInSlice(a, []myStruct{s1, s2}, s2, true)
	testInSlice(a, []myStruct{s1, s2}, s3, false)
}

func testInSlice[T any](a *assert.Assertion, slice []T, element T, expected bool) {
	cache := MakeSliceCacheMap(slice)
	defer ReleaseCacheMap(cache)
	result := InSlice(element, slice, cache)
	a.EqualNow(expected, result)
}

func TestEqual(t *testing.T) {
	a := assert.New(t)

	a.TrueNow(Equal(1, 1))
	a.TrueNow(Equal(int8(1), int8(1)))
	a.TrueNow(Equal(int16(1), int16(1)))
	a.TrueNow(Equal(int32(1), int32(1)))
	a.TrueNow(Equal(int64(1), int64(1)))
	a.TrueNow(Equal(uint(1), uint(1)))
	a.TrueNow(Equal(uint8(1), uint8(1)))
	a.TrueNow(Equal(uint16(1), uint16(1)))
	a.TrueNow(Equal(uint32(1), uint32(1)))
	a.TrueNow(Equal(uint64(1), uint64(1)))
	a.TrueNow(Equal(uintptr(1), uintptr(1)))
	a.TrueNow(Equal("hello", "hello"))
	a.TrueNow(Equal(1.23, 1.23))
	a.TrueNow(Equal(float32(1.23), float32(1.23)))
	a.TrueNow(Equal(complex64(1+2i), complex64(1+2i)))
	a.TrueNow(Equal(complex128(1+2i), complex128(1+2i)))
	a.TrueNow(Equal(true, true))
	a.TrueNow(Equal([]int{1, 2, 3}, []int{1, 2, 3}))
	a.TrueNow(Equal(map[string]int{"a": 1, "b": 2}, map[string]int{"a": 1, "b": 2}))

	a.NotTrueNow(Equal(1, 2))
	a.NotTrueNow(Equal(int8(1), int8(2)))
	a.NotTrueNow(Equal(int16(1), int16(2)))
	a.NotTrueNow(Equal(int32(1), int32(2)))
	a.NotTrueNow(Equal(int64(1), int64(2)))
	a.NotTrueNow(Equal(uint(1), uint(2)))
	a.NotTrueNow(Equal(uint8(1), uint8(2)))
	a.NotTrueNow(Equal(uint16(1), uint16(2)))
	a.NotTrueNow(Equal(uint32(1), uint32(2)))
	a.NotTrueNow(Equal(uint64(1), uint64(2)))
	a.NotTrueNow(Equal(uintptr(1), uintptr(2)))
	a.NotTrueNow(Equal("hello", "world"))
	a.NotTrueNow(Equal(1.23, 4.56))
	a.NotTrueNow(Equal(float32(1.23), float32(4.56)))
	a.NotTrueNow(Equal(complex64(1+2i), complex64(3+4i)))
	a.NotTrueNow(Equal(complex128(1+2i), complex128(3+4i)))
	a.NotTrueNow(Equal(true, false))
	a.NotTrueNow(Equal([]int{1, 2, 3}, []int{4, 5, 6}))
	a.NotTrueNow(Equal(map[string]int{"a": 1, "b": 2}, map[string]int{"a": 3, "b": 4}))
}
