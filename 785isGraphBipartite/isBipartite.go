package isgraphbipartite

func IsBipartite(graph [][]int) bool {
	n := len(graph)
	ok = true
	color = make([]bool, n)
	visited = make([]bool, n)

	// 图不一定是联通的，可能存在多个子图
	// 把每个节点都作为起点进行遍历
	// 任何一个子图不是二分，则都不算二分
	for v := 0; v < n; v++ {
		if !visited[v] {
			traverse(graph, v)
		}
	}
	return ok
}

// 是否符合二分图性质
var ok bool

// 节点颜色
var color []bool

// 是否被访问
var visited []bool

// DFS 遍历
func traverse(graph [][]int, v int) {
	// 是否二分
	if !ok {
		return
	}
	// 标记访问
	visited[v] = true

	// 遍历邻接是否被访问
	for _, w := range graph[v] {
		if !visited[w] {
			// 相邻没有被访问
			// 给节点w涂上和节点v不同颜色
			color[w] = !color[v]
			// 继续遍历 w
			traverse(graph, w)
		} else {
			// w 已经被访问
			// 根据 v和w的颜色判断是否二分
			if color[w] == color[v] {
				// 相同则不是二分
				ok = false
			}
		}
	}
}

// BFS 遍历
func bfs(graph [][]int, start int) {
	q := []int{}
	visited[start] = true
	q = append(q, start)

	for len(q) != 0 && ok {
		v := q[0]
		q = q[1:]

		for _, w := range graph[v] {
			if !visited[w] {
				color[w] = !color[v]
				visited[w] = true
				q = append(q, w)
			} else {
				if color[w] == color[v] {
					ok = false
				}
			}
		}
	}
}
