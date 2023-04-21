package main

import (
	"github.com/tanveerprottoy/ds-algo-problem-solving/internal/ds/linkedlist"
)

func main() {
	l2 := linkedlist.NewLinkedList(&linkedlist.Node[int]{Val: 3, Next: nil})
	l1 := linkedlist.NewLinkedList(&linkedlist.Node[int]{Val: 2, Next: l2.Head})
	l0 := linkedlist.NewLinkedList(&linkedlist.Node[int]{Val: 1, Next: l1.Head})
	// l0.Traverse()
	l0.TraverseRecur(l0.Head)
}
