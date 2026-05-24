package treeset_test

import (
	"testing"

	"github.com/sachin/go-stl/treeset"
)

func TestTreeSet(t *testing.T) {
	s := treeset.New(func(a, b int) bool { return a < b })
	for _, v := range []int{5, 3, 1, 4, 2} {
		s.Insert(v)
	}
	if s.Len() != 5 {
		t.Fatalf("want 5, got %d", s.Len())
	}

	// duplicates ignored
	s.Insert(3)
	if s.Len() != 5 {
		t.Fatalf("duplicate should not increase len")
	}

	// iterate in order
	got := s.Slice()
	for i, v := range got {
		if v != i+1 {
			t.Fatalf("want %d at index %d, got %d", i+1, i, v)
		}
	}

	s.Delete(3)
	if s.Contains(3) {
		t.Fatal("3 should be deleted")
	}
}

func TestFloorCeiling(t *testing.T) {
	s := treeset.New(func(a, b int) bool { return a < b })
	for _, v := range []int{10, 20, 30, 40} {
		s.Insert(v)
	}
	if f, ok := s.Floor(25); !ok || f != 20 {
		t.Fatalf("floor(25) want 20, got %d", f)
	}
	if c, ok := s.Ceiling(25); !ok || c != 30 {
		t.Fatalf("ceiling(25) want 30, got %d", c)
	}
}
