package main

import (
	"fmt"
)

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func main() {
	n := 3
	fmt.Println(climbStairs(n))
	n = 2
	fmt.Println(climbStairs(n))
}
