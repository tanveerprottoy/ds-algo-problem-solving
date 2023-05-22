package singly

import "fmt"

type Node[T any] struct {
	Val  T
	Next *Node[T]
}

func NewNode[T any](val T, next *Node[T]) *Node[T] {
	return &Node[T]{Val: val, Next: next}
}

type LinkedList[T any] struct {
	Head *Node[T]
}

func NewLinkedList[T any](head *Node[T]) *LinkedList[T] {
	return &LinkedList[T]{Head: head}
}

func (l *LinkedList[T]) Traverse() {
	tmp := l.Head
	for tmp.Next != nil {
		fmt.Println("val: ", tmp.Val)
		fmt.Println("next: ", tmp.Next)
		tmp = tmp.Next
	}
}

func (l *LinkedList[T]) TraverseRecur(node *Node[T]) {
	if node.Next == nil {
		return
	}
	fmt.Println("val: ", node.Val)
	fmt.Println("next: ", node.Next)
	l.TraverseRecur(node.Next)
}

func (l *LinkedList[T]) InsertAtHead(e T) {
	nxt := l.Head.Next
	n := NewNode(e, nxt)
	l.Head = n
}

func (l *LinkedList[T]) InsertAtTail(e T) {
	tmp := l.Head
	for tmp.Next != nil {
		tmp = tmp.Next
	}
	n := NewNode(e, nil)
	tmp.Next = n
}

func (l *LinkedList[T]) InsertAtPosition(e T, pos int) {

}

func (l *LinkedList[T]) Find(e T) {

}

func (l *LinkedList[T]) Sort() {

}

func (l *LinkedList[T]) Delete(e T) {

}
