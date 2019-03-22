package main

func arrangeCoins(n int) int {
	start, end := 1, n
	for start <= end {
		mid := start + (end-start)/2
		sum := (mid + 1) * mid / 2
		if sum == n {
			return mid
		} else if sum > n {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return start - 1
}
