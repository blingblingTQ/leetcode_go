package main

func dfs(graph map[int][]int, source int, target int, visited map[int]bool) bool {
	if _, ok := visited[source]; !ok {
		visited[source] = true
		if source == target {
			return true
		}
		for _, next := range graph[source] {
			if dfs(graph, next, target, visited) {
				return true
			}
		}
	}
	return false
}

func findRedundantConnection(edges [][]int) []int {
	graph := make(map[int][]int)
	for i := 0; i <= 1000; i++ {
		graph[i] = make([]int, 0)
	}

	for i := 0; i < len(edges); i++ {
		visited := make(map[int]bool)
		edge := edges[i]
		//dfs
		if len(graph[edge[0]]) > 0 && len(graph[edge[1]]) > 0 &&
			dfs(graph, edge[0], edge[1], visited) {
			return edge
		}
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}
	return nil
}

//union-find
type DSU struct {
	Parents []int
}

func constructor(n int) DSU {
	dsu := DSU{Parents: make([]int, n)}
	for i := 0; i < n; i++ {
		dsu.Parents[i] = i
	}
	return dsu
}

func (d *DSU) find(x int) int {
	if d.Parents[x] != x {
		d.Parents[x] = d.find(d.Parents[x])
	}
	return d.Parents[x]
}

func (d *DSU) union(x int, y int) {
	d.Parents[d.find(x)] = d.find(y)
}

func findRedundantConnection(edges [][]int) []int {
	dsu := constructor(1001)
	for _, edge := range edges {
		if dsu.find(edge[0]) == dsu.find(edge[1]) {
			return edge
		}
		dsu.union(edge[0], edge[1])
	}
	return nil
}
