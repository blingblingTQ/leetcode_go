package main

func core(nums []int) {
	start, end := 0, len(nums)-1
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func rotate(nums []int, k int) {
	n := len(nums)
	if n == 0 || k <= 0 {
		return
	}
	k = k % n

	core(nums[:n-k])
	core(nums[n-k:])
	core(nums)
}
