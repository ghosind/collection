package internal

import (
	"testing"
	"time"

	"github.com/ghosind/go-assert"
)

type strType struct{}

func (s strType) MarshalText() (text []byte, err error) {
	return []byte("custom string"), nil
}

func TestValueString(t *testing.T) {
	a := assert.New(t)

	a.EqualNow(ValueString(nil), "<nil>")
	a.EqualNow(ValueString("hello"), "hello")
	a.EqualNow(ValueString(123), "123")
	a.EqualNow(ValueString(int8(123)), "123")
	a.EqualNow(ValueString(int16(123)), "123")
	a.EqualNow(ValueString(int32(123)), "123")
	a.EqualNow(ValueString(int64(123)), "123")
	a.EqualNow(ValueString(uint(123)), "123")
	a.EqualNow(ValueString(uint8(123)), "123")
	a.EqualNow(ValueString(uint16(123)), "123")
	a.EqualNow(ValueString(uint32(123)), "123")
	a.EqualNow(ValueString(uint64(123)), "123")
	a.EqualNow(ValueString(uintptr(123)), "123")
	a.EqualNow(ValueString(float32(1.23)), "1.23")
	a.EqualNow(ValueString(1.23), "1.23")
	a.EqualNow(ValueString(complex64(1+2i)), "(1+2i)")
	a.EqualNow(ValueString(complex128(1+2i)), "(1+2i)")
	a.EqualNow(ValueString(true), "true")
	a.EqualNow(ValueString([]int{1, 2, 3}), "[1 2 3]")
	a.EqualNow(ValueString(map[string]int{"a": 1, "b": 2}), "map[a:1 b:2]")
	a.EqualNow(ValueString(strType{}), "custom string")
	a.EqualNow(ValueString(time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)), "2000-01-01 00:00:00 +0000 UTC")
}
