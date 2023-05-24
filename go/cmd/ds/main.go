package main

import "github.com/tanveerprottoy/ds-algo-problem-solving/internal/ds/linkedlist/singly/generic"

func main() {
	l2 := generic.NewNode(3, nil)
	l1 := generic.NewNode(2, l2)
	l0 := generic.NewNode(1, l1)
	list := generic.NewLinkedList(l0, nil)
	list.Traverse()
}
