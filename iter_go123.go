//go:build go1.23

package collection

import "iter"

type Iterable[T any] interface {
	// Iter returns a channel of all elements in this collection.
	Iter() iter.Seq[T]
}
