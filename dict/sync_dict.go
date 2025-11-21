package dict

import (
	"bytes"
	"encoding/json"
	"sync"
	"sync/atomic"

	"github.com/ghosind/collection"
	"github.com/ghosind/collection/internal"
)

// SyncDict is a thread-safe map implementation based on sync.Map's algorithm.
type SyncDict[K comparable, V any] struct {
	mu       sync.Mutex
	read     atomic.Pointer[internal.SyncReadOnly[K, V]]
	dirty    map[K]*internal.SyncEntry[V]
	misses   int
	zero     V
	expunged *V
}

// NewSyncDict creates a new SyncDict.
func NewSyncDict[K comparable, V any]() *SyncDict[K, V] {
	d := new(SyncDict[K, V])

	d.expunged = new(V)

	return d
}

func (d *SyncDict[K, V]) loadReadOnly() internal.SyncReadOnly[K, V] {
	if p := d.read.Load(); p != nil {
		return *p
	}
	return internal.SyncReadOnly[K, V]{}
}

func (d *SyncDict[K, V]) loadPresentReadOnly() internal.SyncReadOnly[K, V] {
	read := d.loadReadOnly()
	if read.Amended {
		d.mu.Lock()
		read = d.loadReadOnly()
		if read.Amended {
			read = internal.SyncReadOnly[K, V]{M: d.dirty}
			copyRead := read
			d.read.Store(&copyRead)
			d.dirty = nil
			d.misses = 0
		}
		d.mu.Unlock()
	}

	return read
}

func (d *SyncDict[K, V]) dirtyLocked() {
	if d.dirty != nil {
		return
	}

	read := d.loadReadOnly()
	d.dirty = make(map[K]*internal.SyncEntry[V], len(read.M))
	for k, e := range read.M {
		if !e.TryExpungeLocked() {
			d.dirty[k] = e
		}
	}
}

func (d *SyncDict[K, V]) missLocked() {
	d.misses++
	if d.misses < len(d.dirty) {
		return
	}

	d.read.Store(&internal.SyncReadOnly[K, V]{M: d.dirty})
	d.dirty = nil
	d.misses = 0
}

// Get returns the value which associated to the specified key.
func (d *SyncDict[K, V]) get(key K, val V) (V, bool) {
	read := d.loadReadOnly()
	e, ok := read.M[key]
	if !ok && read.Amended {
		d.mu.Lock()
		read = d.loadReadOnly()
		e, ok = read.M[key]
		if !ok && read.Amended {
			e, ok = d.dirty[key]
			d.missLocked()
		}
		d.mu.Unlock()
	}
	if !ok {
		return val, false
	}
	return e.Load(val)
}

func (d *SyncDict[K, V]) swap(key K, val V, ignore bool) (*V, bool) {
	read := d.loadReadOnly()
	if e, ok := read.M[key]; ok {
		return e.TrySwap(&val)
	}

	d.mu.Lock()
	defer d.mu.Unlock()

	read = d.loadReadOnly()
	if e, ok := read.M[key]; ok {
		if e.UnexpungeLocked() {
			d.dirty[key] = e
		}
		if v := e.SwapLocked(&val); v != nil {
			return v, true
		}
	} else if e, ok := d.dirty[key]; ok {
		if v := e.SwapLocked(&val); v != nil {
			return v, true
		}
	} else if !ignore {
		if !read.Amended {
			d.dirtyLocked()
			d.read.Store(&internal.SyncReadOnly[K, V]{M: read.M, Amended: true})
		}
		d.dirty[key] = internal.NewSyncEntry(val, d.expunged)
	}
	return nil, false
}

// // Clear removes all key-value pairs in this dictionary.
func (d *SyncDict[K, V]) Clear() {
	d.mu.Lock()
	defer d.mu.Unlock()
	read := d.loadReadOnly()
	if read.Amended {
		d.dirty = nil
		d.misses = 0
	}
	read = internal.SyncReadOnly[K, V]{M: make(map[K]*internal.SyncEntry[V])}
	copyRead := read
	d.read.Store(&copyRead)
}

// Clone returns a copy of this dictionary.
func (d *SyncDict[K, V]) Clone() collection.Dict[K, V] {
	read := d.loadPresentReadOnly()
	m := make(map[K]*internal.SyncEntry[V])
	expunged := new(V)

	for k, e := range read.M {
		v, ok := e.Load(d.zero)
		if ok {
			m[k] = internal.NewSyncEntry(v, expunged)
		}
	}

	newDict := new(SyncDict[K, V])
	newDict.expunged = expunged
	newDict.read.Store(&internal.SyncReadOnly[K, V]{M: m})

	return newDict
}

// ContainsKey returns true if this dictionary contains a key-value pair with the specified key.
func (d *SyncDict[K, V]) ContainsKey(key K) bool {
	_, ok := d.get(key, d.zero)
	return ok
}

