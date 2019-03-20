package main

func core(x float64, n uint) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}

	result := core(x, n>>1)
	result *= result
	if n&0x1 == 1 {
		result *= x
	}
	return result
}

func myPow(x float64, n int) float64 {
	var exp uint = uint(n)
	if n < 0 {
		exp = uint(-n)
	}
	result := core(x, exp)
	if n < 0 {
		return 1.0 / result
	}
	return result
}
