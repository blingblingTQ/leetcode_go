package main

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

func (d *DSU) find(x int) int {
	if d.Parent[x] != x {
		d.Parent[x] = d.find(d.Parent[x])
	}
	return d.Parent[x]
}

func (d *DSU) union(x int, y int) {
	d.Parent[d.find(x)] = d.find(y)
}

func regionsBySlashes(grid []string) int {
	n := len(grid)
	dsu := constructor(4 * n * n)
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			root := 4 * (r*n + c)
			val := grid[r][c]
			if val != '\\' {
				dsu.union(root+0, root+1)
				dsu.union(root+2, root+3)
			}
			if val != '/' {
				dsu.union(root+0, root+2)
				dsu.union(root+1, root+3)
			}

			if r+1 < n {
				dsu.union(root+3, (root+4*n)+0)
			}
			if r-1 >= 0 {
				dsu.union(root+0, (root-4*n)+3)
			}
			if c+1 < n {
				dsu.union(root+2, (root+4)+1)
			}
			if c-1 >= 0 {
				dsu.union(root+1, (root-4)+2)
			}
		}
	}
	ans := 0
	for x := 0; x < 4*n*n; x++ {
		if dsu.find(x) == x {
			ans++
		}
	}
	return ans
}
