package main

import (
	"fmt"
)

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	cumulateProduct := 1
	res[0] = 1
	for i := 1; i < len(nums); i++ {
		cumulateProduct *= nums[i-1]
		res[i] = cumulateProduct
	}

	cumulateProduct = 1
	for i := len(nums) - 2; i >= 0; i-- {
		cumulateProduct *= nums[i+1]
		res[i] *= cumulateProduct
	}
	return res
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums))
}
