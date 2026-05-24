// Package stack provides a generic LIFO stack equivalent to C++ std::stack.
//
// Time complexity:
//   Push  O(1) amortized
//   Pop   O(1)
//   Top   O(1)
//   Len   O(1)
package stack

// Stack is a generic LIFO container backed by a slice.
type Stack[T any] struct {
	data []T
}

// New returns an empty Stack.
func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

// NewWithCapacity returns a Stack pre-allocated to the given capacity.
func NewWithCapacity[T any](cap int) *Stack[T] {
	return &Stack[T]{data: make([]T, 0, cap)}
}

// Push pushes val onto the top of the stack. O(1) amortized.
func (s *Stack[T]) Push(val T) {
	s.data = append(s.data, val)
}

// Pop removes and returns the top element.
// Returns zero value and false if empty.
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.data) == 0 {
		var zero T
		return zero, false
	}
	top := s.data[len(s.data)-1]
	var zero T
	s.data[len(s.data)-1] = zero // allow GC
	s.data = s.data[:len(s.data)-1]
	return top, true
}

// Top returns the top element without removing it.
// Returns zero value and false if empty.
func (s *Stack[T]) Top() (T, bool) {
	if len(s.data) == 0 {
		var zero T
		return zero, false
	}
	return s.data[len(s.data)-1], true
}

// Len returns the number of elements.
func (s *Stack[T]) Len() int { return len(s.data) }

// Size is an alias for Len, matching C++ std::stack::size().
func (s *Stack[T]) Size() int { return len(s.data) }

// Empty reports whether the stack has no elements.
func (s *Stack[T]) Empty() bool { return len(s.data) == 0 }

// IsEmpty is an alias for Empty.
func (s *Stack[T]) IsEmpty() bool { return len(s.data) == 0 }

// Clear removes all elements.
func (s *Stack[T]) Clear() {
	var zero T
	for i := range s.data {
		s.data[i] = zero
	}
	s.data = s.data[:0]
}
