package utils

func HasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		// 快慢相遇则含有环
		if slow == fast {
			return true
		}
	}
	return false
}

// 计算环起点
func DetectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		// 有环
		if slow == fast {
			break
		}
	}

	// fast遇空指针说明没有环
	if fast == nil || fast.Next == nil {
		return nil
	}

	// 重指向head
	slow = head
	// 快慢同步，相交则环起点
	for slow != fast {
		fast = fast.Next
		slow = slow.Next
	}

	return slow
}
