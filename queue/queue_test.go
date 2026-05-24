package queue_test

import (
	"testing"

	"github.com/sachinraghuwanshiofficial-blip/sac-stl/queue"
)

func TestQueue(t *testing.T) {
	q := queue.New[int]()
	for i := 1; i <= 5; i++ {
		q.Push(i)
	}
	if q.Len() != 5 {
		t.Fatalf("want 5, got %d", q.Len())
	}
	for i := 1; i <= 5; i++ {
		if front, ok := q.Front(); !ok || front != i {
			t.Fatalf("Front: want (%d, true), got (%d, %v)", i, front, ok)
		}
		val, ok := q.Pop()
		if !ok || val != i {
			t.Fatalf("Pop: want (%d, true), got (%d, %v)", i, val, ok)
		}
	}
	if _, ok := q.Pop(); ok {
		t.Fatal("Pop on empty should return false")
	}
}

func TestQueueWrapAround(t *testing.T) {
	// ensure ring buffer wrap-around works
	q := queue.New[int]()
	for i := 0; i < 100; i++ {
		q.Push(i)
		if i%3 == 0 {
			q.Pop()
		}
	}
	for !q.Empty() {
		q.Pop()
	}
	if q.Len() != 0 {
		t.Fatal("queue should be empty")
	}
}

func BenchmarkQueuePushPop(b *testing.B) {
	q := queue.NewWithCapacity[int](1024)
	for i := 0; i < b.N; i++ {
		q.Push(i)
		q.Pop()
	}
}
