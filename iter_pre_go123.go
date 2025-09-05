//go:build !go1.23

package collection

type Iterable[T any] interface {
	// Iter returns a channel of all elements in this collection.
	Iter() <-chan T
}

type Iterable2[K comparable, V any] interface {
}

type DictIter[K comparable, V any] interface {
	// KeysIter returns a channel over the keys in the dictionary.
	KeysIter() <-chan K
	// ValuesIter returns a channel over the values in the dictionary.
	ValuesIter() <-chan V
}
