package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestConsecutive(nums []int) int {
	m := make(map[int]int)
	maxLength := 0
	for _, num := range nums {
		if _, ok := m[num]; !ok {
			sum := m[num-1] + m[num+1] + 1
			m[num] = sum
			maxLength = max(maxLength, sum)

			m[num-m[num-1]] = sum
			m[num+m[num+1]] = sum
		}
	}
	return maxLength
}

func longestConsecutive(nums []int) int {
	m := make(map[int]bool)
	maxLength := 0
	for _, num := range nums {
		m[num] = true
	}

	for _, num := range nums {
		if _, ok := m[num-1]; !ok {
			count := 1
			currNum := num
			for m[currNum+1] == true {
				currNum += 1
				count += 1
			}
			maxLength = max(maxLength, count)
		}
	}
	return maxLength
}
