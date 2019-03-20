package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxProduct(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}
	localMin, localMax, globalMax := nums[0], nums[0], nums[0]
	for i := 1; i < size; i++ {
		a := localMin * nums[i]
		b := localMax * nums[i]
		localMin = min(min(a, b), nums[i])
		localMax = max(max(a, b), nums[i])
		globalMax = max(localMax, globalMax)
	}
	return globalMax
}
