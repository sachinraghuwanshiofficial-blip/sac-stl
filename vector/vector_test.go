package vector_test

import (
	"testing"

	"github.com/sachin/go-stl/vector"
)

func TestPushPopBack(t *testing.T) {
	v := vector.New[int]()
	for i := 0; i < 10; i++ {
		v.PushBack(i)
	}
	if v.Len() != 10 {
		t.Fatalf("want len 10, got %d", v.Len())
	}
	for i := 9; i >= 0; i-- {
		val, ok := v.PopBack()
		if !ok || val != i {
			t.Fatalf("want (%d, true), got (%d, %v)", i, val, ok)
		}
	}
	if _, ok := v.PopBack(); ok {
		t.Fatal("PopBack on empty should return false")
	}
}

func TestInsertErase(t *testing.T) {
	v := vector.NewFrom([]int{1, 2, 3, 4, 5})
	v.Insert(2, 99)
	if v.At(2) != 99 {
		t.Fatalf("want 99 at index 2, got %d", v.At(2))
	}
	if v.Len() != 6 {
		t.Fatalf("want len 6, got %d", v.Len())
	}
	v.Erase(2)
	if v.At(2) != 3 {
		t.Fatalf("want 3 at index 2, got %d", v.At(2))
	}
	if v.Len() != 5 {
		t.Fatalf("want len 5, got %d", v.Len())
	}
}

func TestReserveResize(t *testing.T) {
	v := vector.NewWithCapacity[int](100)
	if v.Cap() < 100 {
		t.Fatalf("want cap >= 100, got %d", v.Cap())
	}
	v.Resize(50)
	if v.Len() != 50 {
		t.Fatalf("want len 50, got %d", v.Len())
	}
}

func TestRange(t *testing.T) {
	v := vector.NewFrom([]int{10, 20, 30})
	sum := 0
	v.Range(func(_ int, val int) bool {
		sum += val
		return true
	})
	if sum != 60 {
		t.Fatalf("want 60, got %d", sum)
	}
}

func BenchmarkVectorPushBack(b *testing.B) {
	v := vector.NewWithCapacity[int](b.N)
	for i := 0; i < b.N; i++ {
		v.PushBack(i)
	}
}

func BenchmarkSliceAppend(b *testing.B) {
	s := make([]int, 0, b.N)
	for i := 0; i < b.N; i++ {
		s = append(s, i)
	}
}
