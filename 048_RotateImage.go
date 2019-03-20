package main

func rotateCore(matrix [][]int, start int) {
	n := len(matrix)
	endX := n - start - 1
	endY := endX
	for i := 0; i <= endX-start-1; i++ {
		tmp := matrix[start][start+i]
		matrix[start][start+i] = matrix[endY-i][start]
		matrix[endY-i][start] = matrix[endX][endY-i]
		matrix[endX][endY-i] = matrix[start+i][endY]
		matrix[start+i][endY] = tmp
	}
}

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		rotateCore(matrix, i)
	}
}
