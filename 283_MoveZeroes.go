package main

import (
	"fmt"
)

func moveZeroes(nums []int) {
	read, write := 0, 0
	size := len(nums)
	for read < size {
		for read < size && nums[read] == 0 {
			read++
		}
		if read == size {
			break
		}
		for read < size && nums[read] != 0 {
			nums[write] = nums[read]
			write++
			read++
		}
	}
	for i := write; i < size; i++ {
		nums[i] = 0
	}
}

func moveZeroes(nums []int) {
    zeroCount := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == 0 {
            zeroCount++
        } else if zeroCount > 0 {
            nums[i-zeroCount] = nums[i]
            nums[i] = 0
        }
    }
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}
