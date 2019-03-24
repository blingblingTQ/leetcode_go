package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minSubArrayLen(s int, nums []int) int {
	minLen := len(nums) + 1
	start, end := 0, 0
	sum := 0
	for end < len(nums) {
		sum += nums[end]
		end++
		for sum >= s {
			minLen = min(minLen, end-start)
			sum -= nums[start]
			start++
		}
	}
	if minLen > len(nums) {
		return 0
	}
	return minLen
}

func main() {
	s := 5
	nums := []int{2, 3, 1, 1, 1, 1}
	fmt.Println(minSubArrayLen(s, nums))
}
