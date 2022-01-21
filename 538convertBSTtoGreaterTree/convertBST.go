package convertbsttogreatertree

// 同 1038. 把二叉搜索树转换为累加树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var sum int

func ConvertBST(root *TreeNode) *TreeNode {
	sum = 0
	traverse(root)
	return root
}

// 降序遍历 记录累加
func traverse(root *TreeNode) {
	if root == nil {
		return
	}

	traverse(root.Right)
	sum += root.Val
	root.Val = sum
	traverse(root.Left)
}
