// Package treeset provides a generic ordered set equivalent to C++ std::set.
//
// Backed by the treemap LLRB implementation. Elements are unique and
// iterated in sorted order.
//
// Time complexity:
//   Insert/Delete/Contains  O(log n)
//   Min/Max                 O(log n)
//   Floor/Ceiling           O(log n)
//   Len                     O(1)
//   Range (in-order)        O(n)
package treeset

import "github.com/sachin/go-stl/treemap"

// Set is a generic ordered set.
type Set[T any] struct {
	m *treemap.Map[T, struct{}]
}

// New returns an empty Set with the given comparator.
func New[T any](less func(a, b T) bool) *Set[T] {
	return &Set[T]{m: treemap.New[T, struct{}](less)}
}

// Insert adds val to the set. O(log n).
func (s *Set[T]) Insert(val T) {
	s.m.Put(val, struct{}{})
}

// Delete removes val from the set. O(log n).
func (s *Set[T]) Delete(val T) {
	s.m.Delete(val)
}

// Contains reports whether val is in the set. O(log n).
func (s *Set[T]) Contains(val T) bool {
	return s.m.Contains(val)
}

// Len returns the number of elements. O(1).
func (s *Set[T]) Len() int { return s.m.Len() }

// Size is an alias for Len, matching C++ std::set::size().
func (s *Set[T]) Size() int { return s.m.Len() }

// Empty reports whether the set has no elements.
func (s *Set[T]) Empty() bool { return s.m.Empty() }

// IsEmpty is an alias for Empty.
func (s *Set[T]) IsEmpty() bool { return s.m.Empty() }

// Erase is an alias for Delete, matching C++ std::set::erase().
func (s *Set[T]) Erase(val T) { s.Delete(val) }

// Count returns 1 if val is present, 0 otherwise (mirrors C++ std::set::count()).
func (s *Set[T]) Count(val T) int {
	if s.Contains(val) {
		return 1
	}
	return 0
}

// Min returns the smallest element. O(log n).
func (s *Set[T]) Min() (T, bool) {
	k, _, ok := s.m.Min()
	return k, ok
}

// Max returns the largest element. O(log n).
func (s *Set[T]) Max() (T, bool) {
	k, _, ok := s.m.Max()
	return k, ok
}

// Floor returns the largest element ≤ val. O(log n).
func (s *Set[T]) Floor(val T) (T, bool) {
	k, _, ok := s.m.Floor(val)
	return k, ok
}

// Ceiling returns the smallest element ≥ val. O(log n).
func (s *Set[T]) Ceiling(val T) (T, bool) {
	k, _, ok := s.m.Ceiling(val)
	return k, ok
}

// Range iterates over all elements in ascending order.
// Stop early by returning false from fn.
func (s *Set[T]) Range(fn func(val T) bool) {
	s.m.Range(func(k T, _ struct{}) bool {
		return fn(k)
	})
}

// Slice returns all elements as a sorted slice.
func (s *Set[T]) Slice() []T {
	return s.m.Keys()
}
