package middleofthelinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func MiddleNode(head *ListNode) *ListNode {

	// 快慢指针 指向head
	slow, fast := head, head
	// 快走2步，慢走1步
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 慢指向中点，如链表长度偶数，返回的是靠后的那个节点
	return slow

}
