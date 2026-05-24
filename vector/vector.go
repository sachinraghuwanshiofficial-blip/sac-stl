// Package vector provides a generic dynamic array equivalent to C++ std::vector.
//
// Time complexity:
//   PushBack   O(1) amortized  (doubles capacity on growth, like C++)
//   PopBack    O(1)
//   At/Set     O(1)
//   Insert     O(n)
//   Erase      O(n)
//   Find       O(n)
//   Reserve    O(n)
//   Len/Cap    O(1)
package vector

// Vector is a generic dynamic array. The zero value is ready to use.
type Vector[T any] struct {
	data []T
}

// New returns an empty Vector.
func New[T any]() *Vector[T] {
	return &Vector[T]{}
}

// NewWithCapacity returns a Vector pre-allocated to the given capacity.
func NewWithCapacity[T any](cap int) *Vector[T] {
	return &Vector[T]{data: make([]T, 0, cap)}
}

// NewFrom creates a Vector from an existing slice (copied).
func NewFrom[T any](s []T) *Vector[T] {
	d := make([]T, len(s))
	copy(d, s)
	return &Vector[T]{data: d}
}

// PushBack appends v to the end. O(1) amortized.
func (v *Vector[T]) PushBack(val T) {
	v.data = append(v.data, val)
}

// PopBack removes and returns the last element.
// Returns the zero value and false if empty.
func (v *Vector[T]) PopBack() (T, bool) {
	if len(v.data) == 0 {
		var zero T
		return zero, false
	}
	last := v.data[len(v.data)-1]
	var zero T
	v.data[len(v.data)-1] = zero // clear for GC
	v.data = v.data[:len(v.data)-1]
	return last, true
}

// Front returns the first element without removing it.
func (v *Vector[T]) Front() (T, bool) {
	if len(v.data) == 0 {
		var zero T
		return zero, false
	}
	return v.data[0], true
}

// Back returns the last element without removing it.
func (v *Vector[T]) Back() (T, bool) {
	if len(v.data) == 0 {
		var zero T
		return zero, false
	}
	return v.data[len(v.data)-1], true
}

// At returns the element at index i. Panics if out of bounds (mirrors C++ .at()).
func (v *Vector[T]) At(i int) T {
	return v.data[i]
}

// Set sets the element at index i.
func (v *Vector[T]) Set(i int, val T) {
	v.data[i] = val
}

// Insert inserts val before index i. O(n).
func (v *Vector[T]) Insert(i int, val T) {
	v.data = append(v.data, val) // grow by 1
	copy(v.data[i+1:], v.data[i:])
	v.data[i] = val
}

// Erase removes the element at index i. O(n).
func (v *Vector[T]) Erase(i int) {
	var zero T
	copy(v.data[i:], v.data[i+1:])
	v.data[len(v.data)-1] = zero
	v.data = v.data[:len(v.data)-1]
}

// Len returns the number of elements.
func (v *Vector[T]) Len() int { return len(v.data) }

// Size is an alias for Len, matching C++ std::vector::size().
func (v *Vector[T]) Size() int { return len(v.data) }

// Cap returns the current allocated capacity.
func (v *Vector[T]) Cap() int { return cap(v.data) }

// Empty reports whether the vector has no elements.
func (v *Vector[T]) Empty() bool { return len(v.data) == 0 }

// Reserve ensures capacity for at least n elements without changing Len.
func (v *Vector[T]) Reserve(n int) {
	if cap(v.data) >= n {
		return
	}
	next := make([]T, len(v.data), n)
	copy(next, v.data)
	v.data = next
}

// Resize changes the length to n. New elements are zero-valued.
func (v *Vector[T]) Resize(n int) {
	if n <= len(v.data) {
		var zero T
		for i := n; i < len(v.data); i++ {
			v.data[i] = zero
		}
		v.data = v.data[:n]
		return
	}
	if n <= cap(v.data) {
		v.data = v.data[:n]
		return
	}
	next := make([]T, n)
	copy(next, v.data)
	v.data = next
}

// Clear removes all elements but retains allocated capacity.
func (v *Vector[T]) Clear() {
	var zero T
	for i := range v.data {
		v.data[i] = zero
	}
	v.data = v.data[:0]
}

// Slice returns a copy of the underlying data as a plain Go slice.
func (v *Vector[T]) Slice() []T {
	out := make([]T, len(v.data))
	copy(out, v.data)
	return out
}

// Range calls fn for each element in order. Stop early by returning false.
func (v *Vector[T]) Range(fn func(i int, val T) bool) {
	for i, val := range v.data {
		if !fn(i, val) {
			return
		}
	}
}
