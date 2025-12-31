package set

import (
	"bytes"
	"encoding/json"
	"sync"
	"sync/atomic"

	"github.com/ghosind/collection"
	"github.com/ghosind/collection/internal"
)

// SyncSet is a thread-safe set implementation that based on sync.Map.
type SyncSet[T comparable] struct {
	mu     sync.Mutex
	read   atomic.Pointer[internal.SyncReadOnly[T, empty]]
	dirty  map[T]*internal.SyncEntry[empty]
	misses int
}

// NewSyncSet creates a new SyncSet.
func NewSyncSet[T comparable]() *SyncSet[T] {
	s := new(SyncSet[T])

	return s
}

// NewSyncSetFrom creates and returns a new SyncSet containing the elements of the
// provided collection.
func NewSyncSetFrom[T comparable](c ...T) *SyncSet[T] {
	s := new(SyncSet[T])
	m := make(map[T]*internal.SyncEntry[empty], len(c))
	for _, e := range c {
		m[e] = internal.NewSyncEntry(emptyZero, nilEmpty)
	}
	s.read.Store(&internal.SyncReadOnly[T, empty]{M: m})

	return s
}

// Add adds the specified element to this collection.
func (s *SyncSet[T]) Add(val T) bool {
	read := s.loadReadOnly()
	if e, ok := read.M[val]; ok {
		_, ok := e.TrySwap(&emptyZero)
		return !ok
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	read = s.loadReadOnly()
	if e, ok := read.M[val]; ok {
		if e.UnexpungeLocked() {
			s.dirty[val] = e
		}
		if v := e.SwapLocked(&emptyZero); v != nil {
			return false
		}
	} else if e, ok := s.dirty[val]; ok {
		if v := e.SwapLocked(&emptyZero); v != nil {
			return false
		}
	} else {
		if !read.Amended {
			s.dirtyLocked()
			s.read.Store(&internal.SyncReadOnly[T, empty]{M: read.M, Amended: true})
		}
		s.dirty[val] = internal.NewSyncEntry(emptyZero, nilEmpty)
	}
	return true
}

// AddAll adds all of the elements in the this collection.
func (s *SyncSet[T]) AddAll(c ...T) bool {
	isChange := false
	isLocked := false

	read := s.loadReadOnly()
	for _, val := range c {
		if e, ok := read.M[val]; ok {
			_, ok := e.TrySwap(&emptyZero)
			return !ok
		}

		if !isLocked {
			s.mu.Lock()
			defer s.mu.Unlock()
			isLocked = true

			read = s.loadReadOnly()
		}

		if e, ok := read.M[val]; ok {
			if e.UnexpungeLocked() {
				s.dirty[val] = e
			}
			if v := e.SwapLocked(&emptyZero); v != nil {
				continue
			}
		} else if e, ok := s.dirty[val]; ok {
			if v := e.SwapLocked(&emptyZero); v != nil {
				continue
			}
		} else {
			if !read.Amended {
				s.dirtyLocked()
				s.read.Store(&internal.SyncReadOnly[T, empty]{M: read.M, Amended: true})
			}
			s.dirty[val] = internal.NewSyncEntry(emptyZero, nilEmpty)
		}
		isChange = true
	}

	return isChange
}

// Clear removes all of the elements from this collection.
func (s *SyncSet[T]) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()

	read := s.loadReadOnly()
	if read.Amended {
		s.dirty = nil
		s.misses = 0
	}
	read = internal.SyncReadOnly[T, empty]{M: make(map[T]*internal.SyncEntry[empty])}
	copyRead := read
	s.read.Store(&copyRead)
}

// Contains returns true if this collection contains the specified element.
func (s *SyncSet[T]) Contains(e T) bool {
	read := s.loadReadOnly()
	entry, ok := read.M[e]
	if !ok && read.Amended {
		s.mu.Lock()
		read = s.loadReadOnly()
		entry, ok = read.M[e]
		if !ok && read.Amended {
			entry, ok = s.dirty[e]
			s.missLocked()
		}
		s.mu.Unlock()
	}
	if !ok {
		return false
	}
	_, ok = entry.Load(emptyZero)
	return ok
}

