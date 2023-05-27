package main

import (
	"fmt"

	"github.com/tanveerprottoy/ds-algo-problem-solving/internal/ds/stackqueue"
)

func main() {
	/* l2 := generic.NewNode(3, nil)
	l1 := generic.NewNode(2, l2)
	l0 := generic.NewNode(1, l1)
	list := generic.NewLinkedList(l0, nil)
	list.Traverse() */
	m := stackqueue.Constructor()
	m.Push(1)
	m.Push(2)
	fmt.Println(m.Peek())
	fmt.Println(m.Pop())
	fmt.Println(m.Empty())
}
