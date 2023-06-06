/*
Max heap: the value of child nodes is smaller than or equal to the value of its
parent. The root node has the maximum value.

h.data representation of heap
 vertex            ║ index                  ║
╠═══════════════════╬════════════════════════╣
║ root              ║ 0                      ║
║ current           ║ i                      ║
║ parent            ║ (i - 1) / 2            ║
║ left child        ║ 2*i + 1                ║
║ right child       ║ 2*i + 2                ║
║ the last non-leaf ║ (h.data length - 2) / 2 ║


Root is at index 0 of the h.data.
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

import "golang.org/x/exp/constraints"

type MaxHeap[T constraints.Ordered] struct {
	data []T
}

func NewMaxHeap[T constraints.Ordered](data []T) *MaxHeap[T] {
	return &MaxHeap[T]{data: make([]T, 0)}
}

// Time: O(n) | Space: O(1)
// heapify an h.data
func (h *MaxHeap[T]) Heapify() {
	lastNonLeafNodeIdx := (len(h.data) - 2) / 2
	for currentIdx := lastNonLeafNodeIdx; currentIdx >= 0; currentIdx-- {
		endIdx := len(h.data) - 1
		h.siftDown(currentIdx, endIdx) // siftDown is more efficient than siftUp
	}
}

// Time: O(logn) | Space: O(1)
// continuously swap the current node with its smaller child node until it is in the correct position
func (h *MaxHeap[T]) siftDown(currentIdx int, endIdx int) {
	leftChildIdx := currentIdx*2 + 1
	for leftChildIdx <= endIdx {
		rightChildIdx := currentIdx*2 + 2
		if rightChildIdx > endIdx {
			rightChildIdx = -1
		}

		// get the larger child node to swap
		idxToSwap := leftChildIdx
		if rightChildIdx != -1 && h.data[rightChildIdx] > h.data[leftChildIdx] {
			idxToSwap = rightChildIdx
		}

		// check if value of swap node is greater than the value at currentIdx
		if h.data[idxToSwap] > h.data[currentIdx] {
			h.swap(idxToSwap, currentIdx)
			currentIdx = idxToSwap
			leftChildIdx = currentIdx*2 + 1

		} else {
			return
		}
	}
}

// Time: O(logn) | Space: O(1)
// continuously swap the current node with its parent node until it is in the correct position
func (h *MaxHeap[T]) siftUp() {
	currentIdx := len(h.data) - 1
	parentIdx := (currentIdx - 1) / 2
	for currentIdx > 0 && h.data[currentIdx] > h.data[parentIdx] {
		h.swap(currentIdx, parentIdx)
		currentIdx = parentIdx
		parentIdx = (currentIdx - 1) / 2
	}
}

// Time: O(logn) | Space: O(1)
// insert a new value to the end of the tree and update heap ordering
func (h *MaxHeap[T]) Insert(value T) {
	h.data = append(h.data, value)
	h.siftUp()
}

// Time: O(logn) | Space: O(1)
// remove and return the minimum value and update heap ordering
func (h *MaxHeap[T]) Remove() T {
	n := len(h.data)
	// swap the first element and the last element in the h.data
	h.swap(0, n-1)
	valueToRemove := h.data[n-1]
	// pop the last element in the h.data
	h.data = h.data[:n-1]
	// call siftDown to update heap ordering
	h.siftDown(0, n-2)

	return valueToRemove
}

// Time: O(1) | Space: O(1)
func (h MaxHeap[T]) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}
