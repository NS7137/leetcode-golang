package maximumbinarytree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func ConstructMaximumBinaryTree(nums []int) *TreeNode {
	return build(nums, 0, len(nums)-1)
}

func build(nums []int, lo int, hi int) *TreeNode {
	// base case
	if lo > hi {
		return nil
	}

	// 找最大值和对应索引
	index := -1
	maxVal := math.MinInt64
	for i := lo; i <= hi; i++ {
		if maxVal < nums[i] {
			index = i
			maxVal = nums[i]
		}
	}

	// 递归构造左右子树
	root := &TreeNode{}
	root.Val = maxVal
	root.Left = build(nums, lo, index-1)
	root.Right = build(nums, index+1, hi)
	return root
}
