package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func dot(a []int, b []int) int {
	sum := 0
	for i := 0; i < len(a); i++ {
		sum += a[i] * b[i]
	}
	return sum
}

func dfs(price []int, special [][]int, needs []int) int {
	res := dot(price, needs)
	j := 0
	for _, s := range special {
		clone := make([]int, len(needs))
		for j = 0; j < len(needs); j++ {
			diff := needs[j] - s[j]
			if diff < 0 {
				break
			}
			clone[j] = diff
		}
		if j == len(needs) {
			res = min(res, s[j]+dfs(price, special, clone))
		}
	}
	return res
}

func shoppingOffers(price []int, special [][]int, needs []int) int {
	return dfs(price, special, needs)
}

func main() {
	price := []int{2, 5}
	special := [][]int{[]int{3, 0, 5}, []int{1, 2, 10}}
	needs := []int{3, 2}
	fmt.Println(shoppingOffers(price, special, needs))
}
