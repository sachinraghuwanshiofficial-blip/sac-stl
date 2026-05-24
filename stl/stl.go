// Package stl is a unified entry point for go-stl.
//
// It re-exports all container types and provides C++-style free functions so
// you can write code that reads almost identically to C++ STL usage:
//
//	v := stl.MakeVector[int]()
//	stl.PushBack(v, 1)
//	stl.PushBack(v, 2)
//	val, ok := stl.PopBack(v)
//
//	s := stl.MakeStack[string]()
//	stl.Push(s, "hello")
//	top, ok := stl.Top(s)
//	stl.Pop(s)
//
//	pq := stl.MakeMaxHeap[int]()
//	stl.PushPQ(pq, 42)
//	best, ok := stl.TopPQ(pq)
//
// Each function is a thin, zero-overhead wrapper — the compiler inlines them.
package stl

import (
	"github.com/sachin/go-stl/hashmap"
	"github.com/sachin/go-stl/hashset"
	"github.com/sachin/go-stl/priorityqueue"
	"github.com/sachin/go-stl/queue"
	"github.com/sachin/go-stl/stack"
	"github.com/sachin/go-stl/treemap"
	"github.com/sachin/go-stl/treeset"
	"github.com/sachin/go-stl/vector"
)

// ── Type aliases ─────────────────────────────────────────────────────────────
// Re-export all container types so a single import suffices.

type Vector[T any] = vector.Vector[T]
type Stack[T any] = stack.Stack[T]
type Queue[T any] = queue.Queue[T]
type PriorityQueue[T any] = priorityqueue.PriorityQueue[T]
type TreeMap[K, V any] = treemap.Map[K, V]
type HashMap[K comparable, V any] = hashmap.Map[K, V]
type TreeSet[T any] = treeset.Set[T]
type HashSet[T comparable] = hashset.Set[T]

// ── Constructors ─────────────────────────────────────────────────────────────

// MakeVector returns a new empty Vector (mirrors std::vector<T> v;).
func MakeVector[T any]() *Vector[T] { return vector.New[T]() }

// MakeVectorN returns a Vector resized to n zero-valued elements.
func MakeVectorN[T any](n int) *Vector[T] {
	v := vector.NewWithCapacity[T](n)
	v.Resize(n)
	return v
}

// MakeVectorFrom creates a Vector from a slice (mirrors initializer-list ctor).
func MakeVectorFrom[T any](s []T) *Vector[T] { return vector.NewFrom(s) }

// MakeStack returns a new empty Stack (mirrors std::stack<T> s;).
func MakeStack[T any]() *Stack[T] { return stack.New[T]() }

// MakeQueue returns a new empty Queue (mirrors std::queue<T> q;).
func MakeQueue[T any]() *Queue[T] { return queue.New[T]() }

// MakeMinHeap returns a min-heap PriorityQueue ordered by less.
// Usage: pq := stl.MakeMinHeap(func(a, b int) bool { return a < b })
func MakeMinHeap[T any](less func(a, b T) bool) *PriorityQueue[T] {
	return priorityqueue.New(less)
}

// MakeMaxHeap returns a max-heap PriorityQueue ordered by less (pass a > b).
// Usage: pq := stl.MakeMaxHeap(func(a, b int) bool { return a > b })
func MakeMaxHeap[T any](less func(a, b T) bool) *PriorityQueue[T] {
	return priorityqueue.New(less)
}

// MakeTreeMap returns a new ordered TreeMap (mirrors std::map<K,V> m;).
func MakeTreeMap[K, V any](less func(a, b K) bool) *TreeMap[K, V] {
	return treemap.New[K, V](less)
}

// MakeHashMap returns a new unordered HashMap (mirrors std::unordered_map<K,V> m;).
func MakeHashMap[K comparable, V any]() *HashMap[K, V] { return hashmap.New[K, V]() }

// MakeTreeSet returns a new ordered TreeSet (mirrors std::set<T> s;).
func MakeTreeSet[T any](less func(a, b T) bool) *TreeSet[T] { return treeset.New(less) }

// MakeHashSet returns a new unordered HashSet (mirrors std::unordered_set<T> s;).
func MakeHashSet[T comparable]() *HashSet[T] { return hashset.New[T]() }

// ── Vector free functions ─────────────────────────────────────────────────────

// PushBack appends val to the end of v. O(1) amortized.
func PushBack[T any](v *Vector[T], val T) { v.PushBack(val) }

// PopBack removes and returns the last element of v.
func PopBack[T any](v *Vector[T]) (T, bool) { return v.PopBack() }

// Front returns the first element of v.
func Front[T any](v *Vector[T]) (T, bool) { return v.Front() }

