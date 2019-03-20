package main

func plusOne(digits []int) []int {
	carry := 1
	size := len(digits)

	res := make([]int, size+1)
	for i := size - 1; i >= 0; i-- {
		res[i+1] = (digits[i] + carry) % 10
		carry = (digits[i] + carry) / 10
	}

	if carry > 0 {
		res[0] = carry
		return res
	} else {
		return res[1:]
	}
}
