package uniquebsts2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return build(1, n)

}

func build(lo, hi int) (res []*TreeNode) {
	if lo > hi {
		res = append(res, nil)
		return
	}
	// 穷举root 节点
	for i := lo; i <= hi; i++ {
		// 递归构造左右子树的所有合法bst
		leftTree := build(lo, i-1)
		rightTree := build(i+1, hi)
		// root 节点穷举所有左右子树组合
		for _, left := range leftTree {
			for _, right := range rightTree {
				root := &TreeNode{Val: i}
				root.Left = left
				root.Right = right
				res = append(res, root)
			}
		}
	}
	return
}
