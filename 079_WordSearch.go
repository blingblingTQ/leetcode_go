package main

func dfs(board [][]byte, row int, col int, word string, index int, visited [][]bool) bool {
	var rows, cols int
	rows = len(board)
	if rows > 0 {
		cols = len(board[0])
	}

	if index == len(word) {
		return true
	}
	if row < 0 || row >= rows {
		return false
	}
	if col < 0 || col >= cols {
		return false
	}
	if board[row][col] != byte(word[index]) {
		return false
	}
	if visited[row][col] {
		return false
	}

	visited[row][col] = true
	res := false
	res = res || dfs(board, row+1, col, word, index+1, visited) ||
		dfs(board, row-1, col, word, index+1, visited) ||
		dfs(board, row, col+1, word, index+1, visited) ||
		dfs(board, row, col-1, word, index+1, visited)
	visited[row][col] = false
	return res
}

func exist(board [][]byte, word string) bool {
	var rows, cols int
	rows = len(board)
	if rows > 0 {
		cols = len(board[0])
	}

	if len(word) == 0 {
		return false
	}

	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if dfs(board, i, j, word, 0, visited) {
				return true
			}
		}
	}
	return false
}
