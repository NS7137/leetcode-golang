package intersectionoftwolinkedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	p1, p2 := headA, headB

	// 在这里第一轮体现在pA和pB第一次到达尾部会移向另一链表的表头, 而第二轮体现在如果pA或pB相交就返回交点, 不相交最后就是null==null
	for p1 != p2 {
		if p1 == nil {
			p1 = headB
		} else {
			p1 = p1.Next
		}

		if p2 == nil {
			p2 = headA
		} else {
			p2 = p2.Next
		}
	}

	return p1
}
