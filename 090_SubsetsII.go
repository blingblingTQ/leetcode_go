package main

import (
	"fmt"
)

func helper(result *[][]int, path []int, nums []int, index int) {
	if index <= len(nums) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*result = append(*result, tmp)
	}

	for i := index; i < len(nums); i++ {
		if i != index && nums[i] == nums[i-1] {
			continue
		}
		path = append(path, nums[i])
		helper(result, path, nums, i+1)
		path = path[:len(path)-1]
	}
}

func subsetsWithDup(nums []int) [][]int {
	var result [][]int
	var path []int
	sort.Ints(nums)
	helper(&result, path, nums, 0)
	return result
}

func main() {
	nums := []int{1, 2, 2}
	fmt.Println(subsetsWithDup(nums))
}
