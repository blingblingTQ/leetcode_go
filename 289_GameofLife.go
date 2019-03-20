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

func gameOfLife(board [][]int) {
	var row, col int
	row = len(board)
	if row > 0 {
		col = len(board[0])
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			//计算周围的存活细胞个数
			count := 0
			for t := max(0, i-1); t < min(i+2, row); t++ {
				for s := max(0, j-1); s < min(j+2, col); s++ {
					count += board[t][s] & 1
				}
			}
			count -= board[i][j]
			//下一轮存活，并将存活的信息编码至前一位
			if count == 3 || (count == 2 && board[i][j] == 1) {
				board[i][j] |= 2
			}
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			board[i][j] >>= 1
		}
	}
}
