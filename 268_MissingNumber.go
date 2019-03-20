package main

func missingNumber(nums []int) int {
	sum := 0
	n := len(nums)
	for _, num := range nums {
		sum += num
	}
	return n*(n+1)/2 - sum
}

func missingNumber2(nums []int) int {
	n := len(nums)
    res := n
	for i := 0; i < n; i++ {
		res ^= nums[i]
		res ^= i
	}
	return res
}


func missingNumber3(nums []int) int {
	sort.Ints(nums)
	lo, hi := 0, len(nums)
	for lo < hi {
        mid := lo + (hi-lo)/2
		if nums[mid] > mid {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}