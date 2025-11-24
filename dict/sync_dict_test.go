package dict

import (
	"testing"

	"github.com/ghosind/collection"
	"github.com/ghosind/go-assert"
)

func syncDictConstructor() collection.Dict[string, string] {
	return NewSyncDict[string, string]()
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
