package singly

import (
	"fmt"
)

type Node[T any] struct {
	Val  any
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

func (l *LinkedList[T]) InsertAtHead(v T) {
	nxt := l.Head.Next
	n := NewNode(v, nxt)
	l.Head = n
}

func (l *LinkedList[T]) InsertAtTail(v T) {
	tmp := l.Head
	for tmp.Next != nil {
		tmp = tmp.Next
	}
	n := NewNode(v, nil)
	tmp.Next = n
}

func (l *LinkedList[T]) InsertAtPosition(v T, pos int) {
	i := 0
	tmp := l.Head
	for i < pos && tmp.Next != nil {
		if i == pos-1 {
			nxt := tmp.Next
			n := NewNode(v, nxt)
			tmp.Next = n
		}
		i++
	}
}

func (l *LinkedList[T]) Find(v any) *Node[T] {
	tmp := l.Head
	for tmp.Next != nil {
		switch v.(type) {
		case int:
			fmt.Print(v == tmp.Val.(int))
			if(v == tmp.Val.(int)) {
				return tmp
			}
		}
		// val, ok := v.(int)
	}
	return nil
}

func (l *LinkedList[T]) Sort() {

}

func (l *LinkedList[T]) Delete(v T) {

}
