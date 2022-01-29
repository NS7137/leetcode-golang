package allpathsfromsourcetotarget

func AllPathsSourceTarget(graph [][]int) [][]int {
	var res [][]int
	// 递归过程中经过的路径
	var path []int
	traverse(graph, 0, &path, &res)
	return res
}

func traverse(graph [][]int, s int, path *[]int, res *[][]int) {
	// 添加节点到路径
	*path = append(*path, s)

	n := len(graph)
	if s == n-1 {
		var tmp []int
		// 到达终点
		tmp = append(tmp, *path...)
		*res = append(*res, tmp)
		*path = (*path)[:len(*path)-1]
		return
	}

	// 递归每个相邻节点
	for _, v := range graph[s] {
		traverse(graph, v, path, res)
	}

	// 从路径移除节点s
	*path = (*path)[:len(*path)-1]
}
