package dict

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/ghosind/collection"
	"github.com/ghosind/go-assert"
)

type dictTestConstructor func() collection.Dict[string, string]

var testDataEn = map[string]string{
	"zero":  "0",
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

var testDataZh = map[string]string{
	"零": "0",
	"一": "1",
	"二": "2",
	"三": "3",
	"四": "4",
	"五": "5",
	"六": "6",
	"七": "7",
	"八": "8",
	"九": "9",
}

func testDict(a *assert.Assertion, constructor dictTestConstructor) {
	testDictClear(a, constructor)
	testDictClone(a, constructor)
	testDictContainsKey(a, constructor)
	testDictEquals(a, constructor)
	testDictForEach(a, constructor)
	testDictGet(a, constructor)
	testDictGetDefault(a, constructor)
	testDictIsEmpty(a, constructor)
	testDictIter(a, constructor)
	testDictKeys(a, constructor)
	testDictKeysIter(a, constructor)
	testDictPut(a, constructor)
	testDictRemove(a, constructor)
	testDictReplace(a, constructor)
	testDictSize(a, constructor)
	testDictString(a, constructor)
	testDictValues(a, constructor)
	testDictValuesIter(a, constructor)
}

func testDictClear(a *assert.Assertion, constructor dictTestConstructor) {
	d := constructor()
	for k, v := range testDataEn {
		d.Put(k, v)
	}
	a.EqualNow(d.Size(), len(testDataEn))
	d.Clear()
	a.EqualNow(d.Size(), 0)
	for k := range testDataEn {
		a.NotTrueNow(d.ContainsKey(k))
	}
}

func testDictClone(a *assert.Assertion, constructor dictTestConstructor) {
	d1 := constructor()
	for k, v := range testDataEn {
		d1.Put(k, v)
	}
	d2 := d1.Clone()
	a.TrueNow(d1.Equals(d2))
}

func testDictContainsKey(a *assert.Assertion, constructor dictTestConstructor) {
	d := constructor()
	for k, v := range testDataEn {
		d.Put(k, v)
	}
	for k := range testDataEn {
		a.TrueNow(d.ContainsKey(k))
	}
	for k := range testDataZh {
		a.NotTrueNow(d.ContainsKey(k))
	}
}

func testDictEquals(a *assert.Assertion, constructor dictTestConstructor) {
	d1 := constructor()
	a.NotTrueNow(d1.Equals(nil))

	for k, v := range testDataEn {
		d1.Put(k, v)
	}

	d2 := constructor()

	for k, v := range testDataEn {
		d2.Put(k, v)
	}
	a.TrueNow(d1.Equals(d2))

	for k, v := range testDataEn {
		d2.Put(k, v+v)
	}
	a.NotTrueNow(d1.Equals(d2))

	for k, v := range testDataZh {
		d2.Put(k, v)
	}
	a.NotTrueNow(d1.Equals(d2))

	d3 := constructor()
	for k, v := range testDataZh {
		d3.Put(k, v)
	}
	a.NotTrueNow(d1.Equals(d3))
}

func testDictForEach(a *assert.Assertion, constructor dictTestConstructor) {
	d := constructor()
	for k, v := range testDataEn {
		d.Put(k, v)
	}
	count := 0
	// ForEach should iterate over all elements
	err := d.ForEach(func(k string, v string) error {
		a.EqualNow(v, testDataEn[k])
		count++
		return nil
	})
	a.NilNow(err)
	a.EqualNow(count, len(testDataEn))

	// ForEach should exit early if an error is returned
	expectedErr := errors.New("expected error")
	count = 0
	err = d.ForEach(func(k string, v string) error {
		count++
		return expectedErr
	})
	a.EqualNow(err, expectedErr)
	a.EqualNow(count, 1)
}

func testDictGet(a *assert.Assertion, constructor dictTestConstructor) {
	d := constructor()
	for k, v := range testDataEn {
		d.Put(k, v)
	}
	for k, v := range testDataEn {
		vv, found := d.Get(k)
		a.TrueNow(found)
		a.EqualNow(v, vv)
	}

	for k := range testDataZh {
		vv, found := d.Get(k)
		a.NotTrueNow(found)
		a.EqualNow("", vv)
	}
}

func testDictGetDefault(a *assert.Assertion, constructor dictTestConstructor) {
	d := constructor()
	for k, v := range testDataEn {
		d.Put(k, v)
	}
	for k, v := range testDataEn {
		vv := d.GetDefault(k, "default")
		a.EqualNow(v, vv)
	}

	for k := range testDataZh {
		vv := d.GetDefault(k, "default")
		a.EqualNow("default", vv)
	}
}

func testDictIsEmpty(a *assert.Assertion, constructor dictTestConstructor) {
	d := constructor()
	a.TrueNow(d.IsEmpty())
	for k, v := range testDataEn {
		d.Put(k, v)
	}
	a.NotTrueNow(d.IsEmpty())
	d.Clear()
	a.TrueNow(d.IsEmpty())
}

func testDictKeys(a *assert.Assertion, constructor dictTestConstructor) {
	d := constructor()
	for k, v := range testDataEn {
		d.Put(k, v)
	}
	keys := d.Keys()
	a.EqualNow(len(keys), len(testDataEn))
	for _, k := range keys {
		a.TrueNow(d.ContainsKey(k))
	}
}

func testDictPut(a *assert.Assertion, constructor dictTestConstructor) {
	d := constructor()
	for k, v := range testDataEn {
		d.Put(k, v)
		a.TrueNow(d.ContainsKey(k))
	}
	a.EqualNow(d.Size(), len(testDataEn))
}

func testDictRemove(a *assert.Assertion, constructor dictTestConstructor) {
	d := constructor()
	for k, v := range testDataEn {
		d.Put(k, v)
	}
	for k := range testDataEn {
		d.Remove(k)
		a.NotTrueNow(d.ContainsKey(k))
	}
	a.EqualNow(d.Size(), 0)
}

func testDictReplace(a *assert.Assertion, constructor dictTestConstructor) {
	d := constructor()
	for k, v := range testDataEn {
		d.Put(k, v)
	}
	for k, v := range testDataEn {
		old, found := d.Replace(k, v+"-new")
		a.TrueNow(found)
		a.EqualNow(v, old)
		a.EqualNow(v+"-new", d.GetDefault(k, ""))
	}
	for k, v := range testDataZh {
		old, found := d.Replace(k, v)
		a.NotTrueNow(found)
		a.EqualNow("", old)
		a.NotTrueNow(d.ContainsKey(k))
	}
}

func testDictSize(a *assert.Assertion, constructor dictTestConstructor) {
	d := constructor()
	a.EqualNow(d.Size(), 0)
	for k, v := range testDataEn {
		d.Put(k, v)
	}
	a.EqualNow(d.Size(), len(testDataEn))
	d.Clear()
	a.EqualNow(d.Size(), 0)
}

func testDictString(a *assert.Assertion, constructor dictTestConstructor) {
	d := constructor()
	for k, v := range testDataEn {
		d.Put(k, v)
	}

	str := d.String()
	a.TrueNow(strings.HasPrefix(str, "dict["))
	a.TrueNow(strings.HasSuffix(str, "]"))

	for k, v := range testDataEn {
		a.TrueNow(strings.Contains(str, fmt.Sprintf("%v: %v", k, v)))
	}
}

func testDictValues(a *assert.Assertion, constructor dictTestConstructor) {
	d := constructor()
	expectedValues := make([]string, 0, len(testDataEn))
	for k, v := range testDataEn {
		d.Put(k, v)
		expectedValues = append(expectedValues, v)
	}

	values := d.Values()
	a.EqualNow(len(values), len(testDataEn))

	sort.Strings(expectedValues)
	sort.Strings(values)

	a.EqualNow(expectedValues, values)
}
