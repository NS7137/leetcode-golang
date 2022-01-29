package coursescheduleii

func FindOrder(numCourses int, prerequisites [][]int) []int {
	graph := buildGraph(numCourses, prerequisites)
	visited = make([]bool, numCourses)
	onPath = make([]bool, numCourses)
	hasCycle = false
	postorder = make([]int, 0)

	for i := 0; i < numCourses; i++ {
		traverse(graph, i)
	}

	// 有环图无法进行拓扑排序
	if hasCycle {
		return []int{}
	}

	// 逆 后序遍历结果 即为拓扑排序结果
	res := reverse(postorder)
	return res

}

// 建图函数
func buildGraph(numCourses int, prerequisites [][]int) [][]int {
	// 图中共有 numCourses 个节点
	graph := make([][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		graph[i] = []int{}
	}

	for _, edge := range prerequisites {
		from := edge[1]
		to := edge[0]
		// 修完from才能修to
		// 图中添加一条从from指向to的有向边
		graph[from] = append(graph[from], to)
	}
	return graph
}

// 记录后序遍历结果
var postorder []int

// 一次travers 递归经过的节点
var onPath []bool

// 是否有环
var hasCycle bool

// 防止重复遍历
var visited []bool

// 从节点s开始DFS遍历，将遍历过的节点标记为true
func traverse(graph [][]int, s int) {
	if onPath[s] {
		// 发现环
		hasCycle = true
	}
	if visited[s] || hasCycle {
		return
	}
	// 前序遍历代码
	// 将当前节点标记为已遍历
	visited[s] = true
	// 开始遍历节点s
	onPath[s] = true
	for _, t := range graph[s] {
		traverse(graph, t)
	}
	// 后序遍历代码
	postorder = append(postorder, s)
	// 节点s遍历完
	onPath[s] = false
}

func reverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