// ContainsAll returns true if this collection contains all of the elements in the specified
// collection.
func (s *SyncSet[T]) ContainsAll(c ...T) bool {
	read := s.loadPresentReadOnly()

	for _, val := range c {
		e, ok := read.M[val]
		if !ok {
			return false
		}
		_, ok = e.Load(emptyZero)
		if !ok {
			return false
		}
	}

	return true
}

// Equals compares this collection with the object pass from parameter.
func (s *SyncSet[T]) Equals(o any) bool {
	os, ok := o.(*SyncSet[T])
	if !ok {
		return false
	}

	read := s.loadPresentReadOnly()
	oRead := os.loadPresentReadOnly()

	rc := 0
	oc := 0

	for k, e := range read.M {
		_, ok := e.Load(emptyZero)
		if !ok {
			continue
		}
		rc++

		oe, ok := oRead.M[k]
		if !ok {
			return false
		}
		_, ok = oe.Load(emptyZero)
		if !ok {
			return false
		}
	}

	for _, e := range oRead.M {
		_, ok := e.Load(emptyZero)
		if !ok {
			continue
		}
		oc++
		if rc < oc {
			return false
		}
	}

	return rc == oc
}

// IsEmpty returns true if this collection contains no elements.
func (s *SyncSet[T]) IsEmpty() bool {
	read := s.loadPresentReadOnly()
	if len(read.M) == 0 {
		return true
	}

	for _, e := range read.M {
		_, ok := e.Load(emptyZero)
		if ok {
			return false
		}
	}

	return true
}

// Remove removes the specified element from this collection.
func (s *SyncSet[T]) Remove(k T) bool {
	read := s.loadReadOnly()
	e, ok := read.M[k]
	if !ok && read.Amended {
		s.mu.Lock()
		read = s.loadReadOnly()
		e, ok = read.M[k]
		if !ok && read.Amended {
			e, ok = s.dirty[k]
			s.missLocked()
		}
		s.mu.Unlock()
	}

	if ok {
		_, ok := e.Delete()
		if ok {
			return true
		}
	}

	return false
}

// RemoveAll removes all of the elements in the specified collection from this collection.
func (s *SyncSet[T]) RemoveAll(c ...T) bool {
	read := s.loadReadOnly()
	isChanged := false

	for _, k := range c {
		e, ok := read.M[k]
		if !ok && read.Amended {
			s.mu.Lock()
			read = s.loadReadOnly()
			e, ok = read.M[k]
			if !ok && read.Amended {
				e, ok = s.dirty[k]
				s.missLocked()
			}
			s.mu.Unlock()
		}

		if ok {
			_, ok := e.Delete()
			if ok {
				isChanged = true
			}
		}
	}

	return isChanged
}

// RemoveIf removes all of the elements of this collection that satisfy the given predicate.
func (s *SyncSet[T]) RemoveIf(filter func(T) bool) bool {
	read := s.loadPresentReadOnly()
	isChanged := false

	for k, e := range read.M {
		_, ok := e.Load(emptyZero)
		if !ok {
			continue
		}
		if filter(k) {
			if _, ok := e.Delete(); ok {
				isChanged = true
			}
		}
	}

	return isChanged
}

// RetainAll retains only the elements in this collection that are contained in the specified
// collection.
func (s *SyncSet[T]) RetainAll(c ...T) bool {
	read := s.loadReadOnly()
	isChanged := false

	cSet := NewHashSet[T]()
	cSet.AddAll(c...)

	for k, e := range read.M {
		_, ok := e.Load(emptyZero)
		if !ok {
			continue
		}
		if !cSet.Contains(k) {
			if _, ok := e.Delete(); ok {
				isChanged = true
			}
		}
	}

	return isChanged
}

