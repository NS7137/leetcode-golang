package serializeanddeserializebinarytree

import "fmt"

func Print(root *TreeNode) {
	if root == nil {
		return
	}
	// 前
	fmt.Print(root.Val)
	Print(root.Left)
	// 中
	Print(root.Right)
	// 后
}

func (c *Codec) LevelPrint(root *TreeNode) {
	if root == nil {
		return
	}
	c.nq = append(c.nq, root)
	for len(c.nq) != 0 {
		cur := c.nq[0]
		c.nq = c.nq[1:]

		fmt.Print(cur.Val)

		if cur.Left != nil {
			c.nq = append(c.nq, cur.Left)
		}

		if cur.Right != nil {
			c.nq = append(c.nq, cur.Right)
		}
	}
}
