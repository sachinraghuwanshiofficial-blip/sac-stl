// Package hashset provides a generic unordered set equivalent to
// C++ std::unordered_set.
//
// Backed by Go's built-in map. Elements are unique and unordered.
//
// Time complexity:
//   Insert/Delete/Contains  O(1) amortized
//   Len                     O(1)
//   Range                   O(n)
package hashset

// Set is a generic unordered set with comparable elements.
type Set[T comparable] struct {
	data map[T]struct{}
}

// New returns an empty Set.
func New[T comparable]() *Set[T] {
	return &Set[T]{data: make(map[T]struct{})}
}

// NewWithCapacity returns a Set pre-sized for n elements.
func NewWithCapacity[T comparable](n int) *Set[T] {
	return &Set[T]{data: make(map[T]struct{}, n)}
}

// NewFrom creates a Set from a slice.
func NewFrom[T comparable](s []T) *Set[T] {
	set := NewWithCapacity[T](len(s))
	for _, v := range s {
		set.data[v] = struct{}{}
	}
	return set
}

// Insert adds val to the set. O(1) amortized.
func (s *Set[T]) Insert(val T) {
	s.data[val] = struct{}{}
}

// Delete removes val from the set. O(1) amortized.
func (s *Set[T]) Delete(val T) {
	delete(s.data, val)
}

// Contains reports whether val is in the set. O(1) amortized.
func (s *Set[T]) Contains(val T) bool {
	_, ok := s.data[val]
	return ok
}

// Len returns the number of elements. O(1).
func (s *Set[T]) Len() int { return len(s.data) }

// Size is an alias for Len, matching C++ std::unordered_set::size().
func (s *Set[T]) Size() int { return len(s.data) }

// Empty reports whether the set has no elements.
func (s *Set[T]) Empty() bool { return len(s.data) == 0 }

// IsEmpty is an alias for Empty.
func (s *Set[T]) IsEmpty() bool { return len(s.data) == 0 }

// Erase is an alias for Delete, matching C++ std::unordered_set::erase().
func (s *Set[T]) Erase(val T) { delete(s.data, val) }

// Count returns 1 if val is present, 0 otherwise (mirrors C++ std::unordered_set::count()).
func (s *Set[T]) Count(val T) int {
	if _, ok := s.data[val]; ok {
		return 1
	}
	return 0
}

// Clear removes all elements.
func (s *Set[T]) Clear() {
	s.data = make(map[T]struct{})
}

// Range iterates over all elements in unspecified order.
// Stop early by returning false from fn.
func (s *Set[T]) Range(fn func(val T) bool) {
	for v := range s.data {
		if !fn(v) {
			return
		}
	}
}

// Slice returns all elements as an unordered slice.
func (s *Set[T]) Slice() []T {
	out := make([]T, 0, len(s.data))
	for v := range s.data {
		out = append(out, v)
	}
	return out
}

// Union returns a new set containing elements from both s and other.
func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	result := NewWithCapacity[T](s.Len() + other.Len())
	for v := range s.data {
		result.data[v] = struct{}{}
	}
	for v := range other.data {
		result.data[v] = struct{}{}
	}
	return result
}

// Intersection returns a new set containing elements present in both sets.
func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	result := New[T]()
	// iterate over the smaller set
	small, large := s, other
	if small.Len() > large.Len() {
		small, large = large, small
	}
	for v := range small.data {
		if large.Contains(v) {
			result.data[v] = struct{}{}
		}
	}
	return result
}

// Difference returns a new set with elements in s that are not in other.
func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	result := New[T]()
	for v := range s.data {
		if !other.Contains(v) {
			result.data[v] = struct{}{}
		}
	}
	return result
}
