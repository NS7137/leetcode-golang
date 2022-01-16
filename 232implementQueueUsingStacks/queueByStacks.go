package implementqueueusingstacks

type MyQueue struct {
	inStack  []int
	outStack []int
}

func Constructor() MyQueue {
	in := []int{}
	out := []int{}
	q := MyQueue{
		inStack:  in,
		outStack: out,
	}
	return q
}

func (q *MyQueue) Push(x int) {
	q.inStack = append(q.inStack, x)
}

func (q *MyQueue) Pop() int {
	val := q.Peek()
	q.outStack = q.outStack[:len(q.outStack)-1]
	return val
}

func (q *MyQueue) Peek() int {
	if len(q.outStack) == 0 {
		for len(q.inStack) != 0 {
			q.outStack = append(q.outStack, q.inStack[len(q.inStack)-1])
			q.inStack = q.inStack[:len(q.inStack)-1]
		}
	}
	return q.outStack[len(q.outStack)-1]
}

func (q *MyQueue) Empty() bool {
	return len(q.inStack) == 0 && len(q.outStack) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
