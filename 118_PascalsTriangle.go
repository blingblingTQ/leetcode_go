package main

func next(nums []int) []int {
	size := len(nums)
	res := make([]int, size+1)
	res[0] = 1
	res[size] = 1
	for i := 1; i < size; i++ {
		res[i] = nums[i-1] + nums[i]
	}
	return res
}

func generate(numRows int) [][]int {
	var res [][]int
	tmp := []int{1}
	for i := 1; i <= numRows; i++ {
		res = append(res, tmp)
		tmp = next(tmp)
	}
	return res
}
