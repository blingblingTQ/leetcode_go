package main

func dfs(grid [][]byte, rows int, cols int, i int, j int) {
	if i < 0 || i >= rows || j < 0 || j >= cols {
		return
	}

	if grid[i][j] == '1' {
		grid[i][j] = '2'
		dfs(grid, rows, cols, i-1, j)
		dfs(grid, rows, cols, i+1, j)
		dfs(grid, rows, cols, i, j-1)
		dfs(grid, rows, cols, i, j+1)
	}
}

func numIslands(grid [][]byte) int {
	var rows, cols int
	rows = len(grid)
	if rows > 0 {
		cols = len(grid[0])
	}

	count := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				dfs(grid, rows, cols, i, j)
				count++
			}
		}
	}
	return count
}

type DSU struct {
	Parents []int
}

func Constructor(n int) DSU {
	dsu := DSU{Parents: make([]int, n)}
	for i := 0; i < n; i++ {
		dsu.Parents[i] = i
	}
	return dsu
}

func (d *DSU) Find(x int) int {
	if d.Parents[x] != x {
		d.Parents[x] = d.Find(d.Parents[x])
	}
	return d.Parents[x]
}

func (d *DSU) Union(x int, y int) {
	d.Parents[d.Find(x)] = d.Find(y)
}

func numIslands(grid [][]byte) int {
	var rows, cols int
	rows = len(grid)
	if rows > 0 {
		cols = len(grid[0])
	}

	dsu := Constructor(rows * cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				x := i*cols + j
				if i+1 < rows && grid[i+1][j] == '1' {
					y := (i+1)*cols + j
					dsu.Union(x, y)
				}
				if j+1 < cols && grid[i][j+1] == '1' {
					y := i*cols + j + 1
					dsu.Union(x, y)
				}
			}
		}
	}

	island := make(map[int]bool)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				island[dsu.Find(i*cols+j)] = true
			}
		}
	}
	return len(island)
}

func main() {
	grid := [][]byte{
		[]byte{'1', '1', '1', '1', '0'},
		[]byte{'1', '1', '0', '1', '0'},
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'0', '0', '0', '0', '0'},
	}

	fmt.Println(numIslands(grid))
}
