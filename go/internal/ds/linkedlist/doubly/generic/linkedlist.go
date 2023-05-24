package generic

import "fmt"

type Node[T any] struct {
	Val      T
	Previous *Node[T]
	Next     *Node[T]
}

func NewNode[T any](val T, previous, next *Node[T]) *Node[T] {
	return &Node[T]{Val: val, Previous: previous, Next: next}
}

type LinkedList[T any] struct {
	Head   *Node[T]
	Tail   *Node[T]
	Length int
}

func NewLinkedList[T any](head, tail *Node[T]) *LinkedList[T] {
	return &LinkedList[T]{Head: head, Tail: tail}
}

func (l *LinkedList[T]) Traverse() {
	curr := l.Head
	for curr != nil {
		fmt.Println("val: ", curr.Val)
		fmt.Println("Next: ", curr.Next)
		curr = curr.Next
	}
}

func (l *LinkedList[T]) TraverseRecur(node *Node[T]) {
	if node.Next == nil {
		return
	}
	fmt.Println("val: ", node.Val)
	fmt.Println("Next: ", node.Next)
	l.TraverseRecur(node.Next)
}

func (l *LinkedList[T]) Size() int {
	if l.Head == nil {
		return 0
	}
	i := 1
	curr := l.Head
	for curr != nil {
		i++
		curr = curr.Next
	}
	return i
}

func (l *LinkedList[T]) InsertAtHead(v T) {
	nxt := l.Head.Next
	n := NewNode(v, nil, nxt)
	l.Head = n
	l.Length++
}

func (l *LinkedList[T]) InsertAtTail(v T) {
	curr := l.Head
	for curr != nil {
		curr = curr.Next
	}
	n := NewNode(v, curr, nil)
	curr.Next = n
	l.Tail = n
	l.Length++
}

func (l *LinkedList[T]) InsertAtPosition(v T, pos int) {
	i := 0
	curr := l.Head
	for i < pos && curr != nil {
		if i == pos-1 {
			nxt := curr.Next
			n := NewNode(v, curr, nxt)
			curr.Next = n
			l.Length++
			return
		}
		i++
		curr = curr.Next
	}
}

func (l *LinkedList[T]) InsertAtMiddle(v T) {
	mid := l.Size() / 2
	i := 0
	curr := l.Head
	for curr != nil {
		if i == mid-1 {
			n := NewNode(v, curr.Previous, curr.Next)
			n.Next = curr
			l.Length++
			return
		}
		curr = curr.Next
		i++
	}
}

func (l *LinkedList[T]) Find(v any) *Node[T] {
	curr := l.Head
	for curr != nil {
		switch v.(type) {
		case int:
			/* fmt.Print(v == curr.Val.(int))
			if v == curr.Val.(int) {
				return curr
			} */
		}
		// val, ok := v.(int)
		curr = curr.Next
	}
	return nil
}

func (l *LinkedList[T]) DeleteAtHead() *Node[T] {
	curr := l.Head
	l.Head = l.Head.Next
	l.Length--
	return curr
}

func (l *LinkedList[T]) DeleteAtTail() *Node[T] {
	if l.Head == nil {
		return nil
	}
	var del *Node[T]
	// corener case for only 1 item
	if l.Head.Next == nil {
		del := l.Head
		l.Head = nil
		l.Length--
		return del
	}
	curr := l.Head
	for curr != nil {
		curr = curr.Next
	}
	del = curr
	curr = nil
	l.Length--
	return del
}

func (l *LinkedList[T]) Delete(v any) *Node[T] {
	if l.Head == nil {
		return nil
	}
	if l.Head.Next == nil {
		tmp := l.Head
		l.Head = nil
		l.Length--
		return tmp
	}
	// slow will be the previous
	// node of fast
	fast := l.Head.Next
	slow := l.Head
	// check for conrner case
	// when head is the target value
	/* if v == slow.Val.(int) {
		slow = slow.Next
	} */
	for fast != nil && slow != nil {
		switch v.(type) {
		case int:
			{
				/* if v == fast.Val.(int) {
					slow.Next = fast.Next
					l.Length--
					return fast
				} else if v == slow.Val.(int) {
					slow = slow.Next
					l.Length--
					return slow
				} */
			}
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return nil
}
