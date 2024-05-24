package dict

import (
	"reflect"
	"sync"
	"sync/atomic"

	"github.com/ghosind/collection"
)

// SyncDict is a thread-safe map implementation based on sync.Map's algorithm.
type SyncDict[K comparable, V any] struct {
	mu       sync.Mutex
	read     atomic.Pointer[syncReadOnly[K, V]]
	dirty    map[K]*syncEntry[V]
	misses   int
	zero     V
	expunged *V
}

type syncReadOnly[K comparable, V any] struct {
	m       map[K]*syncEntry[V]
	amended bool
}

type syncEntry[T any] struct {
	p        atomic.Pointer[T]
	expunged *T
}

func newSyncEntry[T any](v T, expunged *T) *syncEntry[T] {
	e := new(syncEntry[T])
	e.p.Store(&v)
	e.expunged = expunged
	return e
}

func (e *syncEntry[T]) load(val T) (value T, ok bool) {
	p := e.p.Load()
	if p == nil || p == e.expunged {
		return val, false
	}
	return *p, true
}

func (e *syncEntry[T]) trySwap(val *T) (*T, bool) {
	for {
		p := e.p.Load()
		if p == e.expunged {
			return nil, false
		}
		if e.p.CompareAndSwap(p, val) {
			return p, true
		}
	}
}

func (e *syncEntry[T]) delete() (*T, bool) {
	for {
		p := e.p.Load()
		if p == nil || p == e.expunged {
			return nil, false
		}
		if e.p.CompareAndSwap(p, nil) {
			return p, true
		}
	}
}

func (e *syncEntry[T]) unexpungeLocked() bool {
	return e.p.CompareAndSwap(e.expunged, nil)
}

func (e *syncEntry[T]) swapLocked(v *T) *T {
	return e.p.Swap(v)
}

func (e *syncEntry[T]) tryExpungeLocked() bool {
	p := e.p.Load()
	for p == nil {
		if e.p.CompareAndSwap(nil, e.expunged) {
			return true
		}
		p = e.p.Load()
	}
	return p == e.expunged
}

// NewSyncDict creates a new SyncDict.
func NewSyncDict[K comparable, V any]() *SyncDict[K, V] {
	d := new(SyncDict[K, V])

	d.expunged = new(V)

	return d
}

func (d *SyncDict[K, V]) loadReadOnly() syncReadOnly[K, V] {
	if p := d.read.Load(); p != nil {
		return *p
	}
	return syncReadOnly[K, V]{}
}

func (d *SyncDict[K, V]) loadPresentReadOnly() syncReadOnly[K, V] {
	read := d.loadReadOnly()
	if read.amended {
		d.mu.Lock()
		read = d.loadReadOnly()
		if read.amended {
			read = syncReadOnly[K, V]{m: d.dirty}
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
	d.dirty = make(map[K]*syncEntry[V], len(read.m))
	for k, e := range read.m {
		if !e.tryExpungeLocked() {
			d.dirty[k] = e
		}
	}
}

func (d *SyncDict[K, V]) missLocked() {
	d.misses++
	if d.misses < len(d.dirty) {
		return
	}

	d.read.Store(&syncReadOnly[K, V]{m: d.dirty})
	d.dirty = nil
	d.misses = 0
}

// Get returns the value which associated to the specified key.
func (d *SyncDict[K, V]) get(key K, val V) (V, bool) {
	read := d.loadReadOnly()
	e, ok := read.m[key]
	if !ok && read.amended {
		d.mu.Lock()
		read = d.loadReadOnly()
		e, ok = read.m[key]
		if !ok && read.amended {
			e, ok = d.dirty[key]
			d.missLocked()
		}
		d.mu.Unlock()
	}
	if !ok {
		return val, false
	}
	return e.load(val)
}

func (d *SyncDict[K, V]) swap(key K, val V, ignore bool) (*V, bool) {
	read := d.loadReadOnly()
	if e, ok := read.m[key]; ok {
		return e.trySwap(&val)
	}

	d.mu.Lock()
	defer d.mu.Unlock()

	read = d.loadReadOnly()
	if e, ok := read.m[key]; ok {
		if e.unexpungeLocked() {
			d.dirty[key] = e
		}
		if v := e.swapLocked(&val); v != nil {
			return v, true
		}
	} else if e, ok := d.dirty[key]; ok {
		if v := e.swapLocked(&val); v != nil {
			return v, true
		}
	} else if !ignore {
		if !read.amended {
			d.dirtyLocked()
			d.read.Store(&syncReadOnly[K, V]{m: read.m, amended: true})
		}
		d.dirty[key] = newSyncEntry(val, d.expunged)
	}
	return nil, false
}

// // Clear removes all key-value pairs in this dictionary.
func (d *SyncDict[K, V]) Clear() {
	d.mu.Lock()
	defer d.mu.Unlock()
	read := d.loadReadOnly()
	if read.amended {
		d.dirty = nil
		d.misses = 0
	}
	read = syncReadOnly[K, V]{m: make(map[K]*syncEntry[V])}
	copyRead := read
	d.read.Store(&copyRead)
}

// Clone returns a copy of this dictionary.
func (d *SyncDict[K, V]) Clone() collection.Dict[K, V] {
	read := d.loadPresentReadOnly()
	m := make(map[K]*syncEntry[V])
	expunged := new(V)

	for k, e := range read.m {
		v, ok := e.load(d.zero)
		if ok {
			m[k] = newSyncEntry(v, expunged)
		}
	}

	newDict := new(SyncDict[K, V])
	newDict.expunged = expunged
	newDict.read.Store(&syncReadOnly[K, V]{m: m})

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

	for k, e := range read.m {
		dv, ok := e.load(d.zero)
		if !ok {
			continue
		}
		rs++

		oe, ok := oRead.m[k]
		if !ok {
			return false
		}
		ov, ok := oe.load(od.zero)
		if !ok {
			return false
		}

		if !reflect.DeepEqual(dv, ov) {
			return false
		}
	}

	for _, e := range oRead.m {
		_, ok := e.load(d.zero)
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

	for k, e := range read.m {
		v, ok := e.load(d.zero)
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
	if len(read.m) == 0 {
		return true
	}

	for _, e := range read.m {
		_, ok := e.load(d.zero)
		if ok {
			return false
		}
	}

	return true
}

// Keys returns a slice that contains all the keys in this dictionary.
func (d *SyncDict[K, V]) Keys() []K {
	read := d.loadPresentReadOnly()

	keys := make([]K, 0, len(read.m))
	for k, e := range read.m {
		_, ok := e.load(d.zero)
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
	e, ok := read.m[key]
	if !ok && read.amended {
		d.mu.Lock()
		read = d.loadReadOnly()
		e, ok = read.m[key]
		if !ok && read.amended {
			e, ok = d.dirty[key]
			delete(d.dirty, key)
			d.missLocked()
		}
		d.mu.Unlock()
	}
	if ok {
		vp, ok := e.delete()
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

	for _, e := range read.m {
		_, ok := e.load(d.zero)
		if !ok {
			continue
		}
		size++
	}

	return size
}

// Values returns a slice that contains all the values in this dictionary.
func (d *SyncDict[K, V]) Values() []V {
	read := d.loadPresentReadOnly()

	keys := make([]V, 0, len(read.m))
	for _, e := range read.m {
		v, ok := e.load(d.zero)
		if !ok {
			continue
		}
		keys = append(keys, v)
	}

	return keys
}
