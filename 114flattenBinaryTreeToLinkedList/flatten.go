package flattenbinarytreetolinkedlist

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 将root的左子树和右子树flatten
// 将root的右子树接到左子树下方，然后将整个左子树作为右子树
func Flatten(root *TreeNode) {
	// base case
	if root == nil {
		return
	}

	Flatten(root.Left)
	Flatten(root.Right)

	// 后序
	// 左右子树已被拉平成链表
	left := root.Left
	right := root.Right

	// 将左子树作为右子树
	root.Left = nil
	root.Right = left

	// 将原右子树接到当前右子树的末端
	p := root
	for p.Right != nil {
		p = p.Right
	}
	p.Right = right
}
