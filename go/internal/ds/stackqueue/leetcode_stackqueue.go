package stackqueue

type LeetcodeStack struct {
	data []int
}

func NewStack() *LeetcodeStack {
	return &LeetcodeStack{}
}

func (s *LeetcodeStack) Push(v int) {
	s.data = append(s.data, v)
}

func (s *LeetcodeStack) Pop() int {
	l := len(s.data)
	v := s.data[l-1]
	s.data = s.data[:l-1]
	return v
}

func (s *LeetcodeStack) Peek() int {
	if s.IsEmpty() {
		return -1
	}
	return s.data[len(s.data)-1]
}

func (s *LeetcodeStack) Size() int {
	return len(s.data)
}

func (s *LeetcodeStack) IsEmpty() bool {
	return len(s.data) < 1
}

type LeetcodeQueue struct {
	s0 *LeetcodeStack
	s1 *LeetcodeStack
}

func Constructor() LeetcodeQueue {
	return LeetcodeQueue{new(LeetcodeStack), new(LeetcodeStack)}
}

func (this *LeetcodeQueue) Push(x int) {
	for !this.s0.IsEmpty() {
		this.s1.Push(this.s0.Pop())
	}
	this.s0.Push(x)
	for !this.s1.IsEmpty() {
		this.s0.Push(this.s1.Pop())
	}
}

func (this *LeetcodeQueue) Pop() int {
	return this.s0.Pop()
}

func (this *LeetcodeQueue) Peek() int {
	return this.s0.Peek()
}

func (this *LeetcodeQueue) Empty() bool {
	return this.s0.IsEmpty()
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