// Size returns the number of elements in this collection.
func (s *SyncSet[T]) Size() int {
	read := s.loadPresentReadOnly()
	size := 0

	for _, e := range read.M {
		_, ok := e.Load(emptyZero)
		if ok {
			size++
		}
	}

	return size
}

// String returns the string representation of this collection.
func (s *SyncSet[T]) String() string {
	buf := bytes.NewBufferString("set[")
	first := true
	read := s.loadPresentReadOnly()
	for k := range read.M {
		_, ok := read.M[k].Load(emptyZero)
		if ok {
			if !first {
				buf.WriteString(" ")
			}
			first = false
			buf.WriteString(internal.ValueString(k))
		}
	}
	buf.WriteString("]")
	return buf.String()
}

// ToSlice returns a slice containing all of the elements in this collection.
func (s *SyncSet[T]) ToSlice() []T {
	read := s.loadPresentReadOnly()
	slice := make([]T, 0, len(read.M))

	for k, e := range read.M {
		_, ok := e.Load(emptyZero)
		if ok {
			slice = append(slice, k)
		}
	}

	return slice
}

// MarshalJSON marshals the SyncSet as a JSON array.
func (s *SyncSet[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.ToSlice())
}

// UnmarshalJSON unmarshals a JSON array into the SyncSet.
func (s *SyncSet[T]) UnmarshalJSON(b []byte) error {
	var items []T
	if err := json.Unmarshal(b, &items); err != nil {
		return err
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	m := make(map[T]*internal.SyncEntry[empty])
	for _, item := range items {
		m[item] = internal.NewSyncEntry(emptyZero, nilEmpty)
	}
	s.read.Store(&internal.SyncReadOnly[T, empty]{M: m})
	s.dirty = nil
	s.misses = 0

	return nil
}

// Clone returns a copy of this set.
func (s *SyncSet[T]) Clone() collection.Set[T] {
	read := s.loadPresentReadOnly()
	m := make(map[T]*internal.SyncEntry[empty])

	for k, e := range read.M {
		_, ok := e.Load(emptyZero)
		if ok {
			m[k] = internal.NewSyncEntry(emptyZero, nilEmpty)
		}
	}

	clone := NewSyncSet[T]()
	clone.read.Store(&internal.SyncReadOnly[T, empty]{M: m})

	return clone
}

// ForEach performs the given handler for each elements in the collection until all elements
// have been processed or the handler returns an error.
func (s *SyncSet[T]) ForEach(handler func(e T) error) error {
	read := s.loadPresentReadOnly()

	for k, e := range read.M {
		_, ok := e.Load(emptyZero)
		if ok {
			if err := handler(k); err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *SyncSet[T]) loadReadOnly() internal.SyncReadOnly[T, empty] {
	if p := s.read.Load(); p != nil {
		return *p
	}
	return internal.SyncReadOnly[T, empty]{}
}

func (s *SyncSet[T]) loadPresentReadOnly() internal.SyncReadOnly[T, empty] {
	read := s.loadReadOnly()
	if read.Amended {
		s.mu.Lock()
		read = s.loadReadOnly()
		if read.Amended {
			read = internal.SyncReadOnly[T, empty]{M: s.dirty}
			copyRead := read
			s.read.Store(&copyRead)
			s.dirty = nil
			s.misses = 0
		}
		s.mu.Unlock()
	}

	return read
}

func (s *SyncSet[T]) dirtyLocked() {
	if s.dirty != nil {
		return
	}

	read := s.loadReadOnly()
	s.dirty = make(map[T]*internal.SyncEntry[empty], len(read.M))
	for k, e := range read.M {
		if !e.TryExpungeLocked() {
			s.dirty[k] = e
		}
	}
}

func (s *SyncSet[T]) missLocked() {
	s.misses++
	if s.misses < len(s.dirty) {
		return
	}

	s.read.Store(&internal.SyncReadOnly[T, empty]{M: s.dirty})
	s.dirty = nil
	s.misses = 0
}
