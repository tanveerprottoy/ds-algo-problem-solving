package linkedlist

import "fmt"

type Node[T any] struct {
	Val T
	Nxt *Node[T]
}

func NewNode[T any](val T, nxt *Node[T]) *Node[T] {
	return &Node[T]{Val: val, Nxt: nxt}
}

func (n *Node[T]) Traverse(head *Node[T]) {
	for head != nil {
		fmt.Println("val: ", head.Val)
		fmt.Println("nxt: ", head.Nxt)
		head = head.Nxt
	}
}

func (n *Node[T]) TraverseRecur(head *Node[T]) {
	if head == nil {
		return
	}
	fmt.Println("val: ", head.Val)
	fmt.Println("nxt: ", head.Nxt)
	n.TraverseRecur(head.Nxt)
}
