package DoublyLinkedList

import "fmt"

type Node[T any] struct {
	Val  T
	Previous *Node[T]
	Next *Node[T]
}

func NewNode[T any](val T, next *Node[T]) *Node[T] {
	return &Node[T]{Val: val, Next: next}
}

type DoublyLinkedList[T any] struct {
	Head *Node[T]
}

func NewDoublyLinkedList[T any](head *Node[T]) *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{Head: head}
}

func (l *DoublyLinkedList[T]) Traverse() {
	tmp := l.Head
	for tmp.Next != nil {
		fmt.Println("val: ", tmp.Val)
		fmt.Println("next: ", tmp.Next)
		tmp = tmp.Next
	}
}

func (l *DoublyLinkedList[T]) TraverseRecur(node *Node[T]) {
	if node.Next == nil {
		return
	}
	fmt.Println("val: ", node.Val)
	fmt.Println("next: ", node.Next)
	l.TraverseRecur(node.Next)
}

func (l *DoublyLinkedList[T]) InsertAtHead(e T) {
	nxt := l.Head.Next
	n := NewNode(e, nxt)
	l.Head = n
}

func (l *DoublyLinkedList[T]) InsertAtTail(e T) {
	tmp := l.Head
	for tmp.Next != nil {
		tmp = tmp.Next
	}
	n := NewNode(e, nil)
	tmp.Next = n
}

func (l *DoublyLinkedList[T]) InsertAtPosition(e T, pos int) {

}

func (l *DoublyLinkedList[T]) Find(e T) {

}

func (l *DoublyLinkedList[T]) Sort() {

}

func (l *DoublyLinkedList[T]) Delete(e T) {

}
