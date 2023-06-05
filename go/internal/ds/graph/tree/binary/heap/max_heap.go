/*
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


Root is at index 0 of the array.
Left child of index i is at (2*i + 1).
Right child of index i is at (2*i + 2).
Parent of index i is at (i-1)/2.
Example:
Input: []int{4, 12, 3, 6, 5}
Max-Heap: []int{12,6,3,4,5}
        12
       /  \
      6    3
     / \
    4   5
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
	h.SiftUp(len(h.data) - 1)
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
	h.SiftDown(0)
	return v
}

func (h *MaxHeap[T]) Heapify(index int) {
	smallest := index
	lChild := Left(index)
	rChild := Right(index)
	if lChild < len(h.data) && h.data[lChild] < h.data[smallest] {
		smallest = lChild
	}
	if rChild < len(h.data) && h.data[rChild] < h.data[smallest] {
		smallest = rChild
	}
	if smallest != index {
		h.swap(index, smallest)
		h.Heapify(smallest)
	}
}

/*
Sift up: move the value up the tree by successively exchanging
the value with its parent node.
*/
func (h *MaxHeap[T]) SiftUp(index int) {
	for h.data[Parent(index)] < h.data[index] {
		h.swap(Parent(index), index)
		index = Parent(index)
	}
}

/*
Sift down: move the value down the tree by successively exchanging the value
with its smaller(for min heap)/larger(for max heap) child node.
*/
func (h *MaxHeap[T]) SiftDown(index int) {
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

// Swap keys in the data
func (h *MaxHeap[T]) swap(i1, i2 int) {
	h.data[i1], h.data[i2] = h.data[i2], h.data[i1]
}