// Equals compares this dictionary with the object pass from parameter.
func (d *SyncDict[K, V]) Equals(o any) bool {
	od, ok := o.(*SyncDict[K, V])
	if !ok {
		return false
	}

	read := d.loadPresentReadOnly()
	oRead := od.loadPresentReadOnly()

	rs := 0
	os := 0

	for k, e := range read.M {
		dv, ok := e.Load(d.zero)
		if !ok {
			continue
		}
		rs++

		oe, ok := oRead.M[k]
		if !ok {
			return false
		}
		ov, ok := oe.Load(od.zero)
		if !ok {
			return false
		}

		if !internal.Equal(dv, ov) {
			return false
		}
	}

	for _, e := range oRead.M {
		_, ok := e.Load(d.zero)
		if !ok {
			continue
		}
		os++
		if rs < os {
			return false
		}
	}

	return rs == os
}

// ForEach performs the given handler for each key-value pairs in the dictionary until all pairs
// have been processed or the handler returns an error.
func (d *SyncDict[K, V]) ForEach(handler func(K, V) error) error {
	read := d.loadPresentReadOnly()

	for k, e := range read.M {
		v, ok := e.Load(d.zero)
		if !ok {
			continue
		}
		if err := handler(k, v); err != nil {
			return err
		}
	}

	return nil
}

// Get returns the value which associated to the specified key.
func (d *SyncDict[K, V]) Get(key K) (V, bool) {
	return d.get(key, d.zero)
}

// GetDefault returns the value associated with the specified key, and returns the default value if
// this dictionary contains no pair with the key.
func (d *SyncDict[K, V]) GetDefault(key K, defaultVal V) V {
	v, _ := d.get(key, defaultVal)
	return v
}

// IsEmpty returns true if this dictionary is empty.
func (d *SyncDict[K, V]) IsEmpty() bool {
	read := d.loadPresentReadOnly()
	if len(read.M) == 0 {
		return true
	}

	for _, e := range read.M {
		_, ok := e.Load(d.zero)
		if ok {
			return false
		}
	}

	return true
}

// Keys returns a slice that contains all the keys in this dictionary.
func (d *SyncDict[K, V]) Keys() []K {
	read := d.loadPresentReadOnly()

	keys := make([]K, 0, len(read.M))
	for k, e := range read.M {
		_, ok := e.Load(d.zero)
		if !ok {
			continue
		}
		keys = append(keys, k)
	}

	return keys
}

// Put associate the specified value with the specified key in this dictionary.
func (d *SyncDict[K, V]) Put(key K, val V) V {
	prev, ok := d.swap(key, val, false)
	if ok {
		return *prev
	} else {
		return d.zero
	}
}

// Remove removes the key-value pair with the specified key.
func (d *SyncDict[K, V]) Remove(key K) V {
	read := d.loadReadOnly()
	e, ok := read.M[key]
	if !ok && read.Amended {
		d.mu.Lock()
		read = d.loadReadOnly()
		e, ok = read.M[key]
		if !ok && read.Amended {
			e, ok = d.dirty[key]
			delete(d.dirty, key)
			d.missLocked()
		}
		d.mu.Unlock()
	}
	if ok {
		vp, ok := e.Delete()
		if ok {
			return *vp
		}
		return d.zero
	}
	return d.zero
}

// Replace replaces the value for the specified key only if it is currently in this dictionary.
func (d *SyncDict[K, V]) Replace(key K, val V) (V, bool) {
	prev, ok := d.swap(key, val, true)
	if ok {
		return *prev, ok
	} else {
		return d.zero, ok
	}
}

// Size returns the number of key-value pairs in this dictionary.
func (d *SyncDict[K, V]) Size() int {
	read := d.loadPresentReadOnly()
	size := 0

	for _, e := range read.M {
		_, ok := e.Load(d.zero)
		if !ok {
			continue
		}
		size++
	}

	return size
}

// String returns the string representation of this dictionary.
func (d *SyncDict[K, V]) String() string {
	buf := bytes.NewBufferString("dict[")
	read := d.loadPresentReadOnly()
	count := 0
	for k, e := range read.M {
		v, ok := e.Load(d.zero)
		if !ok {
			continue
		}
		if count > 0 {
			buf.WriteString(" ")
		}
		buf.WriteString(internal.ValueString(k))
		buf.WriteString(": ")
		buf.WriteString(internal.ValueString(v))
		count++
	}
	buf.WriteString("]")
	return buf.String()
}

// Values returns a slice that contains all the values in this dictionary.
func (d *SyncDict[K, V]) Values() []V {
	read := d.loadPresentReadOnly()

	keys := make([]V, 0, len(read.M))
	for _, e := range read.M {
		v, ok := e.Load(d.zero)
		if !ok {
			continue
		}
		keys = append(keys, v)
	}

	return keys
}

// MarshalJSON marshals the SyncDict as a JSON object (map).
func (d *SyncDict[K, V]) MarshalJSON() ([]byte, error) {
	m := make(map[K]V)
	_ = d.ForEach(func(k K, v V) error {
		m[k] = v
		return nil
	})
	return json.Marshal(m)
}

// UnmarshalJSON unmarshals a JSON object into the SyncDict.
func (d *SyncDict[K, V]) UnmarshalJSON(b []byte) error {
	var tmp map[K]V
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	d.Clear()
	for k, v := range tmp {
		d.Put(k, v)
	}
	return nil
}
