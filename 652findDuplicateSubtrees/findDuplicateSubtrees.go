package findduplicatesubtrees

import "fmt"

/*
二叉树(子树)长啥样
其他节点为根的子树长啥样
每个节点子树序列化
hashmap记录子树 及出现次数
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func FindDuplicateSubtrees(root *TreeNode) []*TreeNode {
	// 记录所有子树
	memo := make(map[string]int)
	// 记录重复的子树根节点
	res := []*TreeNode{}
	if root == nil {
		return res
	}
	traverse(root, memo, &res)
	return res
}

func traverse(root *TreeNode, memo map[string]int, res *[]*TreeNode) string {
	if root == nil {
		return "#"
	}

	left := traverse(root.Left, memo, res)
	right := traverse(root.Right, memo, res)
	// 序列化
	subTree := fmt.Sprintf("%v,%v,%v", left, right, root.Val)
	memo[subTree]++
	// 只有重复才入队列
	if memo[subTree] == 2 {
		*res = append(*res, root)
	}
	return subTree
}
