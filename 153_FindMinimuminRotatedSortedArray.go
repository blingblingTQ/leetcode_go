package main

import "fmt"

func findMin(nums []int) int {
	start, end := 0, len(nums)-1
	for start < end {
		mid := start + (end-start)/2
		if nums[mid] > nums[end] {
			start = mid + 1
		} else if nums[mid] < nums[end] {
			end = mid
		} else {
			end--
		}
	}
	return nums[start]
}

func main() {
	fmt.Println("vim-go")
}
