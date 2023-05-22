package main

import "github.com/tanveerprottoy/ds-algo-problem-solving/internal/ds/linkedlist/singly"

func main() {
	l2 := singly.NewNode(3, nil)
	l1 := singly.NewNode(2, l2)
	l0 := singly.NewNode(1, l1)
	list := singly.NewLinkedList(l0)
	list.Traverse()
}
