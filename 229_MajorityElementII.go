package main

func majorityElement(nums []int) []int {
	var res []int
	if len(nums) == 0 {
		return res
	}
	candidate1 := nums[0]
	candidate2 := nums[0]
	count1 := 0
	count2 := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == candidate1 {
			count1++
		} else if nums[i] == candidate2 {
			count2++
		} else if count1 == 0 {
			candidate1 = nums[i]
			count1++
		} else if count2 == 0 {
			candidate2 = nums[i]
			count2++
		} else {
			count1--
			count2--
		}
	}

	count1 = 0
	count2 = 0
	for _, num := range nums {
		if num == candidate1 {
			count1++
		} else if num == candidate2 {
			count2++
		}
	}
	if count1 > len(nums)/3 {
		res = append(res, candidate1)
	}
	if count2 > len(nums)/3 {
		res = append(res, candidate2)
	}
	return res
}
