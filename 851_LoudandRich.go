package main

func dfs(person int, graph map[int][]int, quiet []int, ans *[]int) int {
	if (*ans)[person] == -1 {
		(*ans)[person] = person
		for _, nextPerson := range graph[person] {
			index := dfs(nextPerson, graph, quiet, ans)
			if quiet[index] < quiet[(*ans)[person]] {
				(*ans)[person] = index
			}
		}
	}
	return (*ans)[person]
}

func loudAndRich(richer [][]int, quiet []int) []int {
	graph := make(map[int][]int)
	for i := 0; i < len(quiet); i++ {
		graph[i] = make([]int, 0)
	}
	for _, r := range richer {
		graph[r[1]] = append(graph[r[1]], r[0])
	}

	res := make([]int, len(quiet))
	for i := 0; i < len(res); i++ {
		res[i] = -1
	}
	for i := 0; i < len(res); i++ {
		dfs(i, graph, quiet, &res)
	}
	return res
}
