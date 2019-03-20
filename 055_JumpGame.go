package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func canJump(nums []int) bool {
	maxReach, i := 0, 0
	for i = 0; i < len(nums) && i <= maxReach; i++ {
		maxReach = max(maxReach, nums[i]+i)
	}
	return i == len(nums)
}
