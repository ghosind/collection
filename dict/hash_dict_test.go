package dict

import (
	"testing"

	"github.com/ghosind/collection"
	"github.com/ghosind/go-assert"
)

func hashDictConstructor(initData ...map[string]string) collection.Dict[string, string] {
	if len(initData) == 0 || len(initData[0]) == 0 {
		return NewHashDict[string, string]()
	}
	return NewHashDictFrom(initData[0])
}

func TestHashDict(t *testing.T) {
	a := assert.New(t)

	testDict(a, hashDictConstructor)
}

func BenchmarkHashDict_Get(b *testing.B) {
	benchmarkDict_Get(b, hashDictConstructor, false)
}

func BenchmarkHashDict_Put(b *testing.B) {
	benchmarkDict_Put(b, hashDictConstructor, false)
}
