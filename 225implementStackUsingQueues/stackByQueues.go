package implementstackusingqueues

type MyStack struct {
	queue []int
}

func Constructor() MyStack {
	queue := []int{}
	s := MyStack{queue: queue}
	return s

}

func (s *MyStack) Push(x int) {
	s.queue = append(s.queue, x)

	n := len(s.queue)
	for i := 0; i < n-1; i++ {
		val := s.Pop()
		s.queue = append(s.queue, val)
	}
}

func (s *MyStack) Pop() int {
	if len(s.queue) != 0 {
		val := s.queue[0]
		s.queue = s.queue[1:]
		return val
	}
	return -1
}

func (s *MyStack) Top() int {
	return s.queue[0]
}

func (s *MyStack) Empty() bool {
	return len(s.queue) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
