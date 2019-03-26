package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfs(grid [][]int, rows int, cols int, i int, j int, visited *[][]bool) int {
	if i < 0 || i >= rows || j < 0 || j >= cols {
		return 0
	}
	if grid[i][j] == 0 || (*visited)[i][j] {
		return 0
	}
	(*visited)[i][j] = true
	return dfs(grid, rows, cols, i-1, j, visited) +
		dfs(grid, rows, cols, i+1, j, visited) +
		dfs(grid, rows, cols, i, j-1, visited) +
		dfs(grid, rows, cols, i, j+1, visited) + 1
}

func maxAreaOfIsland(grid [][]int) int {
	var rows, cols int
	rows = len(grid)
	if rows > 0 {
		cols = len(grid[0])
	}

	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}

	maxArea := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				maxArea = max(maxArea, dfs(grid, rows, cols, i, j, &visited))
			}
		}
	}
	return maxArea
}
