package main

import (
	"fmt"
	"math"
)

func max(a, b int) {
	if a > b {
		return a
	}
	return b
}

func maxSubArray(nums []int) int {
	size := len(nums)
	sum := 0
	maxSum := math.MinInt32
	for i := 0; i < size; i++ {
		sum += nums[i]
		maxSum = max(maxSum, sum)
		if sum < 0 {
			sum = 0
		}
	}
	return maxSum
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}
