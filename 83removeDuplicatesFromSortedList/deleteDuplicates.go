package removeduplicatesfromsortedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func DeleteDuplicates(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	slow, fast := head, head
	for fast != nil {
		if fast.Val != slow.Val {
			slow.Next = fast
			slow = slow.Next
		}
		fast = fast.Next
	}
	// 断开与后面重复元素的连接
	slow.Next = nil
	return head

}
