package utils

// 返回链表的倒数第k个节点
func FindFromEnd(head *ListNode, k int) *ListNode {
	fast := head
	// fast 先走k步
	for i := 0; i < k; i++ {
		fast = fast.Next
	}

	slow := head
	// fast 和 slow 同时走 n-k 步
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	// slow 指向 n-k 个节点
	return slow
}
