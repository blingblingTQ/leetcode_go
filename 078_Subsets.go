package main

import (
	"fmt"
)

func dfs(nums []int, path []int) [][]int{
	result := [][]int{}
	result = append(result, path)
	for i := 0; i < len(nums); i++ {
		nextPath := []int{}
		nextPath = append(nextPath, path...)
		nextPath = append(nextPath, nums[i])
		result = append(result, dfs(nums[i+1:], nextPath)...)
	}
	return result
}

func subsets(nums []int) [][]int {
	return dfs(nums, []int{})
}

func subsets2(nums []int) [][]int {
	result := [][]int{[]int{}}

	for i := 0; i < len(nums); i++ {
		n := len(result)
		for j := 0; j < n; j++ {
			cop := []int{}
			cop = append(cop, result[j]...)
			cop = append(cop, nums[i])
			result = append(result, cop)
		}
	}
	return result
}


func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))
}
