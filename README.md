# go-stl

[![Go Version](https://img.shields.io/badge/go-%3E%3D1.21-blue?logo=go)](https://go.dev)
[![Go Reference](https://pkg.go.dev/badge/github.com/sachin/go-stl.svg)](https://pkg.go.dev/github.com/sachin/go-stl)
[![License: MIT](https://img.shields.io/badge/license-MIT-green)](LICENSE)
[![Tests](https://img.shields.io/badge/tests-passing-brightgreen)](#)

> Generic, performant C++ STL–inspired data structures for Go.

**go-stl** brings the ergonomics of `std::vector`, `std::stack`, `std::queue`, `std::priority_queue`, `std::map`, `std::unordered_map`, `std::set`, and `std::unordered_set` to Go — with full generics, C++-style free functions, and zero external dependencies.

---

## Feature Highlights

| Feature | Details |
|---|---|
| **Fully generic** | All containers use Go generics (`[T any]`, `[K comparable, V any]`) |
| **C++ naming** | `PushBack`, `PopBack`, `Top`, `Front`, `Size`, `Empty`, `Insert`, `Erase`, `Find`, `Count` |
| **Free functions** | Unified `stl` package with `stl.Push`, `stl.PopBack`, `stl.Top`, … |
| **Ordered containers** | `treemap` / `treeset` backed by a Left-Leaning Red-Black Tree |
| **No dependencies** | Pure Go — no third-party packages |
| **Ring-buffer queue** | True O(1) dequeue — no element shifting |
| **O(n) heap build** | `NewFrom` heapifies in linear time |
| **Set algebra** | `Union`, `Intersection`, `Difference` on `hashset` |

---

## Installation

```bash
go get github.com/sachin/go-stl
```

Requires **Go 1.21+**.

---

## Quick Start

### Using the unified `stl` package (C++ style)

```go
import "github.com/sachin/go-stl/stl"

// ── Vector ───────────────────────────────────────────────────────────────────
v := stl.MakeVector[int]()
stl.PushBack(v, 1)
stl.PushBack(v, 2)
stl.PushBack(v, 3)
val, ok := stl.PopBack(v) // val=3, ok=true
front, _  := stl.Front(v) // front=1

// ── Stack ────────────────────────────────────────────────────────────────────
s := stl.MakeStack[string]()
stl.Push(s, "hello")
stl.Push(s, "world")
top, ok := stl.Top(s)   // top="world", ok=true
stl.Pop(s)

// ── Queue ────────────────────────────────────────────────────────────────────
q := stl.MakeQueue[int]()
stl.Enqueue(q, 10)
stl.Enqueue(q, 20)
front, _ := stl.QueueFront(q) // front=10
stl.Dequeue(q)

// ── Min-Heap ─────────────────────────────────────────────────────────────────
pq := stl.MakeMinHeap(func(a, b int) bool { return a < b })
stl.PushPQ(pq, 5)
stl.PushPQ(pq, 1)
stl.PushPQ(pq, 3)
best, _ := stl.TopPQ(pq) // best=1

// ── Ordered Map (like std::map) ───────────────────────────────────────────────
m := stl.MakeTreeMap[int, string](func(a, b int) bool { return a < b })
stl.MapInsert(m, 42, "answer")
v2, ok := stl.MapFind(m, 42) // v2="answer", ok=true
stl.MapErase(m, 42)

// ── Hash Map (like std::unordered_map) ───────────────────────────────────────
hm := stl.MakeHashMap[string, int]()
stl.HMapInsert(hm, "score", 100)
score, _ := stl.HMapFind(hm, "score") // score=100

// ── Ordered Set (like std::set) ──────────────────────────────────────────────
ts := stl.MakeTreeSet(func(a, b int) bool { return a < b })
stl.SetInsert(ts, 10)
stl.SetInsert(ts, 5)
stl.SetInsert(ts, 20)
// iterates: 5, 10, 20

// ── Hash Set (like std::unordered_set) ───────────────────────────────────────
hs := stl.MakeHashSet[string]()
stl.HSetInsert(hs, "go")
stl.HSetInsert(hs, "stl")
fmt.Println(stl.HSetContains(hs, "go")) // true
```

### Using packages directly (Go method style)

```go
import (
    "github.com/sachin/go-stl/vector"
    "github.com/sachin/go-stl/treemap"
    "github.com/sachin/go-stl/priorityqueue"
)

// Vector
v := vector.NewFrom([]int{3, 1, 4, 1, 5})
v.PushBack(9)
v.Insert(0, 0)        // insert at front
v.Erase(1)            // erase index 1
v.Range(func(i int, val int) bool {
    fmt.Println(i, val)
    return true
})

// Ordered map with floor/ceiling (no stdlib equivalent)
m := treemap.New[int, string](func(a, b int) bool { return a < b })
m.Put(10, "ten")
m.Put(20, "twenty")
m.Put(30, "thirty")

k, v2, _ := m.Floor(25)   // k=20, v2="twenty"
k, v2, _ = m.Ceiling(15)  // k=20, v2="twenty"
m.Range(func(k int, v string) bool {
    fmt.Printf("%d → %s\n", k, v)
    return true
})

// Priority queue with custom struct
type Job struct { priority int; name string }
pq := priorityqueue.New(func(a, b Job) bool { return a.priority < b.priority })
pq.Push(Job{3, "low"})
pq.Push(Job{1, "critical"})
pq.Push(Job{2, "normal"})
for !pq.Empty() {
    job, _ := pq.Pop()
    fmt.Println(job.name) // critical, normal, low
}
```

---

## Packages

### `vector` — Dynamic Array (`std::vector<T>`)

```go
v := vector.New[int]()
v := vector.NewWithCapacity[int](128)
v := vector.NewFrom([]int{1, 2, 3})
```

| Method | C++ equivalent | Complexity |
|---|---|---|
| `PushBack(val)` | `push_back` | O(1) amortized |
| `PopBack()` | `pop_back` | O(1) |
| `Front()` | `front()` | O(1) |
| `Back()` | `back()` | O(1) |
| `At(i)` | `at(i)` | O(1) |
| `Set(i, val)` | `v[i] = val` | O(1) |
| `Insert(i, val)` | `insert` | O(n) |
| `Erase(i)` | `erase` | O(n) |
| `Len()` / `Size()` | `size()` | O(1) |
| `Cap()` | `capacity()` | O(1) |
| `Empty()` | `empty()` | O(1) |
| `Reserve(n)` | `reserve` | O(n) |
| `Resize(n)` | `resize` | O(n) |
| `Clear()` | `clear` | O(n) |
| `Slice()` | `data()` (copy) | O(n) |
| `Range(fn)` | range-for | O(n) |

---

### `stack` — LIFO Stack (`std::stack<T>`)

```go
s := stack.New[int]()
s := stack.NewWithCapacity[int](64)
```

| Method | C++ equivalent | Complexity |
|---|---|---|
| `Push(val)` | `push` | O(1) amortized |
| `Pop()` | `pop` + `top` | O(1) |
| `Top()` | `top` | O(1) |
| `Len()` / `Size()` | `size()` | O(1) |
| `Empty()` / `IsEmpty()` | `empty()` | O(1) |
| `Clear()` | — | O(n) |

---

### `queue` — FIFO Queue (`std::queue<T>`)

Backed by a **power-of-two ring buffer** — O(1) push *and* pop with no element shifting.

```go
q := queue.New[int]()
q := queue.NewWithCapacity[int](256)
```

| Method | C++ equivalent | Complexity |
|---|---|---|
| `Push(val)` | `push` | O(1) amortized |
| `Pop()` | `pop` + `front` | O(1) |
| `Front()` | `front()` | O(1) |
| `Back()` | `back()` | O(1) |
| `Len()` / `Size()` | `size()` | O(1) |
| `Empty()` / `IsEmpty()` | `empty()` | O(1) |
| `Clear()` | — | O(n) |

---

### `priorityqueue` — Binary Heap (`std::priority_queue<T>`)

Ordering is controlled by a **user-supplied comparator** — making min-heap, max-heap, or any custom ordering trivial.

```go
// Min-heap
pq := priorityqueue.New(func(a, b int) bool { return a < b })

// Max-heap
pq := priorityqueue.New(func(a, b int) bool { return a > b })

// From existing slice — O(n) heapify
pq := priorityqueue.NewFrom([]int{5,3,1,4,2}, func(a, b int) bool { return a < b })
```

| Method | C++ equivalent | Complexity |
|---|---|---|
| `Push(val)` | `push` | O(log n) |
| `Pop()` | `pop` + `top` | O(log n) |
| `Top()` | `top()` | O(1) |
| `Len()` / `Size()` | `size()` | O(1) |
| `Empty()` / `IsEmpty()` | `empty()` | O(1) |
| `Clear()` | — | O(n) |

---

### `treemap` — Ordered Map (`std::map<K,V>`)

Backed by a **Left-Leaning Red-Black Tree** (Sedgewick 2008). Adds `Floor`/`Ceiling` not present in C++ std::map.

```go
m := treemap.New[int, string](func(a, b int) bool { return a < b })
```

| Method | C++ equivalent | Complexity |
|---|---|---|
| `Put(k, v)` / `Insert(k, v)` | `insert` / `operator[]` | O(log n) |
| `Get(k)` / `Find(k)` | `find` | O(log n) |
| `Delete(k)` / `Erase(k)` | `erase` | O(log n) |
| `Contains(k)` | `count(k)` | O(log n) |
| `Min()` | `begin()` | O(log n) |
| `Max()` | `rbegin()` | O(log n) |
| `Floor(k)` | `lower_bound` (adjusted) | O(log n) |
| `Ceiling(k)` | `lower_bound` | O(log n) |
| `Len()` / `Size()` | `size()` | O(1) |
| `Empty()` / `IsEmpty()` | `empty()` | O(1) |
| `Range(fn)` | range-for (ordered) | O(n) |
| `RangeFrom(k, fn)` | `lower_bound` + loop | O(log n + k) |
| `Keys()` | — | O(n) |
| `Values()` | — | O(n) |

---

### `hashmap` — Unordered Map (`std::unordered_map<K,V>`)

```go
m := hashmap.New[string, int]()
m := hashmap.NewWithCapacity[string, int](1024)
```

| Method | C++ equivalent | Complexity |
|---|---|---|
| `Put(k, v)` / `Insert(k, v)` | `insert` / `operator[]` | O(1) avg |
| `Get(k)` / `Find(k)` | `find` | O(1) avg |
| `GetOrDefault(k, d)` | `find` + fallback | O(1) avg |
| `Delete(k)` / `Erase(k)` | `erase` | O(1) avg |
| `Contains(k)` | `count(k)` | O(1) avg |
| `Len()` / `Size()` | `size()` | O(1) |
| `Empty()` / `IsEmpty()` | `empty()` | O(1) |
| `Range(fn)` | range-for | O(n) |
| `Keys()` / `Values()` | — | O(n) |

---

### `treeset` — Ordered Set (`std::set<T>`)

Built on `treemap`. Elements are **unique and sorted**.

```go
s := treeset.New(func(a, b int) bool { return a < b })
```

| Method | C++ equivalent | Complexity |
|---|---|---|
| `Insert(val)` | `insert` | O(log n) |
| `Delete(val)` / `Erase(val)` | `erase` | O(log n) |
| `Contains(val)` | `count(val)` | O(log n) |
| `Count(val)` | `count(val)` → 0 or 1 | O(log n) |
| `Min()` / `Max()` | `begin()` / `rbegin()` | O(log n) |
| `Floor(val)` / `Ceiling(val)` | `lower_bound` | O(log n) |
| `Len()` / `Size()` | `size()` | O(1) |
| `Empty()` / `IsEmpty()` | `empty()` | O(1) |
| `Range(fn)` | range-for (ordered) | O(n) |
| `Slice()` | — | O(n) |

---

### `hashset` — Unordered Set (`std::unordered_set<T>`)

```go
s := hashset.New[string]()
s := hashset.NewFrom([]string{"a", "b", "c"})
```

| Method | C++ equivalent | Complexity |
|---|---|---|
| `Insert(val)` | `insert` | O(1) avg |
| `Delete(val)` / `Erase(val)` | `erase` | O(1) avg |
| `Contains(val)` | `count(val)` | O(1) avg |
| `Count(val)` | `count(val)` → 0 or 1 | O(1) avg |
| `Len()` / `Size()` | `size()` | O(1) |
| `Empty()` / `IsEmpty()` | `empty()` | O(1) |
| `Union(other)` | — | O(n+m) |
| `Intersection(other)` | — | O(min(n,m)) |
| `Difference(other)` | — | O(n) |
| `Range(fn)` | range-for | O(n) |
| `Slice()` | — | O(n) |

---

## C++ → Go Cheatsheet

| C++ | go-stl (method style) | go-stl (free function style) |
|---|---|---|
| `std::vector<int> v` | `v := vector.New[int]()` | `v := stl.MakeVector[int]()` |
| `v.push_back(x)` | `v.PushBack(x)` | `stl.PushBack(v, x)` |
| `v.pop_back()` | `v.PopBack()` | `stl.PopBack(v)` |
| `v.front()` | `v.Front()` | `stl.Front(v)` |
| `v.back()` | `v.Back()` | `stl.Back(v)` |
| `v.size()` | `v.Size()` | `stl.Size(v)` |
| `v.empty()` | `v.Empty()` | `stl.Empty(v)` |
| `v.at(i)` | `v.At(i)` | `stl.At(v, i)` |
| `std::stack<int> s` | `s := stack.New[int]()` | `s := stl.MakeStack[int]()` |
| `s.push(x)` | `s.Push(x)` | `stl.Push(s, x)` |
| `s.pop()` | `s.Pop()` | `stl.Pop(s)` |
| `s.top()` | `s.Top()` | `stl.Top(s)` |
| `std::queue<int> q` | `q := queue.New[int]()` | `q := stl.MakeQueue[int]()` |
| `q.push(x)` | `q.Push(x)` | `stl.Enqueue(q, x)` |
| `q.pop()` | `q.Pop()` | `stl.Dequeue(q)` |
| `q.front()` | `q.Front()` | `stl.QueueFront(q)` |
| `std::priority_queue<int> pq` | `pq := priorityqueue.New(less)` | `pq := stl.MakeMinHeap(less)` |
| `pq.push(x)` | `pq.Push(x)` | `stl.PushPQ(pq, x)` |
| `pq.pop()` | `pq.Pop()` | `stl.PopPQ(pq)` |
| `pq.top()` | `pq.Top()` | `stl.TopPQ(pq)` |
| `std::map<int,string> m` | `m := treemap.New[int,string](less)` | `m := stl.MakeTreeMap[int,string](less)` |
| `m[k] = v` | `m.Put(k, v)` / `m.Insert(k, v)` | `stl.MapInsert(m, k, v)` |
| `m.find(k)` | `m.Find(k)` / `m.Get(k)` | `stl.MapFind(m, k)` |
| `m.erase(k)` | `m.Erase(k)` / `m.Delete(k)` | `stl.MapErase(m, k)` |
| `std::unordered_map<K,V> m` | `m := hashmap.New[K,V]()` | `m := stl.MakeHashMap[K,V]()` |
| `std::set<int> s` | `s := treeset.New(less)` | `s := stl.MakeTreeSet(less)` |
| `s.insert(x)` | `s.Insert(x)` | `stl.SetInsert(s, x)` |
| `s.erase(x)` | `s.Erase(x)` | `stl.SetErase(s, x)` |
| `s.count(x)` | `s.Count(x)` / `s.Contains(x)` | `stl.SetContains(s, x)` |
| `std::unordered_set<T> s` | `s := hashset.New[T]()` | `s := stl.MakeHashSet[T]()` |

---

## Benchmarks

Measured on **Apple M5**, Go 1.26.3, N = 100,000 operations.

```
BenchmarkVector_PushBack      192,808 ns/op    802,818 B/op    1 allocs/op
BenchmarkSlice_Append          56,311 ns/op    802,826 B/op    1 allocs/op

BenchmarkStack_PushPop        342,394 ns/op    802,816 B/op    1 allocs/op
BenchmarkList_PushBackPopBack 3,566,095 ns/op  5,598,028 B/op  199,745 allocs/op   ← 10× slower

BenchmarkQueue_PushPop        304,916 ns/op   1,048,577 B/op   1 allocs/op
BenchmarkList_PushBackPopFront 3,460,356 ns/op 5,598,024 B/op  199,745 allocs/op   ← 11× slower

BenchmarkPQ_PushPop          5,953,455 ns/op    802,816 B/op    1 allocs/op
BenchmarkHeap_PushPop        9,243,797 ns/op  5,697,350 B/op  199,517 allocs/op    ← 1.6× slower

BenchmarkHashMap_Put           907,756 ns/op  2,364,557 B/op    257 allocs/op
BenchmarkBuiltinMap_Put        970,888 ns/op  2,364,557 B/op    257 allocs/op      ← equal

BenchmarkHashMap_Get               4.614 ns/op        0 B/op    0 allocs/op
BenchmarkBuiltinMap_Get            4.796 ns/op        0 B/op    0 allocs/op        ← equal

BenchmarkHashSet_Insert        940,487 ns/op  2,364,558 B/op    257 allocs/op
BenchmarkBuiltinSet_Insert     994,901 ns/op  2,364,558 B/op    257 allocs/op      ← equal

BenchmarkTreeMap_Get            37.06 ns/op         0 B/op    0 allocs/op
```

**Key takeaways:**
- `stack` and `queue` are **10–11× faster** than `container/list` with **200,000× fewer allocations**
- `priorityqueue` is **1.6× faster** than `container/heap` (no interface boxing)
- `hashmap` and `hashset` match the raw built-in map — zero overhead wrapper
- `treemap` fills a gap: there is no ordered map in the Go standard library

---

## Package Structure

```
go-stl/
├── stl/            unified entry point — type aliases + free functions
├── vector/         dynamic array        (std::vector)
├── stack/          LIFO stack           (std::stack)
├── queue/          ring-buffer FIFO     (std::queue)
├── priorityqueue/  binary heap          (std::priority_queue)
├── treemap/        LLRB ordered map     (std::map)
├── hashmap/        hash map wrapper     (std::unordered_map)
├── treeset/        ordered set          (std::set)
├── hashset/        hash set             (std::unordered_set)
└── benchmarks/     comparative benchmarks
```

---

## Design Principles

1. **Generics-first** — every container is fully type-safe with no `any` casting at the call site.
2. **No panics on empty** — all read operations return `(value, ok bool)` pairs.
3. **Comparator at construction** — pass `less` once; no need to implement interfaces.
4. **C++ naming parity** — every C++ method name has a direct Go alias (`Size`, `Empty`, `Insert`, `Erase`, `Find`, `Count`).
5. **Zero dependencies** — the LLRB tree is implemented from scratch; no third-party packages required.
6. **Idiomatic iteration** — `Range(func(k, v) bool)` — stop early by returning `false`.

---

## License

MIT © Sachin