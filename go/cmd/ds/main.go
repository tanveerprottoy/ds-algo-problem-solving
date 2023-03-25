package main

import (
	"github.com/tanveerprottoy/ds-algo-problem-solving/internal/ds/linkedlist"
)

func main() {
	l2 := linkedlist.NewLinkedList(&linkedlist.Node[int]{Val: 3, Nxt: nil})
	l1 := linkedlist.NewLinkedList(&linkedlist.Node[int]{Val: 2, Nxt: l2.Head})
	l0 := linkedlist.NewLinkedList(&linkedlist.Node[int]{Val: 1, Nxt: l1.Head})
	l0.Traverse()
}
