package main

func maxProfit(prices []int) int {
	profit := 0
	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}
