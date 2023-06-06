/**
Min heap: the value of child nodes is greater than or equal to the value of
its parent. The root node has the minimum value.

h.data representation of heap
 vertex            ║ index                  ║
╠═══════════════════╬════════════════════════╣
║ root              ║ 0                      ║
║ current           ║ i                      ║
║ parent            ║ (i - 1) / 2            ║
║ left child        ║ 2*i + 1                ║
║ right child       ║ 2*i + 2                ║
║ the last non-leaf ║ (h.data length - 2) / 2 ║

BuildHeap: all non-leaf nodes call SiftDown to heapify the h.data. We starts with the last non-leaf node.
SiftDown: continuously swap the current node with its smaller child node until it is in the correct position. (Move down)
SiftUp: continuously swap the current node with its parent node until it is in the correct position. (Move up)
Remove: remove the root node by swap the root node with the last node in the tree and then pop the root node which is currently at the end of tree . Call SiftDown to update heap ordering
Insert: insert a new value at the end of the tree and call SiftUp to update heap ordering

*/

package heap

import "golang.org/x/exp/constraints"

type MinHeap[T constraints.Ordered] struct {
	data []T
}

func NewMinHeap[T constraints.Ordered](data []T) *MinHeap[T] {
	return &MinHeap[T]{data: data}
}

// Time: O(n) | Space: O(1)
// heapify an h.data
func (h *MinHeap[T]) Heapify() {
	lastNonLeafNodeIdx := (len(h.data) - 2) / 2
	for currentIdx := lastNonLeafNodeIdx; currentIdx >= 0; currentIdx-- {
		endIdx := len(h.data) - 1
		h.siftDown(currentIdx, endIdx) // siftDown is more efficient than siftUp
	}
}

// Time: O(logn) | Space: O(1)
// continuously swap the current node with its smaller child node until it is in the correct position
func (h *MinHeap[T]) siftDown(currentIdx int, endIdx int) {
	leftChildIdx := currentIdx*2 + 1
	for leftChildIdx <= endIdx {
		rightChildIdx := currentIdx*2 + 2
		if rightChildIdx > endIdx {
			rightChildIdx = -1
		}

		// get the smaller child node to swap
		idxToSwap := leftChildIdx
		if rightChildIdx != -1 && h.data[rightChildIdx] < h.data[leftChildIdx] {
			idxToSwap = rightChildIdx
		}

		// check if value of swap node is less than the value at currentIdx
		if h.data[idxToSwap] < h.data[currentIdx] {
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
func (h *MinHeap[T]) siftUp() {
	currentIdx := len(h.data) - 1
	parentIdx := (currentIdx - 1) / 2
	for currentIdx > 0 && h.data[currentIdx] < h.data[parentIdx] {
		h.swap(currentIdx, parentIdx)
		currentIdx = parentIdx
		parentIdx = (currentIdx - 1) / 2
	}
}

// Time: O(logn) | Space: O(1)
// insert a new value to the end of the tree and update heap ordering
func (h *MinHeap[T]) Insert(v T) {
	h.data = append(h.data, v)
	h.siftUp()
}

// Time: O(logn) | Space: O(1)
// remove and return the minimum value and update heap ordering
func (h *MinHeap[T]) Remove() T {
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
func (h MinHeap[T]) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}
