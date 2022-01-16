package removenthnodefromendoflist

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	// 找到倒数第n+1个节点
	x := findFromEnd(dummy, n+1)
	// 删倒数第n个节点
	x.Next = x.Next.Next
	return dummy.Next

}

// 返回链表的倒数第k个节点
func findFromEnd(head *ListNode, k int) *ListNode {
	p1 := head
	// p1 先走k步
	for i := 0; i < k; i++ {
		p1 = p1.Next
	}

	p2 := head
	// p1 和 p2 同时走 n-k 步
	for p1 != nil {
		p2 = p2.Next
		p1 = p1.Next
	}
	// p2 指向 n-k 个节点
	return p2
}
