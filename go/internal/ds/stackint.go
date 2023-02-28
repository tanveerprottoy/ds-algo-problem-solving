package ds

type StackInt struct {
	d []int
}

func NewStackInt() *StackInt {
	return &StackInt{make([]int, 0)}
}

func (s *StackInt) IsEmpty() bool {
	return len(s.d) == 0
}

func (s *StackInt) Length() int {
	return len(s.d)
}

func (s *StackInt) Push(v int) {
	s.d = append(s.d, v)
}

func (s *StackInt) Pop() int {
	var res int
	l := s.Length()
	if l == 0 {
		return res
	}
	res = s.d[l-1]
	s.d = s.d[:l-1]
	return res
}
