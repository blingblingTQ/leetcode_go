package main

import (
	"fmt"
	"math"
)

type Edge struct {
	To    int
	Delay int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func networkDelayTime(times [][]int, N int, K int) int {
	graph := make(map[int][]Edge)
	for _, time := range times {
		graph[time[0]] = append(graph[time[0]], Edge{
			To:    time[1],
			Delay: time[2],
		})
	}

	visited := make(map[int]int)
	for i := 1; i <= N; i++ {
		visited[i] = math.MaxInt32
	}
	dfs(graph, K, 0, visited)
	time := 0
	for i := 1; i <= N; i++ {
		if i == K {
			continue
		}
		if visited[i] == math.MaxInt32 {
			return -1
		}
		time = max(time, visited[i])
	}
	return time
}

func dfs(graph map[int][]Edge, node int, elapsed int, visited map[int]int) {
	if elapsed >= visited[node] {
		return
	}

	visited[node] = elapsed
	for _, edge := range graph[node] {
		dfs(graph, edge.To, elapsed+edge.Delay, visited)
	}
}

func main() {
	times := [][]int{[]int{2, 1, 1}, []int{2, 3, 1}, []int{3, 4, 1}}
	N := 4
	K := 2
	fmt.Println(networkDelayTime(times, N, K))
}
