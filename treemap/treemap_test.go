package treemap_test

import (
	"testing"

	"github.com/sachin/go-stl/treemap"
)

func intLess(a, b int) bool { return a < b }

func TestPutGet(t *testing.T) {
	m := treemap.New[int, string](intLess)
	m.Put(3, "three")
	m.Put(1, "one")
	m.Put(2, "two")

	if v, ok := m.Get(2); !ok || v != "two" {
		t.Fatalf("want (two, true), got (%q, %v)", v, ok)
	}
	if _, ok := m.Get(99); ok {
		t.Fatal("should not find 99")
	}
}

func TestDelete(t *testing.T) {
	m := treemap.New[int, string](intLess)
	for i := 1; i <= 10; i++ {
		m.Put(i, "x")
	}
	m.Delete(5)
	if m.Contains(5) {
		t.Fatal("5 should be deleted")
	}
	if m.Len() != 9 {
		t.Fatalf("want 9, got %d", m.Len())
	}
}

func TestInOrder(t *testing.T) {
	m := treemap.New[int, int](intLess)
	for _, v := range []int{5, 3, 7, 1, 9, 2, 8, 4, 6} {
		m.Put(v, v)
	}
	prev := -1
	m.Range(func(k, _ int) bool {
		if k <= prev {
			t.Fatalf("out of order: %d after %d", k, prev)
		}
		prev = k
		return true
	})
}

func TestMinMax(t *testing.T) {
	m := treemap.New[int, int](intLess)
	for i := 1; i <= 5; i++ {
		m.Put(i, i)
	}
	if k, _, _ := m.Min(); k != 1 {
		t.Fatalf("want min=1, got %d", k)
	}
	if k, _, _ := m.Max(); k != 5 {
		t.Fatalf("want max=5, got %d", k)
	}
}

func TestFloorCeiling(t *testing.T) {
	m := treemap.New[int, int](intLess)
	for _, v := range []int{1, 3, 5, 7, 9} {
		m.Put(v, v)
	}
	if k, _, ok := m.Floor(6); !ok || k != 5 {
		t.Fatalf("floor(6) want 5, got %d", k)
	}
	if k, _, ok := m.Ceiling(4); !ok || k != 5 {
		t.Fatalf("ceiling(4) want 5, got %d", k)
	}
}

func TestUpdate(t *testing.T) {
	m := treemap.New[int, int](intLess)
	m.Put(1, 100)
	m.Put(1, 200)
	if v, _ := m.Get(1); v != 200 {
		t.Fatalf("want 200, got %d", v)
	}
	if m.Len() != 1 {
		t.Fatalf("want len 1, got %d", m.Len())
	}
}

func BenchmarkTreeMapPut(b *testing.B) {
	m := treemap.New[int, int](intLess)
	for i := 0; i < b.N; i++ {
		m.Put(i, i)
	}
}

func BenchmarkTreeMapGet(b *testing.B) {
	m := treemap.New[int, int](intLess)
	for i := 0; i < b.N; i++ {
		m.Put(i, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Get(i)
	}
}
