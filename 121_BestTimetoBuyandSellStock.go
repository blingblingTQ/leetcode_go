package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {
	if prices == nil {
		return 0
	}
	minBuy := 0
	profit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < prices[minBuy] {
			minBuy = i
		} else {
			profit = max(profit, prices[i]-prices[minBuy])
		}
	}
	return profit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}
