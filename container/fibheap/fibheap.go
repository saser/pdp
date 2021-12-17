package fibheap

import (
	"constraints"
	"container/list"

	"github.com/Saser/pdp/adventofcode/go/intmath"
)

type Heap[T any] struct {
	less  func(i, j T) bool
	roots *list.List    // list of *node[T]
	min   *list.Element // contains a *node[T]
	len   int           // the number of elements in the heap
}

func New[T constraints.Ordered]() *Heap[T] {
	return NewFunc(func(i, j T) bool { return i < j })
}

func NewFunc[T any](less func(i, j T) bool) *Heap[T] {
	return &Heap[T]{
		less:  less,
		roots: list.New(),
		min:   nil,
		len:   0,
	}
}

type node[T any] struct {
	value    T
	marked   bool
	parent   *node[T]
	children []*node[T]
}

func (n *node[T]) rank() int {
	return len(n.children)
}

func (n *node[T]) cutChildren() []*node[T] {
	cut := make([]*node[T], 0, len(n.children))
	for _, child := range n.children {
		child.parent = nil
		cut = append(cut, child)
	}
	return cut
}

func (h *Heap[T]) Len() int {
	return h.len
}

func (h *Heap[T]) Pop() T {
	if h.len == 0 {
		panic("popping from empty")
	}
	h.len--

	// Remove the minimum node and meld its children into the list of roots.
	minNode := h.roots.Remove(h.min).(*node[T])
	popped := minNode.value
	for _, child := range minNode.cutChildren() {
		h.roots.PushFront(child)
	}

	// Update the minimum element. If the heap is now empty, return early.
	minElement := h.roots.Front()
	if minElement == nil {
		h.min = nil
		return popped
	}
	// While we're finding the minimum element, also find the max rank.
	maxRank := 0
	for e := minElement.Next(); e != nil; e = e.Next() {
		current := minElement.Value.(*node[T])
		candidate := e.Value.(*node[T])
		if h.less(candidate.value, current.value) {
			minElement = e
		}
		maxRank = intmath.Max(maxRank, current.rank(), candidate.rank())
	}
	h.min = minElement

	rootByRank := make([]*list.Element, maxRank) // rank -> root element with that rank
	for e := h.roots.Front(); e != nil; e = e.Next() {
		///////////////////////////////////////////////////////////////////////
		// Too tired right now, but we should be looping here until there is no
		// existing node with the same rank as the one we're currently on.
		///////////////////////////////////////////////////////////////////////
		rootElement := e
		rootNode := rootElement.Value.(*node[T])
		rank := rootNode.rank()
		existingElement := rootByRank[rank]
		if existingElement == nil {
			rootByRank[rank] = e
			continue
		}
		existingNode := existingElement.Value.(*node[T])
		if h.less(rootNode.value, existingNode.value) {
			rootNode.parent = existingNode
			existingNode.children = append(existingNode.children, rootNode)
		} else {

		}
	}

	return popped
}

func (h *Heap[T]) Push(v T) {
	h.len++

	root := h.roots.PushFront(&node[T]{
		value:    v,
		marked:   false,
		parent:   nil,
		children: nil,
	})
	if min, ok := h.minval(); !ok || h.less(v, min) {
		h.min = root
	}
}

// minval returns the minimum value of the heap. If there is no minimum value
// (i.e., the heap is empty), minval returns the zero-value and false.
func (h *Heap[T]) minval() (T, bool) {
	if h.min == nil {
		var v T
		return v, false
	}
	return h.min.Value.(*node[T]).value, true
}
