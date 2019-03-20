package main

import (
	"fmt"
)

func numDecodings(s string) int {
	size := len(s)
	if size == 0 {
		return 0
	}

	dp := make([]int, size+1)
	dp[0] = 1
	for i := 1; i < size+1; i++ {
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}
		if i >= 2 && s[i-2:i] <= "26" && s[i-2:i] >= "10" {
			dp[i] += dp[i-2]
		}
	}
	return dp[size]
}

func main() {
	str := "23"
	fmt.Println(numDecodings(str))
}
