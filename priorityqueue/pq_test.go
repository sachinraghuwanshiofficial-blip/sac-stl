package priorityqueue_test

import (
	"testing"

	"github.com/sachin/go-stl/priorityqueue"
)

func TestMinHeap(t *testing.T) {
	pq := priorityqueue.New(func(a, b int) bool { return a < b })
	for _, v := range []int{5, 1, 3, 2, 4} {
		pq.Push(v)
	}
	for i := 1; i <= 5; i++ {
		top, ok := pq.Top()
		if !ok || top != i {
			t.Fatalf("Top: want (%d, true), got (%d, %v)", i, top, ok)
		}
		val, ok := pq.Pop()
		if !ok || val != i {
			t.Fatalf("Pop: want (%d, true), got (%d, %v)", i, val, ok)
		}
	}
}

func TestMaxHeap(t *testing.T) {
	pq := priorityqueue.New(func(a, b int) bool { return a > b })
	for _, v := range []int{5, 1, 3, 2, 4} {
		pq.Push(v)
	}
	for i := 5; i >= 1; i-- {
		val, ok := pq.Pop()
		if !ok || val != i {
			t.Fatalf("Pop: want (%d, true), got (%d, %v)", i, val, ok)
		}
	}
}

func TestNewFrom(t *testing.T) {
	pq := priorityqueue.NewFrom([]int{9, 3, 7, 1, 5}, func(a, b int) bool { return a < b })
	if top, ok := pq.Top(); !ok || top != 1 {
		t.Fatalf("want top=1, got %d", top)
	}
}

func TestCustomStruct(t *testing.T) {
	type Task struct {
		priority int
		name     string
	}
	pq := priorityqueue.New(func(a, b Task) bool { return a.priority < b.priority })
	pq.Push(Task{3, "low"})
	pq.Push(Task{1, "high"})
	pq.Push(Task{2, "medium"})

	top, _ := pq.Pop()
	if top.name != "high" {
		t.Fatalf("want 'high', got %q", top.name)
	}
}

func BenchmarkPQPush(b *testing.B) {
	pq := priorityqueue.NewWithCapacity(b.N, func(a, b int) bool { return a < b })
	for i := 0; i < b.N; i++ {
		pq.Push(i)
	}
}

func BenchmarkPQPop(b *testing.B) {
	pq := priorityqueue.NewWithCapacity(b.N, func(a, b int) bool { return a < b })
	for i := 0; i < b.N; i++ {
		pq.Push(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pq.Pop()
	}
}
