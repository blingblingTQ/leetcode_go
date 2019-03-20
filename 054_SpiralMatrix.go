package main

func appendCycle(matrix [][]int, rows int, cols int, row int, col int) []int {
	startRow, startCol, endRow, endCol := row, col, rows-row-1, cols-col-1
	result := make([]int, 0)
	for i := startCol; i <= endCol; i++ {
		result = append(result, matrix[startRow][i])
	}
	if endRow > startRow {
		for i := startRow + 1; i <= endRow; i++ {
			result = append(result, matrix[i][endCol])
		}
	}
	if endRow > startRow && endCol > startCol {
		for i := endCol - 1; i >= startCol; i-- {
			result = append(result, matrix[endRow][i])
		}
	}
	if endCol > startCol && startRow < endRow-1 {
		for i := endRow - 1; i >= startRow+1; i-- {
			result = append(result, matrix[i][startCol])
		}
	}
	return result
}

func spiralOrder(matrix [][]int) []int {
	var rows, cols int
	rows = len(matrix)
	if rows > 0 {
		cols = len(matrix[0])
	}

	result := make([]int, 0)
	for i := 0; rows > i*2 && cols > i*2; i++ {
		result = append(result, appendCycle(matrix, rows, cols, i, i)...)
	}
	return result
}
