package main

import (
	"github.com/tanveerprottoy/ds-algo-problem-solving/internal/ds"
)

func main() {
	l2 := new(ds.BasicLinkedList)
	l2.Val = 3
	l1 := new(ds.BasicLinkedList)
	l1.Val = 2
	l1.Nxt = l2
	l0 := new(ds.BasicLinkedList)
	l0.Val = 1
	l0.Nxt = l1
	l0.Traverse()
}
