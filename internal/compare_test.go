package internal

import (
	"testing"

	"github.com/ghosind/go-assert"
)

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