// Back returns the last element of v.
func Back[T any](v *Vector[T]) (T, bool) { return v.Back() }

// At returns the element at index i of v (panics on out-of-bounds).
func At[T any](v *Vector[T], i int) T { return v.At(i) }

// Insert inserts val before index i in v. O(n).
func Insert[T any](v *Vector[T], i int, val T) { v.Insert(i, val) }

// Erase removes the element at index i from v. O(n).
func Erase[T any](v *Vector[T], i int) { v.Erase(i) }

// Size returns the number of elements in v.
func Size[T any](v *Vector[T]) int { return v.Len() }

// Empty reports whether v has no elements.
func Empty[T any](v *Vector[T]) bool { return v.Empty() }

// ── Stack free functions ──────────────────────────────────────────────────────

// Push pushes val onto s. O(1) amortized.
func Push[T any](s *Stack[T], val T) { s.Push(val) }

// Pop removes and returns the top of s.
func Pop[T any](s *Stack[T]) (T, bool) { return s.Pop() }

// Top returns the top element of s without removing it.
func Top[T any](s *Stack[T]) (T, bool) { return s.Top() }

// StackSize returns the number of elements in s.
func StackSize[T any](s *Stack[T]) int { return s.Len() }

// IsEmpty reports whether s has no elements.
func IsEmpty[T any](s *Stack[T]) bool { return s.Empty() }

// ── Queue free functions ──────────────────────────────────────────────────────

// Enqueue pushes val to the back of q. O(1) amortized.
func Enqueue[T any](q *Queue[T], val T) { q.Push(val) }

// Dequeue removes and returns the front of q.
func Dequeue[T any](q *Queue[T]) (T, bool) { return q.Pop() }

// QueueFront returns the front element of q without removing it.
func QueueFront[T any](q *Queue[T]) (T, bool) { return q.Front() }

// QueueBack returns the back element of q without removing it.
func QueueBack[T any](q *Queue[T]) (T, bool) { return q.Back() }

// QueueSize returns the number of elements in q.
func QueueSize[T any](q *Queue[T]) int { return q.Len() }

// ── PriorityQueue free functions ──────────────────────────────────────────────

// PushPQ pushes val into the priority queue pq. O(log n).
func PushPQ[T any](pq *PriorityQueue[T], val T) { pq.Push(val) }

// PopPQ removes and returns the highest-priority element. O(log n).
func PopPQ[T any](pq *PriorityQueue[T]) (T, bool) { return pq.Pop() }

// TopPQ returns the highest-priority element without removing it. O(1).
func TopPQ[T any](pq *PriorityQueue[T]) (T, bool) { return pq.Top() }

// ── Map free functions ────────────────────────────────────────────────────────

// MapInsert inserts or updates key→val in m.
func MapInsert[K, V any](m *TreeMap[K, V], key K, val V) { m.Put(key, val) }

// MapFind returns the value for key in m.
func MapFind[K, V any](m *TreeMap[K, V], key K) (V, bool) { return m.Get(key) }

// MapErase removes key from m.
func MapErase[K, V any](m *TreeMap[K, V], key K) { m.Delete(key) }

// MapSize returns the number of entries in m.
func MapSize[K, V any](m *TreeMap[K, V]) int { return m.Len() }

// HMapInsert inserts or updates key→val in a HashMap.
func HMapInsert[K comparable, V any](m *HashMap[K, V], key K, val V) { m.Put(key, val) }

// HMapFind returns the value for key in a HashMap.
func HMapFind[K comparable, V any](m *HashMap[K, V], key K) (V, bool) { return m.Get(key) }

// HMapErase removes key from a HashMap.
func HMapErase[K comparable, V any](m *HashMap[K, V], key K) { m.Delete(key) }

// ── Set free functions ────────────────────────────────────────────────────────

// SetInsert inserts val into the ordered TreeSet.
func SetInsert[T any](s *TreeSet[T], val T) { s.Insert(val) }

// SetErase removes val from the ordered TreeSet.
func SetErase[T any](s *TreeSet[T], val T) { s.Delete(val) }

// SetContains reports whether val is in the TreeSet.
func SetContains[T any](s *TreeSet[T], val T) bool { return s.Contains(val) }

// HSetInsert inserts val into the HashSet.
func HSetInsert[T comparable](s *HashSet[T], val T) { s.Insert(val) }

// HSetErase removes val from the HashSet.
func HSetErase[T comparable](s *HashSet[T], val T) { s.Delete(val) }

// HSetContains reports whether val is in the HashSet.
func HSetContains[T comparable](s *HashSet[T], val T) bool { return s.Contains(val) }
