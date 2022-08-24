package collection

type Map[K comparable, V any] interface {
	// Clear removes all key-value pairs in this map.
	Clear()

	// ContainsKey returns true if this map contains a key-value pair with the specified key.
	ContainsKey(k K) bool

	// Equals compares this map with the object pass from parameter.
	Equals(o any) bool

	// ForEach performs the given handler for each key-value pairs in the map until all pairs have
	// been processed or the handler returns an error.
	ForEach(handler func(k K, v V) error) error

	// Get returns the value which associated to the specified key.
	Get(k K) (V, bool)

	// IsEmpty returns true if this map is empty.
	IsEmpty() bool

	// Put associate the specified value with the specified key in this map.
	Put(k K, v V) V

	// Remove removes the key-value pair with the specified key.
	Remove(k K) V

	// Size returns the number of key-value pairs in this map.
	Size() int
}
