package mergetwosortedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	head := &ListNode{Val: 0, Next: nil}
	// 保存原始头
	ptr := head
	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			head.Next = list2
			list2 = list2.Next
		} else {
			head.Next = list1
			list1 = list1.Next
		}
		// 头指针往后移动
		head = head.Next
	}

	if list1 != nil {
		head.Next = list1
	}

	if list2 != nil {
		head.Next = list2
	}

	return ptr.Next
}

func MergeTwoListsByRecursion(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else if list1.Val < list2.Val {
		list1.Next = MergeTwoListsByRecursion(list1.Next, list2)
		return list1
	} else {
		list2.Next = MergeTwoListsByRecursion(list1, list2.Next)
		return list2
	}
}
