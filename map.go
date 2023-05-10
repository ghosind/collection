package collection

type Map[K comparable, V any] interface {
	// Clear removes all key-value pairs in this map.
	Clear()

	// Clone returns a copy of this map.
	Clone() Map[K, V]

	// ContainsKey returns true if this map contains a key-value pair with the specified key.
	ContainsKey(k K) bool

	// Equals compares this map with the object pass from parameter.
	Equals(o any) bool

	// ForEach performs the given handler for each key-value pairs in the map until all pairs have
	// been processed or the handler returns an error.
	ForEach(handler func(k K, v V) error) error

	// Get returns the value which associated to the specified key.
	Get(k K) (V, bool)

	// GetDefault returns the value associated with the specified key, and returns the default value
	// if this map contains no pair with the key.
	GetDefault(k K, defaultVal V) V

	// IsEmpty returns true if this map is empty.
	IsEmpty() bool

	// Keys returns a slice that contains all the keys in this map.
	Keys() []K

	// Put associate the specified value with the specified key in this map.
	Put(k K, v V) V

	// Remove removes the key-value pair with the specified key.
	Remove(k K) V

	// Replace replaces the value for the specified key only if it is currently in this map.
	Replace(k K, v V) (V, bool)

	// Size returns the number of key-value pairs in this map.
	Size() int

	// Values returns a slice that contains all the values in this map.
	Values() []V
}
