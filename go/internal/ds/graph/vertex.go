package graph

type Vertex[T any] struct {
	val       T
	adjacents []Vertex[T]
}

func NewVertex[T any](val T, adjacents []Vertex[T]) *Vertex[T] {
	return &Vertex[T]{val: val, adjacents: adjacents}
}
