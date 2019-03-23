package main

import (
	"fmt"
	"math"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func valid(piles []int, k int, H int) bool {
	spend := 0
	for _, pile := range piles {
		spend += int(math.Ceil(float64(pile) / float64(k)))
	}
	return spend <= H
}

func minEatingSpeed(piles []int, H int) int {
	maxSpeed := 0
	for _, pile := range piles {
		maxSpeed = max(maxSpeed, pile)
	}
	start, end := 1, maxSpeed
	for start < end {
		mid := start + (end-start)/2
		if valid(piles, mid, H) {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return start
}

func main() {
	piles := []int{312884470}
	H := 968709470
	fmt.Println(minEatingSpeed(piles, H))
}
