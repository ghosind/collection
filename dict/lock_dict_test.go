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
	return NewLockDictFrom(initData[0])
}

func TestLockDict(t *testing.T) {
	a := assert.New(t)

	testDict(a, lockDictConstructor)
}

func BenchmarkLockDict_Get(b *testing.B) {
	benchmarkDict_Get(b, lockDictConstructor, true)
}

func BenchmarkLockDict_Put(b *testing.B) {
	benchmarkDict_Put(b, lockDictConstructor, true)
}
