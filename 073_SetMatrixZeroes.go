package main

//方法一：O(m+n)的空间复杂度
func setZeroes(matrix [][]int) {
	var m, n int
	m = len(matrix)
	if m > 0 {
		n = len(matrix[0])
	}
	setZeroM := make([]bool, m)
	setZeroN := make([]bool, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				setZeroM[i] = true
				setZeroN[j] = true
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if setZeroM[i] || setZeroN[j] {
				matrix[i][j] = 0
			}
		}
	}
}

//方法二：O(1)空间复杂度
func setZeroes(matrix [][]int) {
	var m, n int
	m = len(matrix)
	if m > 0 {
		n = len(matrix[0])
	}
	x, y := -1, -1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				if x == -1 {
					x = i
					y = j
				} else {
					matrix[x][j] = 0
					matrix[i][y] = 0
				}
			}
		}
	}

	if x == -1 {
		return
	}
	for i := 0; i < m; i++ {
		if i != x && matrix[i][y] == 0 {
			for k := 0; k < n; k++ {
				matrix[i][k] = 0
			}
		}
	}

	for j := 0; j < n; j++ {
		if matrix[x][j] == 0 {
			for k := 0; k < m; k++ {
				matrix[k][j] = 0
			}
		}
	}
	for j := 0; j < n; j++ {
		matrix[x][j] = 0
	}
}
