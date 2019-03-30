package main

//solution one: DFS, to find cycle
func eventualSafeNodes(graph [][]int) []int {
	size := len(graph)
	state := make([]int, 0)
	//0: unvisited, 1: safe, 2: unsafe
	color := make([]int, size)
	for i := 0; i < size; i++ {
		if dfs(graph, i, &color) {
			state = append(state, i)
		}
	}
	return state
}

func dfs(graph [][]int, i int, color *[]int) bool {
	//如果已经访问过，则返回之前的状态
	if (*color)[i] != 0 {
		return (*color)[i] == 1
	}

	//先标记成unsafe
	(*color)[i] = 2
	for _, next := range graph[i] {
		if !dfs(graph, next, color) {
			return false
		}
	}
	(*color)[i] = 1
	return true
}
