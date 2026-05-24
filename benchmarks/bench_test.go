// Package benchmarks compares go-stl structures against Go stdlib equivalents.
package benchmarks_test

import (
	"container/heap"
	"container/list"
	"testing"

	"github.com/sachinraghuwanshiofficial-blip/sac-stl/hashmap"
	"github.com/sachinraghuwanshiofficial-blip/sac-stl/hashset"
	"github.com/sachinraghuwanshiofficial-blip/sac-stl/priorityqueue"
	"github.com/sachinraghuwanshiofficial-blip/sac-stl/queue"
	"github.com/sachinraghuwanshiofficial-blip/sac-stl/stack"
	"github.com/sachinraghuwanshiofficial-blip/sac-stl/treemap"
	"github.com/sachinraghuwanshiofficial-blip/sac-stl/vector"
)

const N = 100_000

// ── Vector vs []int ──────────────────────────────────────────────────────────

func BenchmarkVector_PushBack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := vector.NewWithCapacity[int](N)
		for j := 0; j < N; j++ {
			v.PushBack(j)
		}
	}
}

func BenchmarkSlice_Append(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]int, 0, N)
		for j := 0; j < N; j++ {
			s = append(s, j)
		}
	}
}

// ── Stack vs container/list ──────────────────────────────────────────────────

func BenchmarkStack_PushPop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := stack.NewWithCapacity[int](N)
		for j := 0; j < N; j++ {
			s.Push(j)
		}
		for j := 0; j < N; j++ {
			s.Pop()
		}
	}
}

func BenchmarkList_PushBackPopBack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := list.New()
		for j := 0; j < N; j++ {
			l.PushBack(j)
		}
		for j := 0; j < N; j++ {
			l.Remove(l.Back())
		}
	}
}

// ── Queue vs container/list ──────────────────────────────────────────────────

func BenchmarkQueue_PushPop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		q := queue.NewWithCapacity[int](N)
		for j := 0; j < N; j++ {
			q.Push(j)
		}
		for j := 0; j < N; j++ {
			q.Pop()
		}
	}
}

func BenchmarkList_PushBackPopFront(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := list.New()
		for j := 0; j < N; j++ {
			l.PushBack(j)
		}
		for j := 0; j < N; j++ {
			l.Remove(l.Front())
		}
	}
}

// ── PriorityQueue vs container/heap ─────────────────────────────────────────

// stdlib heap adapter
type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *intHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *intHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func BenchmarkPQ_PushPop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pq := priorityqueue.NewWithCapacity(N, func(a, b int) bool { return a < b })
		for j := 0; j < N; j++ {
			pq.Push(j)
		}
		for j := 0; j < N; j++ {
			pq.Pop()
		}
	}
}

func BenchmarkHeap_PushPop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		h := &intHeap{}
		heap.Init(h)
		for j := 0; j < N; j++ {
			heap.Push(h, j)
		}
		for j := 0; j < N; j++ {
			heap.Pop(h)
		}
	}
}

// ── HashMap vs builtin map ───────────────────────────────────────────────────

func BenchmarkHashMap_Put(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := hashmap.NewWithCapacity[int, int](N)
		for j := 0; j < N; j++ {
			m.Put(j, j)
		}
	}
}

func BenchmarkBuiltinMap_Put(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := make(map[int]int, N)
		for j := 0; j < N; j++ {
			m[j] = j
		}
	}
}

func BenchmarkHashMap_Get(b *testing.B) {
	m := hashmap.NewWithCapacity[int, int](N)
	for j := 0; j < N; j++ {
		m.Put(j, j)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Get(i % N)
	}
}

func BenchmarkBuiltinMap_Get(b *testing.B) {
	m := make(map[int]int, N)
	for j := 0; j < N; j++ {
		m[j] = j
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m[i%N]
	}
}

// ── HashSet vs map[T]struct{} ────────────────────────────────────────────────

func BenchmarkHashSet_Insert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := hashset.NewWithCapacity[int](N)
		for j := 0; j < N; j++ {
			s.Insert(j)
		}
	}
}

func BenchmarkBuiltinSet_Insert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := make(map[int]struct{}, N)
		for j := 0; j < N; j++ {
			m[j] = struct{}{}
		}
	}
}

// ── TreeMap (ordered) ────────────────────────────────────────────────────────

func BenchmarkTreeMap_Put(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := treemap.New[int, int](func(a, b int) bool { return a < b })
		for j := 0; j < N; j++ {
			m.Put(j, j)
		}
	}
}

func BenchmarkTreeMap_Get(b *testing.B) {
	m := treemap.New[int, int](func(a, b int) bool { return a < b })
	for j := 0; j < N; j++ {
		m.Put(j, j)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Get(i % N)
	}
}
