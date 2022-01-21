package validatebinarysearchtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsValidBST(root *TreeNode) bool {
	return isValidHelper(root, nil, nil)
}

func isValidHelper(root *TreeNode, min *TreeNode, max *TreeNode) bool {
	// base case
	if root == nil {
		return true
	}
	// 若 root.val 不符合 max 和 min 限制，说明不是合法 BST
	if min != nil && root.Val <= min.Val {
		return false
	}
	if max != nil && root.Val >= max.Val {
		return false
	}

	// 限定左子树的最大值 是 root.val 右子树的最小值是 root.val
	return isValidHelper(root.Left, min, root) && isValidHelper(root.Right, root, max)
}
