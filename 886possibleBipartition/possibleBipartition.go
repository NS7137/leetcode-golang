package possiblebipartition

var ok bool
var color []bool
var visited []bool

func PossibleBipartition(n int, dislikes [][]int) bool {
	ok = true
	color = make([]bool, n+1)
	visited = make([]bool, n+1)
	// 转邻接表
	graph := buildGraph(n, dislikes)

	for v := 1; v <= n; v++ {
		if !visited[v] {
			traverse(graph, v)
		}
	}
	return ok
}

// 建图
func buildGraph(n int, dislikes [][]int) [][]int {
	// 编号 1...n
	graph := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		graph[i] = []int{}
	}

	for _, edge := range dislikes {
		v := edge[0]
		w := edge[1]
		// 无向图
		graph[v] = append(graph[v], w)
		graph[w] = append(graph[w], v)
	}
	return graph
}

// dfs 双色问题
func traverse(graph [][]int, v int) {
	if !ok {
		return
	}
	visited[v] = true

	for _, w := range graph[v] {
		if !visited[w] {
			color[w] = !color[v]
			traverse(graph, w)
		} else {
			if color[w] == color[v] {
				ok = false
			}
		}
	}
}

// BFS 遍历着色 与 785一样
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
