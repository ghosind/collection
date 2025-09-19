package internal

import (
	"testing"

	"github.com/ghosind/collection"
	"github.com/ghosind/go-assert"
)

func TestCheckIndex(t *testing.T) {
	a := assert.New(t)

	a.NotPanicNow(func() {
		CheckIndex(0, 1)
		CheckIndex(0, 10)
		CheckIndex(9, 10)
	})

	a.PanicOfNow(func() {
		CheckIndex(-1, 10)
	}, collection.ErrOutOfBounds)

	a.PanicOfNow(func() {
		CheckIndex(10, 10)
	}, collection.ErrOutOfBounds)
}
