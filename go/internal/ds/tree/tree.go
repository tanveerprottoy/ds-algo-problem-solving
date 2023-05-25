package tree

type Tree[T any] struct {
	Node T
	Edge *T
}