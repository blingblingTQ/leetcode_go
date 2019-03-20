package main

func removeDuplicates(nums []int) int {
	wIdx := 1
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] != nums[i-1] {
			nums[wIdx] = nums[i]
			wIdx++
		}
	}
	return wIdx
}
