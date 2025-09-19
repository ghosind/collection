//go:build go1.23

package collection

import "iter"

type Iterable[T any] interface {
	// Iter returns an iterator of all elements in this collection.
	Iter() iter.Seq[T]
}

type Iterable2[K comparable, V any] interface {
	// Iter returns an iterator of all key-value pairs in this collection.
	Iter() iter.Seq2[K, V]
}

type DictIter[K comparable, V any] interface {
	// KeysIter returns an iterator over the keys in the dictionary.
	KeysIter() iter.Seq[K]
	// ValuesIter returns an iterator over the values in the dictionary.
	ValuesIter() iter.Seq[V]
}
