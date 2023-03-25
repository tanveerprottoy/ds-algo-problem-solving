package stack

type StackInt struct {
	data []int
}

func NewStackInt() *StackInt {
	return &StackInt{make([]int, 0)}
}

func (s *StackInt) IsEmpty() bool {
	return len(s.data) < 1
}

func (s *StackInt) Length() int {
	return len(s.data)
}

func (s *StackInt) Push(v int) {
	s.data = append(s.data, v)
}

func (s *StackInt) Pop() int {
	var res int
	l := s.Length()
	if l == 0 {
		return res
	}
	res = s.data[l-1]
	s.data = s.data[:l-1]
	return res
}