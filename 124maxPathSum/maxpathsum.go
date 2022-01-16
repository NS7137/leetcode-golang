package maxpathsum

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MaxPathSum(root *TreeNode) int {
	ret := math.MinInt32
	helper(root, &ret)
	return ret
}

func helper(root *TreeNode, ret *int) int {
	if root == nil {
		return 0
	}

	l := helper(root.Left, ret)
	r := helper(root.Right, ret)
	sum := root.Val
	if l > 0 {
		sum += l
	}
	if r > 0 {
		sum += r
	}

	*ret = max(*ret, sum)
	if max(l, r) > 0 {
		return max(l, r) + root.Val
	}
	return root.Val
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
