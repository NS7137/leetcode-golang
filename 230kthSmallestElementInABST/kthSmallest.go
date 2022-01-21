package kthsmallestelementinabst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 记录结果
var res int

// 记录当前元素排名
var rank int

func KthSmallest(root *TreeNode, k int) int {
	res = 0
	rank = 0
	traverse(root, k)
	return res
}

// 中序遍历，BST 升序排序
func traverse(root *TreeNode, k int) {
	if root == nil {
		return
	}
	traverse(root.Left, k)
	rank++
	if k == rank {
		res = root.Val
		return
	}
	traverse(root.Right, k)
}

/*
O(logN) 每个节点知道自己的排名 m ， 比较 k 和 m大小
如果 m = k，则 找到第k个元素
如果 k < m，则 第k元素 在左子树，所以去左子树搜索第k个元素
如果 k > m，则 第k元素 在右子树，所以去右子树搜索第k-m-1个元素

节点中可以维护 自己为根的二叉树有多少个节点
*/

func Kth(root *TreeNode, k int) int {
	leftSum := sumTree(root.Left)
	if leftSum+1 == k {
		return root.Val
	} else if leftSum+1 > k {
		return Kth(root.Left, k)
	} else if leftSum+1 < k {
		return Kth(root.Right, k-leftSum-1)
	}
	return root.Val
}

func sumTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + sumTree(root.Left) + sumTree(root.Right)
}
