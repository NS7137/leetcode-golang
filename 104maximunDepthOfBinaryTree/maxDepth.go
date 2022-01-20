package maximundepthofbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 最大深度
var res int

// 遍历到的节点深度
var depth int

func MaxDepth(root *TreeNode) int {
	res = 0
	depth = 0
	traverse(root)
	return res
}

// 通过遍历返回
func traverse(root *TreeNode) {
	if root == nil {
		res = max(res, depth)
		return
	}
	// 前序位置 进入节点+1
	depth++
	traverse(root.Left)
	traverse(root.Right)
	// 后序位置 离开节点-1
	depth--
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 通过子树最大高度推导
func maxDepthByDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftMax := maxDepthByDepth(root.Left)
	rightMax := maxDepthByDepth(root.Right)
	// 最大深度等于左右子树的最大值 + 根节点本身
	res = max(leftMax, rightMax) + 1
	return res
}
