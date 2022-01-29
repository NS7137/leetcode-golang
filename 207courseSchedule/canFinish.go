package courseschedule

func CanFinish(numCourses int, prerequisites [][]int) bool {

	graph := buildGraph(numCourses, prerequisites)
	visited = make([]bool, numCourses)
	onPath = make([]bool, numCourses)
	hasCycle = false

	for i := 0; i < numCourses; i++ {
		// 遍历图所有节点
		traverse(graph, i)
	}
	// 只要没有环可完成所有课程
	return !hasCycle
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
	if visited[s] {
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
	// 节点s遍历完
	onPath[s] = false
}
