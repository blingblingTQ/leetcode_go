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

func coinChange(coins []int, amount int) int {
    if amount < 0 {
        return -1
    }
    dp := make([]int, amount+1)
    coinsSize := len(coins)
    for i := 1; i <= amount; i++ {
        minVal := math.MaxInt32
        for j := 0; j < coinsSize; j++ {
            if i >= coins[j] && dp[i-coins[j]] != -1{
                minVal = min(dp[i-coins[j]]+1, minVal)
            }
        }
        if minVal == math.MaxInt32 {
            dp[i] = -1
        } else {
            dp[i] = minVal
        }
    }
    return dp[amount]
}

func main() {
	coins := []int{1, 2, 5}
	amount := 11
	fmt.Println(coinChange(coins, amount))
}
