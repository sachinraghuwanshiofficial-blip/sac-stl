// Package treemap provides a generic ordered map equivalent to C++ std::map.
//
// Backed by a Left-Leaning Red-Black (LLRB) tree (Sedgewick 2008).
// Keys are ordered by a user-supplied less comparator.
//
// Time complexity:
//   Put/Get/Delete  O(log n)
//   Min/Max         O(log n)
//   Floor/Ceiling   O(log n)
//   Len             O(1)
//   Range (in-order) O(n)
package treemap

const red = true
const black = false

type node[K, V any] struct {
	key         K
	val         V
	left, right *node[K, V]
	color       bool // red = true
	size        int
}

// Map is a generic ordered map backed by an LLRB tree.
type Map[K, V any] struct {
	root *node[K, V]
	less func(a, b K) bool
}

// New returns an empty Map with the given comparator.
// less(a, b) must return true when a < b (strict).
func New[K, V any](less func(a, b K) bool) *Map[K, V] {
	return &Map[K, V]{less: less}
}

// Put inserts or updates the key-value pair. O(log n).
func (m *Map[K, V]) Put(key K, val V) {
	m.root = m.put(m.root, key, val)
	m.root.color = black
}

// Get returns the value for key and true if found. O(log n).
func (m *Map[K, V]) Get(key K) (V, bool) {
	n := m.root
	for n != nil {
		switch {
		case m.less(key, n.key):
			n = n.left
		case m.less(n.key, key):
			n = n.right
		default:
			return n.val, true
		}
	}
	var zero V
	return zero, false
}

// Contains reports whether key is present. O(log n).
func (m *Map[K, V]) Contains(key K) bool {
	_, ok := m.Get(key)
	return ok
}

// Delete removes the key from the map. O(log n).
func (m *Map[K, V]) Delete(key K) {
	if !m.Contains(key) {
		return
	}
	if !isRed(m.root.left) && !isRed(m.root.right) {
		m.root.color = red
	}
	m.root = m.delete(m.root, key)
	if m.root != nil {
		m.root.color = black
	}
}

// Len returns the number of key-value pairs. O(1).
func (m *Map[K, V]) Len() int { return sz(m.root) }

// Size is an alias for Len, matching C++ std::map::size().
func (m *Map[K, V]) Size() int { return sz(m.root) }

// Empty reports whether the map has no entries.
func (m *Map[K, V]) Empty() bool { return m.root == nil }

// IsEmpty is an alias for Empty.
func (m *Map[K, V]) IsEmpty() bool { return m.root == nil }

// Insert is an alias for Put, matching C++ std::map::insert().
func (m *Map[K, V]) Insert(key K, val V) { m.Put(key, val) }

// Find returns the value for key and true if found (mirrors C++ std::map::find()).
func (m *Map[K, V]) Find(key K) (V, bool) { return m.Get(key) }

// Erase is an alias for Delete, matching C++ std::map::erase().
func (m *Map[K, V]) Erase(key K) { m.Delete(key) }

// Min returns the smallest key and its value. O(log n).
func (m *Map[K, V]) Min() (K, V, bool) {
	if m.root == nil {
		var zk K
		var zv V
		return zk, zv, false
	}
	n := min(m.root)
	return n.key, n.val, true
}

// Max returns the largest key and its value. O(log n).
func (m *Map[K, V]) Max() (K, V, bool) {
	if m.root == nil {
		var zk K
		var zv V
		return zk, zv, false
	}
	n := m.root
	for n.right != nil {
		n = n.right
	}
	return n.key, n.val, true
}

// Floor returns the largest key ≤ given key. O(log n).
func (m *Map[K, V]) Floor(key K) (K, V, bool) {
	n := m.floor(m.root, key)
	if n == nil {
		var zk K
		var zv V
		return zk, zv, false
	}
	return n.key, n.val, true
}

// Ceiling returns the smallest key ≥ given key. O(log n).
func (m *Map[K, V]) Ceiling(key K) (K, V, bool) {
	n := m.ceiling(m.root, key)
	if n == nil {
		var zk K
		var zv V
		return zk, zv, false
	}
	return n.key, n.val, true
}

// Range iterates over all key-value pairs in ascending key order.
// Stop early by returning false from fn.
func (m *Map[K, V]) Range(fn func(key K, val V) bool) {
	m.inorder(m.root, fn)
}

// RangeFrom iterates over keys ≥ from in ascending order.
func (m *Map[K, V]) RangeFrom(from K, fn func(key K, val V) bool) {
	m.rangeFrom(m.root, from, fn)
}

// Keys returns all keys in ascending order.
func (m *Map[K, V]) Keys() []K {
	keys := make([]K, 0, sz(m.root))
	m.inorder(m.root, func(k K, _ V) bool {
		keys = append(keys, k)
		return true
	})
	return keys
}

// Values returns all values in key-ascending order.
func (m *Map[K, V]) Values() []V {
	vals := make([]V, 0, sz(m.root))
	m.inorder(m.root, func(_ K, v V) bool {
		vals = append(vals, v)
		return true
	})
	return vals
}

