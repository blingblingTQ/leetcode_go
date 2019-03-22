package main

func isPerfectSquare(num int) bool {
	if num < 0 {
		return false
	}
	start, end := 0, num
	for start <= end {
		mid := start + (end-start)/2
		if mid*mid == num {
			return true
		} else if mid*mid < num {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return false
}
