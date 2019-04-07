package main

func findJudge(N int, trust [][]int) int {
	in := make(map[int][]int)
	out := make(map[int][]int)
	for _, edge := range trust {
		in[edge[1]] = append(in[edge[1]], edge[0])
		out[edge[0]] = append(out[edge[0]], edge[1])
	}

	for i := 1; i <= N; i++ {
		if len(in[i]) == N-1 && len(out[i]) == 0 {
			return i
		}
	}
	return -1
}
