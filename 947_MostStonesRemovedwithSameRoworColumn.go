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

func removeStones(stones [][]int) int {
	n := len(stones)
	dsu := constructor(20000)

	for _, stone := range stones {
		dsu.union(stone[0], stone[1]+10000)
	}

	seen := make(map[int]bool)

	for _, stone := range stones {
		seen[dsu.find(stone[0])] = true
	}
	return n - len(seen)
}
