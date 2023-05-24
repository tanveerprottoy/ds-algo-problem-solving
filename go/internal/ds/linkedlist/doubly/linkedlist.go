package doubly

import "fmt"

type Node struct {
	Val      any
	Previous *Node
	Next     *Node
}

func NewNode(val any, previous, next *Node) *Node {
	return &Node{Val: val, Previous: previous, Next: next}
}

type LinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

func NewLinkedList(head, tail *Node) *LinkedList {
	return &LinkedList{Head: head, Tail: tail}
}

func (l *LinkedList) Traverse() {
	curr := l.Head
	for curr != nil {
		fmt.Println("val: ", curr.Val)
		fmt.Println("Next: ", curr.Next)
		curr = curr.Next
	}
}

func (l *LinkedList) TraverseRecur(node *Node) {
	if node.Next == nil {
		return
	}
	fmt.Println("val: ", node.Val)
	fmt.Println("Next: ", node.Next)
	l.TraverseRecur(node.Next)
}

func (l *LinkedList) Size() int {
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

func (l *LinkedList) InsertAtHead(v any) {
	nxt := l.Head.Next
	n := NewNode(v, nil, nxt)
	l.Head = n
	l.Length++
}

func (l *LinkedList) InsertAtTail(v any) {
	curr := l.Head
	for curr != nil {
		curr = curr.Next
	}
	n := NewNode(v, curr, nil)
	curr.Next = n
	l.Tail = n
	l.Length++
}

func (l *LinkedList) InsertAtPosition(v any, pos int) {
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

func (l *LinkedList) InsertAtMiddle(v any) {
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

func (l *LinkedList) Find(v any) *Node {
	curr := l.Head
	for curr != nil {
		switch v.(type) {
		case int:
			fmt.Print(v == curr.Val.(int))
			if v == curr.Val.(int) {
				return curr
			}
		}
		// val, ok := v.(int)
		curr = curr.Next
	}
	return nil
}

func (l *LinkedList) DeleteAtHead() *Node {
	curr := l.Head
	l.Head = l.Head.Next
	l.Length--
	return curr
}

func (l *LinkedList) DeleteAtTail() *Node {
	if l.Head == nil {
		return nil
	}
	var del *Node
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

func (l *LinkedList) Delete(v any) *Node {
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
				if v == fast.Val.(int) {
					slow.Next = fast.Next
					l.Length--
					return fast
				} else if v == slow.Val.(int) {
					slow = slow.Next
					l.Length--
					return slow
				}
			}
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return nil
}
