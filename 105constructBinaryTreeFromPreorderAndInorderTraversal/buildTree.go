package constructbinarytreefrompreorderandinordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildTree(preorder []int, inorder []int) *TreeNode {
	return build(preorder, 0, len(preorder)-1,
		inorder, 0, len(inorder)-1)
}

func build(preorder []int, preStart int, preEnd int,
	inorder []int, inStart int, inEnd int) *TreeNode {

	if preStart > preEnd {
		return nil
	}
	// root 节点对应的值就是前序遍历数组的第一个元素
	rootVal := preorder[preStart]
	// 找 rootVal 在中序遍历数组的索引
	index := 0
	for i := inStart; i <= inEnd; i++ {
		if inorder[i] == rootVal {
			index = i
			break
		}
	}
	root := &TreeNode{}
	root.Val = rootVal
	// 左子树的长度
	leftSize := index - inStart
	// 递归构造左右子树
	root.Left = build(preorder, preStart+1, preStart+leftSize,
		inorder, inStart, index-1)
	root.Right = build(preorder, preStart+leftSize+1, preEnd,
		inorder, index+1, inEnd)

	return root
}
