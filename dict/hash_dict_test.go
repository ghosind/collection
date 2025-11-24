package dict

import (
	"testing"

	"github.com/ghosind/collection"
	"github.com/ghosind/go-assert"
)

func hashDictConstructor() collection.Dict[string, string] {
	return NewHashDict[string, string]()
}

func TestHashDict(t *testing.T) {
	a := assert.New(t)

	testDict(a, hashDictConstructor)
}

func BenchmarkHashDictGet(b *testing.B) {
	benchmarkDictGet(b, hashDictConstructor, false)
}

func BenchmarkHashDictPut(b *testing.B) {
	benchmarkDictPut(b, hashDictConstructor, false)
}
