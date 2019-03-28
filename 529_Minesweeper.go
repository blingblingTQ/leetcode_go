package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func idDigit(c byte) bool {
	if c-'0' > 0 && c-'0' <= 8 {
		return true
	}
	return false
}

func getMineCount(board [][]byte, rows int, cols int, row int, col int) int {
	count := 0
	for i := max(0, row-1); i <= min(row+1, rows-1); i++ {
		for j := max(0, col-1); j <= min(col+1, cols-1); j++ {
			if board[i][j] == 'M' {
				count++
			}
		}
	}
	return count
}

func dfs(board *[][]byte, rows int, cols int, row int, col int) {
	if row < 0 || row >= rows || col < 0 || col >= cols {
		return
	}
	if (*board)[row][col] == 'M' {
		(*board)[row][col] = 'X'
		return
	} else if (*board)[row][col] == 'B' {
		return
	} else {
		if count := getMineCount(*board, rows, cols, row, col); count > 0 {
			(*board)[row][col] = byte(count + '0')
		} else {
			(*board)[row][col] = 'B'
			dfs(board, rows, cols, row+1, col+1)
			dfs(board, rows, cols, row+1, col)
			dfs(board, rows, cols, row+1, col-1)
			dfs(board, rows, cols, row, col+1)
			dfs(board, rows, cols, row, col-1)
			dfs(board, rows, cols, row-1, col+1)
			dfs(board, rows, cols, row-1, col)
			dfs(board, rows, cols, row-1, col-1)
		}
	}
}

func updateBoard(board [][]byte, click []int) [][]byte {
	rows, cols, row, col := len(board), len(board[0]), click[0], click[1]
	dfs(&board, rows, cols, row, col)
	return board
}
