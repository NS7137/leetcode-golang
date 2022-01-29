package connectingcitieswithminimumcost

import (
	"leetcode-golang/utils"
	"sort"
)

/*
最小生成数 mst
1.包含图中所有节点
2.形成的结构是树，即不存在环
3.权重和最小
1,2由union-find处理
3 用到贪心思路
权重从小到大排序，从最小开始遍历，是否会和mst集合中的其他边形成环
不会加入mst集合，会则略过
最后mst集合中的边就形成了最小生成树
*/

// kruskal
func MinimumCost(n int, connections [][]int) int {
	// 城市编号 1...n
	uf := utils.NewUF(n + 1)
	// 对所有边从小到大排序
	sort.Sort(ByWeight(connections))
	// 记录权重之和
	mst := 0
	for _, edge := range connections {
		u := edge[0]
		v := edge[1]
		weight := edge[2]
		// 若产生环，则不能加入mst
		if uf.Connected(u, v) {
			continue
		}
		mst += weight
		uf.Union(u, v)
	}
	// 保证所有节点都被连通 0未被使用所以连通量为2
	if uf.Count() == 2 {
		return mst
	} else {
		return -1
	}
}

// 按权重升序排序
type ByWeight [][]int

func (a ByWeight) Len() int {
	return len(a)
}
func (a ByWeight) Less(i, j int) bool {
	return a[i][2] < a[j][2]
}
func (a ByWeight) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
