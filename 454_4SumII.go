package main

func fourSumCount(A []int, B []int, C []int, D []int) int {
	m := make(map[int]int)
	for _, a := range A {
		for _, b := range B {
			m[a+b]++
		}
	}
	res := 0
	for _, c := range C {
		for _, d := range D {
			res += m[-c-d]
		}
	}
	return res
}
