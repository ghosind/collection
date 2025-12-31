package dict

import (
	"testing"

	"github.com/ghosind/collection"
	"github.com/ghosind/go-assert"
)

func lockDictConstructor(initData ...map[string]string) collection.Dict[string, string] {
	if len(initData) == 0 || len(initData[0]) == 0 {
		return NewLockDict[string, string]()
	}
	return NewLockDictFrom[string, string](initData[0])
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
