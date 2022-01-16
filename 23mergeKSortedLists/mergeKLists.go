package mergeksortedlists

import (
	"container/heap"
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeKLists(lists []*ListNode) *ListNode {
	n := []int{}
	if len(lists) == 0 {
		return nil
	}
	// 把lists中val全遍历出来放入n
	for _, list := range lists {
		tmp := list
		for tmp != nil {
			n = append(n, tmp.Val)
			tmp = tmp.Next
		}
	}
	if len(n) == 0 {
		return nil
	}
	// 排序n
	sort.Ints(n)
	start := &ListNode{Val: n[0]}
	ptr := start
	// 按序链接
	for i := 0; i < len(n); i++ {
		ptr.Next = &ListNode{Val: n[i]}
		ptr = ptr.Next
	}
	return start.Next
}

type NodeHeap []*ListNode

func (h NodeHeap) Len() int           { return len(h) }
func (h NodeHeap) Less(i, j int) bool { return h[i].Val <= h[j].Val }
func (h NodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *NodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func MergeKListsByHeap(lists []*ListNode) *ListNode {
	mergeHeap := &NodeHeap{}
	heap.Init(mergeHeap)

	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			heap.Push(mergeHeap, lists[i])
		}
	}

	var head *ListNode
	var tail *ListNode
	var next *ListNode

	for mergeHeap.Len() > 0 {
		if tail != nil {
			next = heap.Pop(mergeHeap).(*ListNode)
			if next.Next != nil {
				heap.Push(mergeHeap, next.Next)
			}
			tail.Next = next
			tail = next
		} else {
			head = heap.Pop(mergeHeap).(*ListNode)
			tail = head
			if head.Next != nil {
				heap.Push(mergeHeap, head.Next)
			}
		}
	}

	return head
}