// ---- internal LLRB operations ----

func (m *Map[K, V]) put(n *node[K, V], key K, val V) *node[K, V] {
	if n == nil {
		return &node[K, V]{key: key, val: val, color: red, size: 1}
	}
	switch {
	case m.less(key, n.key):
		n.left = m.put(n.left, key, val)
	case m.less(n.key, key):
		n.right = m.put(n.right, key, val)
	default:
		n.val = val
	}
	return balance(n)
}

func (m *Map[K, V]) delete(n *node[K, V], key K) *node[K, V] {
	if m.less(key, n.key) {
		if !isRed(n.left) && !isRed(n.left.left) {
			n = moveRedLeft(n)
		}
		n.left = m.delete(n.left, key)
	} else {
		if isRed(n.left) {
			n = rotateRight(n)
		}
		if !m.less(n.key, key) && !m.less(key, n.key) && n.right == nil {
			return nil
		}
		if !isRed(n.right) && !isRed(n.right.left) {
			n = moveRedRight(n)
		}
		if !m.less(n.key, key) && !m.less(key, n.key) {
			small := min(n.right)
			n.key = small.key
			n.val = small.val
			n.right = deleteMin(n.right)
		} else {
			n.right = m.delete(n.right, key)
		}
	}
	return balance(n)
}

func (m *Map[K, V]) inorder(n *node[K, V], fn func(K, V) bool) bool {
	if n == nil {
		return true
	}
	if !m.inorder(n.left, fn) {
		return false
	}
	if !fn(n.key, n.val) {
		return false
	}
	return m.inorder(n.right, fn)
}

func (m *Map[K, V]) rangeFrom(n *node[K, V], from K, fn func(K, V) bool) bool {
	if n == nil {
		return true
	}
	if m.less(n.key, from) {
		return m.rangeFrom(n.right, from, fn)
	}
	if !m.rangeFrom(n.left, from, fn) {
		return false
	}
	if !fn(n.key, n.val) {
		return false
	}
	return m.rangeFrom(n.right, from, fn)
}

func (m *Map[K, V]) floor(n *node[K, V], key K) *node[K, V] {
	if n == nil {
		return nil
	}
	if !m.less(n.key, key) && !m.less(key, n.key) {
		return n
	}
	if m.less(key, n.key) {
		return m.floor(n.left, key)
	}
	t := m.floor(n.right, key)
	if t != nil {
		return t
	}
	return n
}

func (m *Map[K, V]) ceiling(n *node[K, V], key K) *node[K, V] {
	if n == nil {
		return nil
	}
	if !m.less(n.key, key) && !m.less(key, n.key) {
		return n
	}
	if m.less(n.key, key) {
		return m.ceiling(n.right, key)
	}
	t := m.ceiling(n.left, key)
	if t != nil {
		return t
	}
	return n
}

// ---- LLRB helper functions ----

func isRed[K, V any](n *node[K, V]) bool {
	return n != nil && n.color == red
}

func sz[K, V any](n *node[K, V]) int {
	if n == nil {
		return 0
	}
	return n.size
}

func setSize[K, V any](n *node[K, V]) {
	n.size = 1 + sz(n.left) + sz(n.right)
}

func rotateLeft[K, V any](h *node[K, V]) *node[K, V] {
	x := h.right
	h.right = x.left
	x.left = h
	x.color = h.color
	h.color = red
	setSize(h)
	setSize(x)
	return x
}

func rotateRight[K, V any](h *node[K, V]) *node[K, V] {
	x := h.left
	h.left = x.right
	x.right = h
	x.color = h.color
	h.color = red
	setSize(h)
	setSize(x)
	return x
}

func flipColors[K, V any](h *node[K, V]) {
	h.color = !h.color
	h.left.color = !h.left.color
	h.right.color = !h.right.color
}

func balance[K, V any](h *node[K, V]) *node[K, V] {
	if isRed(h.right) && !isRed(h.left) {
		h = rotateLeft(h)
	}
	if isRed(h.left) && isRed(h.left.left) {
		h = rotateRight(h)
	}
	if isRed(h.left) && isRed(h.right) {
		flipColors(h)
	}
	setSize(h)
	return h
}

func moveRedLeft[K, V any](h *node[K, V]) *node[K, V] {
	flipColors(h)
	if isRed(h.right.left) {
		h.right = rotateRight(h.right)
		h = rotateLeft(h)
		flipColors(h)
	}
	return h
}

func moveRedRight[K, V any](h *node[K, V]) *node[K, V] {
	flipColors(h)
	if isRed(h.left.left) {
		h = rotateRight(h)
		flipColors(h)
	}
	return h
}

func min[K, V any](n *node[K, V]) *node[K, V] {
	for n.left != nil {
		n = n.left
	}
	return n
}

func deleteMin[K, V any](h *node[K, V]) *node[K, V] {
	if h.left == nil {
		return nil
	}
	if !isRed(h.left) && !isRed(h.left.left) {
		h = moveRedLeft(h)
	}
	h.left = deleteMin(h.left)
	return balance(h)
}
