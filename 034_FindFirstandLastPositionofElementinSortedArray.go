package main

func firstPosition(nums []int, target int) int {
	start, end := 0, len(nums)-1
	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] < target {
			start = mid + 1
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			if mid == 0 || (mid > 0 && nums[mid-1] != target) {
				return mid
			}
			end = mid - 1
		}
	}
	return -1
}

func lastPosition(nums []int, target int) int {
	start, end := 0, len(nums)-1
	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] < target {
			start = mid + 1
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			if mid == len(nums)-1 || (mid < len(nums)-1 && nums[mid+1] != target) {
				return mid
			}
			start = mid + 1
		}
	}
	return -1
}

func searchRange(nums []int, target int) []int {
	first := firstPosition(nums, target)
	last := lastPosition(nums, target)
	return []int{first, last}
}
