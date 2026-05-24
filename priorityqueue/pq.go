// Package priorityqueue provides a generic binary-heap priority queue
// equivalent to C++ std::priority_queue.
//
// The ordering is controlled by a user-supplied less function:
//   - Min-heap (smallest element on top): less = func(a, b T) bool { return a < b }
//   - Max-heap (largest element on top):  less = func(a, b T) bool { return a > b }
//
// Time complexity:
//   Push  O(log n)
//   Pop   O(log n)
//   Top   O(1)
//   Len   O(1)
package priorityqueue

// PriorityQueue is a generic binary heap. Elements with higher priority
// (as defined by less) are served first.
type PriorityQueue[T any] struct {
	data []T
	less func(a, b T) bool
}

// New returns an empty PriorityQueue with the given comparator.
// less(a, b) should return true when a should be served before b.
func New[T any](less func(a, b T) bool) *PriorityQueue[T] {
	return &PriorityQueue[T]{less: less}
}

// NewWithCapacity returns a PriorityQueue pre-allocated to the given capacity.
func NewWithCapacity[T any](cap int, less func(a, b T) bool) *PriorityQueue[T] {
	return &PriorityQueue[T]{data: make([]T, 0, cap), less: less}
}

// NewFrom builds a PriorityQueue from an existing slice in O(n) via heapify.
func NewFrom[T any](s []T, less func(a, b T) bool) *PriorityQueue[T] {
	pq := &PriorityQueue[T]{data: make([]T, len(s)), less: less}
	copy(pq.data, s)
	// bottom-up heapify
	for i := len(pq.data)/2 - 1; i >= 0; i-- {
		pq.down(i)
	}
	return pq
}

// Push inserts val into the queue. O(log n).
func (pq *PriorityQueue[T]) Push(val T) {
	pq.data = append(pq.data, val)
	pq.up(len(pq.data) - 1)
}

// Pop removes and returns the highest-priority element. O(log n).
// Returns zero value and false if empty.
func (pq *PriorityQueue[T]) Pop() (T, bool) {
	if len(pq.data) == 0 {
		var zero T
		return zero, false
	}
	top := pq.data[0]
	n := len(pq.data) - 1
	pq.data[0] = pq.data[n]
	var zero T
	pq.data[n] = zero // allow GC
	pq.data = pq.data[:n]
	if n > 0 {
		pq.down(0)
	}
	return top, true
}

// Top returns the highest-priority element without removing it. O(1).
// Returns zero value and false if empty.
func (pq *PriorityQueue[T]) Top() (T, bool) {
	if len(pq.data) == 0 {
		var zero T
		return zero, false
	}
	return pq.data[0], true
}

// Len returns the number of elements.
func (pq *PriorityQueue[T]) Len() int { return len(pq.data) }

// Size is an alias for Len, matching C++ std::priority_queue::size().
func (pq *PriorityQueue[T]) Size() int { return len(pq.data) }

// Empty reports whether the queue has no elements.
func (pq *PriorityQueue[T]) Empty() bool { return len(pq.data) == 0 }

// IsEmpty is an alias for Empty.
func (pq *PriorityQueue[T]) IsEmpty() bool { return len(pq.data) == 0 }

// Clear removes all elements.
func (pq *PriorityQueue[T]) Clear() {
	var zero T
	for i := range pq.data {
		pq.data[i] = zero
	}
	pq.data = pq.data[:0]
}

func (pq *PriorityQueue[T]) up(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if !pq.less(pq.data[i], pq.data[parent]) {
			break
		}
		pq.data[i], pq.data[parent] = pq.data[parent], pq.data[i]
		i = parent
	}
}

func (pq *PriorityQueue[T]) down(i int) {
	n := len(pq.data)
	for {
		left := 2*i + 1
		if left >= n {
			break
		}
		j := left
		if right := left + 1; right < n && pq.less(pq.data[right], pq.data[left]) {
			j = right
		}
		if !pq.less(pq.data[j], pq.data[i]) {
			break
		}
		pq.data[i], pq.data[j] = pq.data[j], pq.data[i]
		i = j
	}
}
