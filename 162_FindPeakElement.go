package main

//常规方法
func findPeakElement(nums []int) int {
	size := len(nums)
	for i := 0; i < size-1; i++ {
		if nums[i] > nums[i+1] {
			return i
		}
	}
	return -1
}

//二分法
func findPeakElement(nums []int) int {
	start, end := 0, len(nums)-1
	for start < end {
		mid := start + (end-start)/2
		if nums[mid] < nums[mid+1] {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return start
}
