package main

//方法一
func sortColors(nums []int) {
	counts := []int{0, 0, 0}
	for _, num := range nums {
		counts[num]++
	}

	write := 0
	for counts[0] > 0 {
		nums[write] = 0
		write++
		counts[0]--
	}
	for counts[1] > 0 {
		nums[write] = 1
		write++
		counts[1]--
	}
	for counts[2] > 0 {
		nums[write] = 2
		write++
		counts[2]--
	}
}

func sortColors(nums []int) {
	lo, hi := 0, len(nums)-1
	index := 0
	for index <= hi {
		if nums[index] == 0 {
			nums[index], nums[lo] = nums[lo], nums[index]
			index++
			lo++
		} else if nums[index] == 2 {
			nums[index], nums[hi] = nums[hi], nums[index]
			hi--
		} else {
			index++
		}
	}
}
