package ds

type Set[T any] struct {
	data map[string]T
}

func NewSet[T any]() *Set[T] {
	return &Set[T]{make(map[string]T)}
}

func (s *Set[T]) Add(v T) {
	
}
