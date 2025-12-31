package dict

import (
	"testing"

	"github.com/ghosind/collection"
	"github.com/ghosind/go-assert"
)

func syncDictConstructor(initData ...map[string]string) collection.Dict[string, string] {
	if len(initData) == 0 || len(initData[0]) == 0 {
		return NewSyncDict[string, string]()
	}
	return NewSyncDictFrom(initData[0])
}

func TestSyncDict(t *testing.T) {
	a := assert.New(t)

	testDict(a, syncDictConstructor)
}

func BenchmarkSyncDictGet(b *testing.B) {
	benchmarkDictGet(b, syncDictConstructor, true)
}

func BenchmarkSyncDictPut(b *testing.B) {
	benchmarkDictPut(b, syncDictConstructor, true)
}
