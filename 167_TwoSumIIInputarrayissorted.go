package main

func twoSum(numbers []int, target int) []int {
	var res []int
	size := len(numbers)
	if size < 2 {
		return res
	}

	a := 0
	b := size - 1
	for a < b {
		sum := numbers[a] + numbers[b]
		if sum == target {
			res = append(res, a+1)
			res = append(res, b+1)
			return res
		} else if sum < target {
			a++
		} else {
			b--
		}
	}
	return res
}
