// Package queue provides a generic FIFO queue equivalent to C++ std::queue.
//
// Backed by a power-of-two ring buffer so both Push and Pop are O(1) amortized
// with no shifting — unlike a naive slice-based approach that would be O(n) on
// front dequeue.
//
// Time complexity:
//   Push   O(1) amortized
//   Pop    O(1)
//   Front  O(1)
//   Back   O(1)
//   Len    O(1)
package queue

const minCap = 4

// Queue is a generic FIFO container backed by a ring buffer.
type Queue[T any] struct {
	buf        []T
	head, tail int
	count      int
}

// New returns an empty Queue.
func New[T any]() *Queue[T] {
	return &Queue[T]{buf: make([]T, minCap)}
}

// NewWithCapacity returns a Queue pre-allocated to at least the given capacity.
func NewWithCapacity[T any](cap int) *Queue[T] {
	c := nextPow2(cap)
	if c < minCap {
		c = minCap
	}
	return &Queue[T]{buf: make([]T, c)}
}

// Push enqueues val at the back. O(1) amortized.
func (q *Queue[T]) Push(val T) {
	if q.count == len(q.buf) {
		q.grow()
	}
	q.buf[q.tail] = val
	q.tail = (q.tail + 1) & (len(q.buf) - 1)
	q.count++
}

// Pop dequeues and returns the front element.
// Returns zero value and false if empty.
func (q *Queue[T]) Pop() (T, bool) {
	if q.count == 0 {
		var zero T
		return zero, false
	}
	val := q.buf[q.head]
	var zero T
	q.buf[q.head] = zero // allow GC
	q.head = (q.head + 1) & (len(q.buf) - 1)
	q.count--
	return val, true
}

// Front returns the front element without removing it.
func (q *Queue[T]) Front() (T, bool) {
	if q.count == 0 {
		var zero T
		return zero, false
	}
	return q.buf[q.head], true
}

// Back returns the back element without removing it.
func (q *Queue[T]) Back() (T, bool) {
	if q.count == 0 {
		var zero T
		return zero, false
	}
	idx := (q.tail - 1 + len(q.buf)) & (len(q.buf) - 1)
	return q.buf[idx], true
}

// Len returns the number of elements.
func (q *Queue[T]) Len() int { return q.count }

// Size is an alias for Len, matching C++ std::queue::size().
func (q *Queue[T]) Size() int { return q.count }

// Empty reports whether the queue has no elements.
func (q *Queue[T]) Empty() bool { return q.count == 0 }

// IsEmpty is an alias for Empty.
func (q *Queue[T]) IsEmpty() bool { return q.count == 0 }

// Clear removes all elements but retains allocated capacity.
func (q *Queue[T]) Clear() {
	var zero T
	for i := range q.buf {
		q.buf[i] = zero
	}
	q.head, q.tail, q.count = 0, 0, 0
}

// grow doubles the ring buffer capacity.
func (q *Queue[T]) grow() {
	newCap := len(q.buf) * 2
	newBuf := make([]T, newCap)
	if q.head < q.tail {
		copy(newBuf, q.buf[q.head:q.tail])
	} else {
		n := copy(newBuf, q.buf[q.head:])
		copy(newBuf[n:], q.buf[:q.tail])
	}
	q.head = 0
	q.tail = q.count
	q.buf = newBuf
}

func nextPow2(n int) int {
	if n <= 0 {
		return minCap
	}
	n--
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	n |= n >> 32
	return n + 1
}
