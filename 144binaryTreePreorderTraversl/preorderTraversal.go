package binarytreepreordertraversl

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 动态规划思路
func PreorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return nil
	}

	// 前序遍历结果特点，第一个是根节点值，接着是左子树，最后是右子树
	res = append(res, root.Val)
	res = append(res, PreorderTraversal(root.Left)...)
	res = append(res, PreorderTraversal(root.Right)...)
	return res
}

// 回溯思路
func PreorderTraversal2(root *TreeNode) []int {
	res := []int{}
	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		res = append(res, root.Val)
		traverse(root.Left)
		traverse(root.Right)
	}
	traverse(root)
	return res
}
