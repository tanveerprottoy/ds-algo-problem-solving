package linkedlist

import "fmt"

type Node[T any] struct {
	Val  T
	Next *Node[T]
}

func NewNode[T any](val T, nxt *Node[T]) *Node[T] {
	return &Node[T]{Val: val, Next: nxt}
}

func (n *Node[T]) Traverse(head *Node[T]) {
	for head != nil {
		fmt.Println("val: ", head.Val)
		fmt.Println("nxt: ", head.Next)
		head = head.Next
	}
}

func (n *Node[T]) TraverseRecur(head *Node[T]) {
	if head == nil {
		return
	}
	fmt.Println("val: ", head.Val)
	fmt.Println("nxt: ", head.Next)
	n.TraverseRecur(head.Next)
}
