package stack_test

import (
	"testing"

	"github.com/sachinraghuwanshiofficial-blip/sac-stl/stack"
)

func TestStack(t *testing.T) {
	s := stack.New[int]()
	for i := 1; i <= 5; i++ {
		s.Push(i)
	}
	if s.Len() != 5 {
		t.Fatalf("want 5, got %d", s.Len())
	}
	for i := 5; i >= 1; i-- {
		if top, ok := s.Top(); !ok || top != i {
			t.Fatalf("Top: want (%d, true), got (%d, %v)", i, top, ok)
		}
		val, ok := s.Pop()
		if !ok || val != i {
			t.Fatalf("Pop: want (%d, true), got (%d, %v)", i, val, ok)
		}
	}
	if _, ok := s.Pop(); ok {
		t.Fatal("Pop on empty should return false")
	}
}

func BenchmarkStackPush(b *testing.B) {
	s := stack.NewWithCapacity[int](b.N)
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
}
