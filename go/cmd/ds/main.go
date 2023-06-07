package main

import (
	"fmt"

	"github.com/tanveerprottoy/ds-algo-problem-solving/internal/ds/graph/tree/binary"
)

func main() {
	/* l2 := generic.NewNode(3, nil)
	l1 := generic.NewNode(2, l2)
	l0 := generic.NewNode(1, l1)
	list := generic.NewLinkedList(l0, nil)
	list.Traverse() */
	/* m := stackqueue.Constructor()
	m.Push(1)
	m.Push(2)
	fmt.Println(m.Peek())
	fmt.Println(m.Pop())
	fmt.Println(m.Empty()) */
	/* data := []int{9, 31, 40, 22, 10, 15, 1, 25, 91}
	minHeap := heap.NewMinHeap[int](data)
	minHeap.Heapify()
	fmt.Println("build min heap: ", *minHeap)

	// apply Remove method
	valueToRemove := minHeap.Remove()
	fmt.Println("root value: ", valueToRemove)
	fmt.Println("min heap after Remove root: ", *minHeap)

	// apply Insert method
	// append a new value, 50
	minHeap.Insert(50)
	fmt.Println("min heap after insert value 50: ", *minHeap) */

	tn := &binary.TreeNode{Val: 2, Left: &binary.TreeNode{Val: 1}, Right: &binary.TreeNode{Val: 3}}
	fmt.Println(binary.InorderTraversal(tn))
}
