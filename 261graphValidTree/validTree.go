package graphvalidtree

import "leetcode-golang/utils"

/*
什么情况下加入一条边会使得树变成图，出现环
如果该边的两个节点本来就在同一个连通分量里，那么添加这条边就会产生环
反之，如果该边的两个节点不在同一个连通分量里，则不产生环
判断是否连通，就可以用union-find
*/

// n个节点  edges无向边列表
func ValidTree(n int, edges [][]int) bool {

	uf := utils.NewUF(n)

	// 遍历所有边，将组成边的两个节点连接
	for _, edge := range edges {
		u := edge[0]
		v := edge[1]
		// 若两个节点已经在同一连通分量中，会产生环
		if uf.Connected(u, v) {
			return false
		}
		uf.Union(u, v)
	}

	// 保证最后只形成一棵树，即只有一个连通分量
	return uf.Count() == 1
}
