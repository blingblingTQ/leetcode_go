package main

type NumArray struct {
	Sum []int
}

func Constructor(nums []int) NumArray {
	numArr := NumArray{
		Sum: make([]int, len(nums)+1),
	}
	for i := 1; i <= len(nums); i++ {
		numArr.Sum[i] = numArr.Sum[i-1] + nums[i-1]
	}
	return numArr
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.Sum[j+1] - this.Sum[i]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
