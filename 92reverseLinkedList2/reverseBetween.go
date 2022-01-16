package reverselinkedlist2

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	// base case
	if left == 1 {
		return reverseN(head, right)
	}
	// 前进到反转起点 触发 base case
	head.Next = ReverseBetween(head.Next, left-1, right-1)
	return head
}

// 后驱
var successor *ListNode

func reverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		// n + 1
		successor = head.Next
		return head
	}

	last := reverseN(head.Next, n-1)
	head.Next.Next = head
	head.Next = successor
	return last
}
