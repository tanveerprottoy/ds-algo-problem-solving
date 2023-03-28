package linkedlist

import "fmt"

type LinkedList[T any] struct {
	Head *Node[T]
}

func NewLinkedList[T any](head *Node[T]) *LinkedList[T] {
	return &LinkedList[T]{Head: head}
}

func (l *LinkedList[T]) Traverse() {
	tmp := l.Head
	for tmp.Nxt != nil {
		fmt.Println("val: ", tmp.Val)
		fmt.Println("next: ", tmp.Nxt)
		tmp = tmp.Nxt
	}
}

func (l *LinkedList[T]) TraverseRecur(node *Node[T]) {
	if node.Nxt == nil {
		return
	}
	fmt.Println("val: ", node.Val)
	fmt.Println("next: ", node.Nxt)
	l.TraverseRecur(node.Nxt)
}

func (l *LinkedList[T]) Insert(e T, pos int) {

}

func (l *LinkedList[T]) Find(e T)  {

}

func (l *LinkedList[T]) Sort() {

}

func (l *LinkedList[T]) Delete(e T) {

}
