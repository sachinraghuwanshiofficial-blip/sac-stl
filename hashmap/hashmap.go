// Package hashmap provides a generic unordered hash map equivalent to
// C++ std::unordered_map.
//
// It wraps Go's built-in map with an ergonomic API that mirrors the rest of
// this library. The underlying map gives the same O(1) amortized guarantees.
//
// Time complexity:
//   Put/Get/Delete  O(1) amortized
//   Contains        O(1) amortized
//   Len             O(1)
//   Range           O(n)
package hashmap

// Map is a generic unordered map with comparable keys.
type Map[K comparable, V any] struct {
	data map[K]V
}

// New returns an empty Map.
func New[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{data: make(map[K]V)}
}

// NewWithCapacity returns a Map pre-sized for the given number of entries.
func NewWithCapacity[K comparable, V any](cap int) *Map[K, V] {
	return &Map[K, V]{data: make(map[K]V, cap)}
}

// Put inserts or updates the key-value pair. O(1) amortized.
func (m *Map[K, V]) Put(key K, val V) {
	m.data[key] = val
}

// Get returns the value for key and true if found. O(1) amortized.
func (m *Map[K, V]) Get(key K) (V, bool) {
	v, ok := m.data[key]
	return v, ok
}

// GetOrDefault returns the value for key, or def if not present.
func (m *Map[K, V]) GetOrDefault(key K, def V) V {
	if v, ok := m.data[key]; ok {
		return v
	}
	return def
}

// Contains reports whether key is present. O(1) amortized.
func (m *Map[K, V]) Contains(key K) bool {
	_, ok := m.data[key]
	return ok
}

// Delete removes the key from the map. O(1) amortized.
func (m *Map[K, V]) Delete(key K) {
	delete(m.data, key)
}

// Len returns the number of key-value pairs. O(1).
func (m *Map[K, V]) Len() int { return len(m.data) }

// Size is an alias for Len, matching C++ std::unordered_map::size().
func (m *Map[K, V]) Size() int { return len(m.data) }

// Empty reports whether the map has no entries.
func (m *Map[K, V]) Empty() bool { return len(m.data) == 0 }

// IsEmpty is an alias for Empty.
func (m *Map[K, V]) IsEmpty() bool { return len(m.data) == 0 }

// Insert is an alias for Put, matching C++ std::unordered_map::insert().
func (m *Map[K, V]) Insert(key K, val V) { m.Put(key, val) }

// Find returns the value for key and true if found (mirrors C++ std::unordered_map::find()).
func (m *Map[K, V]) Find(key K) (V, bool) { return m.Get(key) }

// Erase is an alias for Delete, matching C++ std::unordered_map::erase().
func (m *Map[K, V]) Erase(key K) { m.Delete(key) }

// Clear removes all entries.
func (m *Map[K, V]) Clear() {
	m.data = make(map[K]V)
}

// Range iterates over all key-value pairs (unordered).
// Stop early by returning false from fn.
func (m *Map[K, V]) Range(fn func(key K, val V) bool) {
	for k, v := range m.data {
		if !fn(k, v) {
			return
		}
	}
}

// Keys returns all keys in unspecified order.
func (m *Map[K, V]) Keys() []K {
	keys := make([]K, 0, len(m.data))
	for k := range m.data {
		keys = append(keys, k)
	}
	return keys
}

// Values returns all values in unspecified order.
func (m *Map[K, V]) Values() []V {
	vals := make([]V, 0, len(m.data))
	for _, v := range m.data {
		vals = append(vals, v)
	}
	return vals
}
