package hashset_test

import (
	"testing"

	"github.com/sachin/go-stl/hashset"
)

func TestHashSet(t *testing.T) {
	s := hashset.NewFrom([]int{1, 2, 3, 4, 5})
	if s.Len() != 5 {
		t.Fatalf("want 5, got %d", s.Len())
	}
	s.Insert(3) // duplicate
	if s.Len() != 5 {
		t.Fatal("duplicate should not increase len")
	}
	s.Delete(3)
	if s.Contains(3) {
		t.Fatal("3 should be deleted")
	}
}

func TestSetOps(t *testing.T) {
	a := hashset.NewFrom([]int{1, 2, 3, 4})
	b := hashset.NewFrom([]int{3, 4, 5, 6})

	u := a.Union(b)
	if u.Len() != 6 {
		t.Fatalf("union: want 6, got %d", u.Len())
	}

	i := a.Intersection(b)
	if i.Len() != 2 || !i.Contains(3) || !i.Contains(4) {
		t.Fatalf("intersection should be {3,4}, got %v", i.Slice())
	}

	d := a.Difference(b)
	if d.Len() != 2 || !d.Contains(1) || !d.Contains(2) {
		t.Fatalf("difference should be {1,2}, got %v", d.Slice())
	}
}

func BenchmarkHashSetInsert(b *testing.B) {
	s := hashset.NewWithCapacity[int](b.N)
	for i := 0; i < b.N; i++ {
		s.Insert(i)
	}
}
