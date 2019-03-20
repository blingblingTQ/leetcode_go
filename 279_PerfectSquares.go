package main

import (
	"fmt"
	"math"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func numSquares(n int) int {
	if n == 0 {
		return 0
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		minVal := math.MaxInt32
		for j := 1; j*j <= i; j++ {
			minVal = min(minVal, dp[i-j*j]+1)
		}
		dp[i] = minVal
	}
	return dp[n]
}

func main() {
	n := 12
	fmt.Println(numSquares(n))
}
