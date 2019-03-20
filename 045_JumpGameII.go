package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func jump(nums []int) int {
	var maxReach, i, last, step int
	for i = 0; i < len(nums); i++ {
		if i > maxReach {
			return -1
		}
		if last < i {
			last = maxReach
			step++
		}
		maxReach = max(maxReach, nums[i]+i)
	}
	return step
}
