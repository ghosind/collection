package dict

import (
	"testing"

	"github.com/ghosind/collection"
	"github.com/ghosind/go-assert"
)

func lockDictConstructor() collection.Dict[string, string] {
	return NewLockDict[string, string]()
}

func TestLockDict(t *testing.T) {
	a := assert.New(t)

	testDict(a, lockDictConstructor)
}

func BenchmarkLockDictGet(b *testing.B) {
	benchmarkDictGet(b, lockDictConstructor, true)
}

func BenchmarkLockDictPut(b *testing.B) {
	benchmarkDictPut(b, lockDictConstructor, true)
}
