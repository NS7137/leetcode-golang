package constructbinaraytreefrominorderandpostordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildTree(inorder []int, postorder []int) *TreeNode {
	return build(inorder, 0, len(inorder)-1,
		postorder, 0, len(postorder)-1)
}

func build(inorder []int, inStart int, inEnd int,
	postorder []int, postStart int, postEnd int) *TreeNode {

	if inStart > inEnd {
		return nil
	}
	// 从后序遍历最后元素 为root节点
	rootVal := postorder[postEnd]
	// 在中序遍历中找 root节点 索引
	index := 0
	for i := inStart; i <= inEnd; i++ {
		if inorder[i] == rootVal {
			index = i
			break
		}
	}

	root := &TreeNode{}
	root.Val = rootVal

	leftSize := index - inStart
	// 递归构造左右子树
	root.Left = build(inorder, inStart, index-1,
		postorder, postStart, postStart+leftSize-1)
	root.Right = build(inorder, index+1, inEnd,
		postorder, postStart+leftSize, postEnd-1)
	return root
}
