package main

//union-find
type DSU struct {
	Parent []int
}

func constructor(n int) DSU {
	dsu := DSU{Parent: make([]int, n)}
	for i := 0; i < n; i++ {
		dsu.Parent[i] = i
	}
	return dsu
}

func (d *DSU) Find(x int) int {
	if d.Parent[x] != x {
		d.Parent[x] = d.Find(d.Parent[x])
	}
	return d.Parent[x]
}

func (d *DSU) Union(x int, y int) {
	d.Parent[d.Find(x)] = d.Find(y)
}

func findCircleNum(M [][]int) int {
	n := len(M)
	dsu := constructor(n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if M[i][j] == 1 {
				dsu.Union(i, j)
			}
		}
	}

	count := 0
	for i := 0; i < n; i++ {
		if dsu.Find(i) == i {
			count++
		}
	}
	return count
}

//dfs
func dfs(M [][]int, i int, j int, visited *[][]bool) {
	n := len(M)
	if (*visited)[i][j] {
		return
	}
	(*visited)[i][j] = true
	for k := 0; k < n; k++ {
		if M[j][k] == 1 {
			dfs(M, j, k, visited)
		}
	}
}

func findCircleNum(M [][]int) int {
	n := len(M)

	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, n)
	}

	count := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if M[i][j] == 1 && !visited[i][j] {
				dfs(M, i, j, &visited)
				count++
			}
		}
	}
	return count
}
