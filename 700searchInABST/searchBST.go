package searchinabst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SearchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	// 小于 root.val 在左树找
	if val < root.Val {
		return SearchBST(root.Left, val)
	}

	// 大于 root.val 在右树找
	if val > root.Val {
		return SearchBST(root.Right, val)
	}

	return root
}

// 普通二叉树穷举所有节点
func SearchBT(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	// 当前没找到就递归左右子树
	left := SearchBT(root.Left, val)
	right := SearchBT(root.Right, val)
	if left != nil {
		return left
	} else {
		return right
	}
}
