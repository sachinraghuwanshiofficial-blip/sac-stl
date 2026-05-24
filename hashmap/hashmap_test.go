package hashmap_test

import (
	"testing"

	"github.com/sachinraghuwanshiofficial-blip/sac-stl/hashmap"
)

func TestHashMap(t *testing.T) {
	m := hashmap.New[string, int]()
	m.Put("a", 1)
	m.Put("b", 2)
	m.Put("c", 3)

	if v, ok := m.Get("b"); !ok || v != 2 {
		t.Fatalf("want (2, true), got (%d, %v)", v, ok)
	}
	if m.GetOrDefault("z", 99) != 99 {
		t.Fatal("missing key should return default")
	}

	m.Delete("b")
	if m.Contains("b") {
		t.Fatal("b should be deleted")
	}
	if m.Len() != 2 {
		t.Fatalf("want 2, got %d", m.Len())
	}
}

func BenchmarkHashMapPut(b *testing.B) {
	m := hashmap.NewWithCapacity[int, int](b.N)
	for i := 0; i < b.N; i++ {
		m.Put(i, i)
	}
}

func BenchmarkBuiltinMapPut(b *testing.B) {
	m := make(map[int]int, b.N)
	for i := 0; i < b.N; i++ {
		m[i] = i
	}
}
