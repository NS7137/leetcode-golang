package diameterofbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var maxDiameter int

func DiameterOfBinaryTree(root *TreeNode) int {
	maxDiameter = 0
	maxDepthByPostorder(root)
	return maxDiameter

}

// 每条二叉树直径长度，就是节点左右子树的最大深度和
// 前序位置无法获取子树信息，所以只能让每个节点调用maxDepth算子树深度
func traverse(root *TreeNode) {
	if root == nil {
		return
	}
	// 对每个节点计算直径 又调用递归，最坏时间复杂度O(N^2)
	leftMax := maxDepth(root.Left)
	rightMax := maxDepth(root.Right)
	diameter := leftMax + rightMax
	// 更新最大直径
	maxDiameter = max(maxDiameter, diameter)

	traverse(root.Left)
	traverse(root.Right)
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftMax := maxDepth(root.Left)
	rightMax := maxDepth(root.Right)
	return 1 + max(leftMax, rightMax)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 把计算逻辑放在后序位置，知道左右子树最大深度  O(N)
func maxDepthByPostorder(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftMax := maxDepthByPostorder(root.Left)
	rightMax := maxDepthByPostorder(root.Right)
	// 后序位置计算直径
	diameter := leftMax + rightMax
	maxDiameter = max(maxDiameter, diameter)
	return 1 + max(leftMax, rightMax)
}
