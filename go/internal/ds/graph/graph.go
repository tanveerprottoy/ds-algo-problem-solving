package graph

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Graph[T constraints.Ordered] struct {
	vertices []Vertex[T]
}

func (g *Graph[T]) AddVertex(v T) error {
	if contains(g.vertices, v) {
		err := fmt.Errorf("Vertex %d already exists", v)
		return err
	} else {
		v := &Vertex[T]{
			val: v,
		}
		g.vertices = append(g.vertices, *v)
	}
	return nil
}

// AddEdge will add ad endge from a vertex to a vertex
func (g *Graph[T]) AddEdge(to, from T) error {
	toVertex := g.getVertex(to)
	fromVertex := g.getVertex(from)
	if toVertex == nil || fromVertex == nil {
		return fmt.Errorf("Not a valid edge from %d ---> %d", from, to)
	} else if contains(fromVertex.adjacents, toVertex.val) {
		return fmt.Errorf("Edge from vertex %d ---> %d already exists", fromVertex.val, toVertex.val)
	} else {
		fromVertex.adjacents = append(fromVertex.adjacents, *toVertex)
		return nil
	}
}

// getVertex will return a vertex point if exists or return nil
func (g *Graph[T]) getVertex(v T) *Vertex[T] {
	for i, v1 := range g.vertices {
		if v1.val == v {
			return &g.vertices[i]
		}
	}
	return nil
}

func contains[T constraints.Ordered](s []Vertex[T], v T) bool {
	for _, v1 := range s {
		if v1.val == v {
			return true
		}
	}
	return false
}

func (g *Graph[T]) Print() {
	for _, v := range g.vertices {
		fmt.Printf("%d : ", v.val)
		for _, v := range v.adjacents {
			fmt.Printf("%d ", v.val)
		}
		fmt.Println()
	}
}
