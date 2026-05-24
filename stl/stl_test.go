package stl_test

import (
	"testing"

	"github.com/sachinraghuwanshiofficial-blip/sac-stl/stl"
)

func TestVectorFreeFunctions(t *testing.T) {
	v := stl.MakeVector[int]()
	stl.PushBack(v, 10)
	stl.PushBack(v, 20)
	stl.PushBack(v, 30)

	if stl.Size(v) != 3 {
		t.Fatalf("want 3, got %d", stl.Size(v))
	}
	if b, ok := stl.Back(v); !ok || b != 30 {
		t.Fatalf("want back=30, got %d", b)
	}
	val, ok := stl.PopBack(v)
	if !ok || val != 30 {
		t.Fatalf("want (30, true), got (%d, %v)", val, ok)
	}
	if stl.Empty(v) {
		t.Fatal("should not be empty")
	}
}

func TestStackFreeFunctions(t *testing.T) {
	s := stl.MakeStack[string]()
	stl.Push(s, "a")
	stl.Push(s, "b")
	stl.Push(s, "c")

	if top, ok := stl.Top(s); !ok || top != "c" {
		t.Fatalf("want top=c, got %q", top)
	}
	stl.Pop(s)
	if top, ok := stl.Top(s); !ok || top != "b" {
		t.Fatalf("want top=b after pop, got %q", top)
	}
	if stl.IsEmpty(s) {
		t.Fatal("should not be empty")
	}
}

func TestQueueFreeFunctions(t *testing.T) {
	q := stl.MakeQueue[int]()
	stl.Enqueue(q, 1)
	stl.Enqueue(q, 2)
	stl.Enqueue(q, 3)

	for i := 1; i <= 3; i++ {
		front, ok := stl.QueueFront(q)
		if !ok || front != i {
			t.Fatalf("want front=%d, got %d", i, front)
		}
		stl.Dequeue(q)
	}
}

func TestMinHeapFreeFunctions(t *testing.T) {
	pq := stl.MakeMinHeap(func(a, b int) bool { return a < b })
	stl.PushPQ(pq, 5)
	stl.PushPQ(pq, 1)
	stl.PushPQ(pq, 3)

	top, ok := stl.TopPQ(pq)
	if !ok || top != 1 {
		t.Fatalf("want top=1, got %d", top)
	}
	stl.PopPQ(pq)
	top, _ = stl.TopPQ(pq)
	if top != 3 {
		t.Fatalf("want top=3, got %d", top)
	}
}

func TestMaxHeapFreeFunctions(t *testing.T) {
	pq := stl.MakeMaxHeap(func(a, b int) bool { return a > b })
	stl.PushPQ(pq, 5)
	stl.PushPQ(pq, 1)
	stl.PushPQ(pq, 3)

	top, _ := stl.TopPQ(pq)
	if top != 5 {
		t.Fatalf("want top=5, got %d", top)
	}
}

func TestTreeMapFreeFunctions(t *testing.T) {
	m := stl.MakeTreeMap[int, string](func(a, b int) bool { return a < b })
	stl.MapInsert(m, 1, "one")
	stl.MapInsert(m, 2, "two")

	v, ok := stl.MapFind(m, 1)
	if !ok || v != "one" {
		t.Fatalf("want (one, true), got (%q, %v)", v, ok)
	}
	stl.MapErase(m, 1)
	if stl.MapSize(m) != 1 {
		t.Fatalf("want size=1, got %d", stl.MapSize(m))
	}
}

func TestHashMapFreeFunctions(t *testing.T) {
	m := stl.MakeHashMap[string, int]()
	stl.HMapInsert(m, "x", 42)
	v, ok := stl.HMapFind(m, "x")
	if !ok || v != 42 {
		t.Fatalf("want (42, true), got (%d, %v)", v, ok)
	}
	stl.HMapErase(m, "x")
	if _, ok := stl.HMapFind(m, "x"); ok {
		t.Fatal("x should be deleted")
	}
}

func TestSetFreeFunctions(t *testing.T) {
	s := stl.MakeTreeSet(func(a, b int) bool { return a < b })
	stl.SetInsert(s, 10)
	stl.SetInsert(s, 20)
	if !stl.SetContains(s, 10) {
		t.Fatal("should contain 10")
	}
	stl.SetErase(s, 10)
	if stl.SetContains(s, 10) {
		t.Fatal("10 should be erased")
	}
}

func TestHashSetFreeFunctions(t *testing.T) {
	s := stl.MakeHashSet[int]()
	stl.HSetInsert(s, 5)
	stl.HSetInsert(s, 10)
	if !stl.HSetContains(s, 5) {
		t.Fatal("should contain 5")
	}
	stl.HSetErase(s, 5)
	if stl.HSetContains(s, 5) {
		t.Fatal("5 should be erased")
	}
}

// MakeVectorFrom and MakeVectorN
func TestMakeVectorHelpers(t *testing.T) {
	v := stl.MakeVectorFrom([]int{1, 2, 3})
	if stl.Size(v) != 3 || stl.At(v, 0) != 1 {
		t.Fatal("MakeVectorFrom failed")
	}
	v2 := stl.MakeVectorN[int](5)
	if stl.Size(v2) != 5 {
		t.Fatalf("MakeVectorN: want 5, got %d", stl.Size(v2))
	}
}
