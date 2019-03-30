package main

func dfs(nums []int, index int, currSum int, S int, res *int) {
	if index == len(nums) {
		if currSum == S {
			(*res)++
		}
		return
	}

	dfs(nums, index+1, currSum+nums[index], S, res)
	dfs(nums, index+1, currSum-nums[index], S, res)
}

func findTargetSumWays(nums []int, S int) int {
	res := 0
	dfs(nums, 0, 0, S, &res)
	return res
}
