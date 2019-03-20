package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	} else if size == 1 {
		return nums[0]
	}

	//第一种方法，如果抢第一家，那么最后一家就不能抢
	dp1 := make([]int, size)
	dp1[0] = nums[0]
	dp1[1] = nums[0]
	for i := 2; i < size-1; i++ {
		dp1[i] = max(dp1[i-2]+nums[i], dp1[i-1])
	}

	//第二种方法，如果不抢第一家，那么最后一家可以抢
	dp2 := make([]int, size)
	dp2[0] = 0
	dp2[1] = nums[1]
	for i := 2; i < size; i++ {
		dp2[i] = max(dp2[i-2]+nums[i], dp2[i-1])
	}
	return max(dp1[size-2], dp2[size-1])
}
