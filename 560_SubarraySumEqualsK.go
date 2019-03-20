package main

func subarraySum(nums []int, k int) int {
	count, sum := 0, 0
	m := map[int]int{
		0: 1,
	}

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		count += m[sum-k]
		m[sum]++
	}
	return count
}
