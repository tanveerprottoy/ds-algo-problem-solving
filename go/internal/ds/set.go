package ds

type Set[T any] struct {
	data map[string]T
}

func NewSet[T any]() *Set[T] {
	return &Set[T]{make(map[string]T)}
}

func (s *Set[T]) Add(k string, v T) {
	s.data[k] = v
}

func (s *Set[T]) Delete(k string) {
	delete(s.data, k)
}

func (s *Set[T]) Has(k string) bool {
	_, ok := s.data[k]
	return ok
}
