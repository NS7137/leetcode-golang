package maximumsumbstinbinarytree

import "math"

/*
计算子树中BST的最大和
1 判断左右子树是否合法BST
2 判断左右子树加上自己是否合法BST  左子树最大值 小于 右子树最小值
3 BST所有节点值之和 和 别的BST比较
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 记录BST最大节点之和
var maxSum int

func MaxSumBST(root *TreeNode) int {
	maxSum = 0
	traverse(root)
	return maxSum
}

// 函数返回 int[]{isBST, min, max, sum}
func traverse(root *TreeNode) []int {
	// base case
	if root == nil {
		return []int{1, math.MaxInt64, math.MinInt64, 0}
	}

	left := traverse(root.Left)
	right := traverse(root.Right)

	// 后序遍历 通过left和right推导返回值，更新maxSum
	res := make([]int, 4)
	// 判断root为根的二叉树是否BST
	isBST := left[0] == 1 && right[0] == 1 && left[2] < root.Val && root.Val < right[1]
	if isBST {
		// root为根的二叉树是BST
		res[0] = 1
		// root为根的BST最小值
		res[1] = min(left[1], root.Val)
		// root 为根的BST最大值
		res[2] = max(right[2], root.Val)
		// root 为根的BST 所有节点和
		res[3] = left[3] + right[3] + root.Val
		// 更新 maxSum
		maxSum = max(maxSum, res[3])
	} else {
		res[0] = 0
	}

	return res
}

func min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
