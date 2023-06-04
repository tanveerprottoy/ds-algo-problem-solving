/**
Max heap: the value of child nodes is smaller than or equal to the value of its
parent. The root node has the maximum value.

array representation of heap
 vertex            ║ index                  ║
╠═══════════════════╬════════════════════════╣
║ root              ║ 0                      ║ 
║ current           ║ i                      ║
║ parent            ║ (i - 1) / 2            ║
║ left child        ║ 2*i + 1                ║
║ right child       ║ 2*i + 2                ║
║ the last non-leaf ║ (array length - 2) / 2 ║

*/

package heap

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type MaxHeap[T constraints.Ordered] struct {
	data []T
}

func NewMaxHeap[T constraints.Ordered]() *MaxHeap[T] {
	return &MaxHeap[T]{data: make([]T, 0)}
}

// Push adds an element to the heap
func (h *MaxHeap[T]) Push(key T) {
	h.data = append(h.data, key)
	h.HeapifyUp(len(h.data) - 1)
}

// Pop returns the largest key, and removes it from heap
func (h *MaxHeap[T]) Pop() T {
	v := h.data[0]
	l := len(h.data) - 1
	if len(h.data) == 0 {
		fmt.Println("Cannot extract because data lenth is 0")
		return v
	}
	h.data[0] = h.data[l]
	h.data = h.data[:l]
	h.HeapifyDown(0)
	return v
}

// HeapifyUp process
func (h *MaxHeap[T]) HeapifyUp(index int) {
	for h.data[Parent(index)] < h.data[index] {
		h.swap(Parent(index), index)
		index = Parent(index)
	}
}

// HeapifyDown process
func (h *MaxHeap[T]) HeapifyDown(index int) {
	lastIndex := len(h.data) - 1
	l, r := Left(index), Right(index)
	childToCompare := 0
	// loop while index has at least one child
	for l <= lastIndex { // when left child is the only child
		if l == lastIndex {
			childToCompare = l
		} else if h.data[l] > h.data[r] { // when left child is larger
			childToCompare = l
		} else { // when right child is larger
			childToCompare = r
		}
		// Compare data value of the current index to larger child and swap if smaller
		if h.data[index] < h.data[childToCompare] {
			h.swap((index), childToCompare)
			index = childToCompare
			l, r = Left(index), Right(index)
		} else {
			return
		}
	}
}

func heapify(heap *[]int, i int) {
	smallest := i
	lChild := 2*i + 1
	rChild := 2*i + 2

	if lChild < len(*heap) && (*heap)[lChild] < (*heap)[smallest] {
		smallest = lChild
	}
	if rChild < len(*heap) && (*heap)[rChild] < (*heap)[smallest] {
		smallest = rChild
	}

	if smallest != i {
		(*heap)[i], (*heap)[smallest] = (*heap)[smallest], (*heap)[i]
		heapify(heap, smallest)
	}
}

// Swap keys in the data
func (h *MaxHeap[T]) swap(i1, i2 int) {
	h.data[i1], h.data[i2] = h.data[i2], h.data[i1]
}
