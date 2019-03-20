package main

func fourSum(nums []int, target int) [][]int {
	var res [][]int
	size := len(nums)
	if size <= 3 {
		return res
	}

	sort.Ints(nums)

	var a, b, c, d int
	for a = 0; a < size-3; a++ {
		if a > 0 && nums[a-1] == nums[a] {
			continue
		}
		for d = size - 1; d > a; d-- {
			if d < size-1 && nums[d] == nums[d+1] {
				continue
			}

			b = a + 1
			c = d - 1
			for b < c {
				sum := nums[a] + nums[b] + nums[c] + nums[d]
				if sum == target {
					res = append(res, []int{nums[a], nums[b], nums[c], nums[d]})
					b++
					c--

					for b < c && nums[b] == nums[b-1] {
						b++
					}
					for b < c && nums[c] == nums[c+1] {
						c--
					}

				} else if sum < target {
					b++
				} else {
					c--
				}
			}
		}
	}
	return res
}
