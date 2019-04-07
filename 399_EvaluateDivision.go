package main

type Edge struct {
	To    string
	Score float64
}

func dfs(graph map[string][]Edge, source string, target string, visited map[string]struct{}) float64 {
	if _, ok := graph[source]; !ok {
		return -1.0
	}
	if _, ok := graph[target]; !ok {
		return -1.0
	}
	if source == target {
		return 1.0
	}

	if _, ok := visited[source]; ok {
		return -1.0
	}
	visited[source] = struct{}{}

	for _, edge := range graph[source] {
		if res := dfs(graph, edge.To, target, visited); res != -1.0 {
			return res * edge.Score
		}
	}
	return -1.0
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := make(map[string][]Edge)
	for i := 0; i < len(values); i++ {
		graph[equations[i][0]] = append(graph[equations[i][0]], Edge{
			To:    equations[i][1],
			Score: values[i],
		})
		graph[equations[i][1]] = append(graph[equations[i][1]], Edge{
			To:    equations[i][0],
			Score: float64(1.0 / values[i]),
		})
	}

	result := make([]float64, len(queries))
	for i := 0; i < len(queries); i++ {
		visited := make(map[string]struct{})
		result[i] = dfs(graph, queries[i][0], queries[i][1], visited)
	}
	return result
}
