package internal

import "sync/atomic"

type SyncReadOnly[K comparable, V any] struct {
	M       map[K]*SyncEntry[V]
	Amended bool
}

type SyncEntry[T any] struct {
	p        atomic.Pointer[T]
	expunged *T
}

func NewSyncEntry[T any](v T, expunged *T) *SyncEntry[T] {
	e := new(SyncEntry[T])
	e.p.Store(&v)
	e.expunged = expunged
	return e
}

func (e *SyncEntry[T]) Load(val T) (value T, ok bool) {
	p := e.p.Load()
	if p == nil || p == e.expunged {
		return val, false
	}
	return *p, true
}

func (e *SyncEntry[T]) TrySwap(val *T) (*T, bool) {
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

func (e *SyncEntry[T]) Delete() (*T, bool) {
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

func (e *SyncEntry[T]) UnexpungeLocked() bool {
	return e.p.CompareAndSwap(e.expunged, nil)
}

func (e *SyncEntry[T]) SwapLocked(v *T) *T {
	return e.p.Swap(v)
}

func (e *SyncEntry[T]) TryExpungeLocked() bool {
	p := e.p.Load()
	for p == nil {
		if e.p.CompareAndSwap(nil, e.expunged) {
			return true
		}
		p = e.p.Load()
	}
	return p == e.expunged
}
