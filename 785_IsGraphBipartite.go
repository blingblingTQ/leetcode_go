package main

func dfs(graph [][]int, depth *[]int, i int) bool {
	for _, next := range graph[i] {
		if (*depth)[next] == -1 {
			(*depth)[next] = (*depth)[i] + 1
			if !dfs(graph, depth, next) {
				return false
			}
		} else {
			if (*depth)[next]%2 == (*depth)[i]%2 {
				return false
			}
		}
	}
	return true
}

func isBipartite(graph [][]int) bool {
	size := len(graph)
	if size <= 1 {
		return true
	}

	depth := make([]int, size)
	for i := 0; i < size; i++ {
		depth[i] = -1 //unvisited
	}

	for i := 0; i < size; i++ {
		if depth[i] == -1 {
			depth[i] = 0
			if !dfs(graph, &depth, i) {
				return false
			}
		}
	}
	return true
}
