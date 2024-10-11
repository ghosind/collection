//go:build go1.23

package set

import "iter"

// Iter returns a channel of all elements in this set.
func (set *ConcurrentHashSet[T]) Iter() iter.Seq[T] {
	return func(yield func(T) bool) {
		set.mutex.RLock()
		defer set.mutex.RUnlock()

		for e := range *set.data {
			if !yield(e) {
				break
			}
		}
	}
}
