package main

import (
	"fmt"
)

func findLength(A []int, B []int) int {
	sizeA := len(A)
	sizeB := len(B)

	dp := make([][]int, sizeA+1)
	for i := 0; i <= sizeA; i++ {
		dp[i] = make([]int, sizeB+1)
	}

	max := 0
	for i := 1; i <= sizeA; i++ {
		for j := 1; j <= sizeB; j++ {
			if A[i-1] == B[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			if max < dp[i][j] {
				max = dp[i][j]
			}
		}
	}
	return max
}

func main() {
	A := []int{1, 2, 3, 2, 1}
	B := []int{3, 2, 1, 4, 7}
	fmt.Println(findLength(A, B))
}
