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

func BenchmarkHashDictGet(b *testing.B) {
	benchmarkDictGet(b, hashDictConstructor, false)
}

func BenchmarkHashDictPut(b *testing.B) {
	benchmarkDictPut(b, hashDictConstructor, false)
}
