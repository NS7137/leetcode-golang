package invertbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InvertTree(root *TreeNode) *TreeNode {
	// base case
	if root == nil {
		return nil
	}

	//前序 root节点需要交换左右子节点
	root.Left, root.Right = root.Right, root.Left

	// 递归
	InvertTree(root.Left)
	InvertTree(root.Right)

	return root

}
