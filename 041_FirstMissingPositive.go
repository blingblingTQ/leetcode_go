package main

func firstMissingPositive(nums []int) int {
	if len(nums) == 0 {
		return 1
	}

	newNums := []int{0}
	newNums = append(newNums, nums...)

	i := 0
	for i < len(newNums) {
		if newNums[i] > 0 && newNums[i] < len(newNums) && newNums[i] != newNums[newNums[i]] {
			newNums[i], newNums[newNums[i]] = newNums[newNums[i]], newNums[i]
		} else {
			i++
		}
	}

	for i = 0; i < len(newNums); i++ {
		if newNums[i] != i {
			return i
		}
	}
	return len(newNums)
}
