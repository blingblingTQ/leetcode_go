package main

import (
	"fmt"
)

//方法一
func findDuplicate1(nums []int) int {
	size := len(nums)
	for i := 0; i < size; i++ {
		for nums[i] != i {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	return -1
}

//
func findDuplicate2(nums []int) int {
	result := 0
	for bit := 0; bit < 32; bit++ {
		mask := 1 << uint(bit)
		a, b := 0, 0
		for i := 0; i < len(nums); i++ {
			if i != 0 && i&mask != 0 {
				a++
			}
			if nums[i]&mask != 0 {
				b++
			}
		}
		if b > a {
			result |= mask
		}
	}
	return result
}

func findDuplicate3(nums []int) int {
	start, end := 0, len(nums)-1
	for start <= end {
		cnt := 0
		mid := start + (end-start)/2
		for _, num := range nums {
			if num <= mid {
				cnt++
			}
		}
		if cnt < mid {
			start = mid + 1
		} else {
			end = mid - 1
		}

	}
	return start
}

func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

func main() {
	nums := []int{1, 3, 4, 2, 2}
	fmt.Println(findDuplicate(nums))
}
