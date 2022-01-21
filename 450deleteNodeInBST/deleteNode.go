package deletenodeinbst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
三种情况 不能破坏BST性质
1 两个子节点为空，直接删除
2 两个子节点一个非空，让这个子节点接替自己位置
3 两个子节点都非空，找左子树最大，或右子树中最小的节点接替自己位置
*/
func DeleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	// 1,2 两种情况
	if root.Val == key {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}

		// 情况3
		// 获得右子树最小的节点
		minNode := getMin(root.Right)
		// 删除右树最小的节点
		root.Right = DeleteNode(root.Right, minNode.Val)
		// 一般不通过修改节点内部值来交换节点，操作应该和内部存储的数据域解耦，没必要关心内部数据
		// 用右树最小的节点替换root节点
		minNode.Left = root.Left
		minNode.Right = root.Right
		root = minNode
	} else if root.Val > key {
		root.Left = DeleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = DeleteNode(root.Right, key)
	}
	return root
}

func getMin(node *TreeNode) *TreeNode {
	// bst 最左边的就是最小的
	for node.Left != nil {
		node = node.Left
	}
	return node
}
