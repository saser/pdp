package fibheap

import (
	"constraints"
	"container/list"
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

func (h *Heap[T]) Len() int {
	return h.len
}

func (h *Heap[T]) Pop() T {
	h.len--
	var v T
	return v
}

func (h *Heap[T]) Push(v T) {
	h.len++
}
