package main

func mySqrt(x int) int {
	if x < 0 {
		//illegle
		return -1
	}

	lo, hi := 0, x
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if mid*mid == x {
			return mid
		} else if mid*mid > x {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return lo - 1
}
