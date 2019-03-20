package main

func threeSum(nums []int) [][]int {
	var res [][]int

	if len(nums) <= 2 {
		return res
	}

	sort.Ints(nums)

	var a, b, c int
	for a = 0; a < len(nums)-2; a++ {
		if a > 0 && nums[a-1] == nums[a] {
			continue
		}

		b = a + 1
		c = len(nums) - 1
		for b < c {
			sum := nums[a] + nums[b] + nums[c]
			if sum == 0 {
				res = append(res, []int{nums[a], nums[b], nums[c]})
				b++
				c--

				for b < c && nums[b-1] == nums[b] {
					b++
				}
				for b < c && nums[c] == nums[c+1] {
					c--
				}
			} else if sum < 0 {
				b++
			} else {
				c--
			}

		}
	}
	return res
}
