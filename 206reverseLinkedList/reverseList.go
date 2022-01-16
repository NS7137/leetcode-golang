package reverselinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	last := ReverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

var successor *ListNode

func ReverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		// 记录第 n+1 个节点
		successor = head.Next
		return head
	}
	last := ReverseN(head.Next, n-1)
	head.Next.Next = head
	head.Next = successor
	return last
}
