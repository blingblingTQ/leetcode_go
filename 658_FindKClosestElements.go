package main

import (
	"fmt"
)

func findClosestElements(arr []int, k int, x int) []int {
	start, end := 0, len(arr)-k
	for start < end {
		mid := start + (end-start)/2
		fmt.Printf("start=%d, end=%d, mid=%d\n", start, end, mid)
		if x-arr[mid] > arr[mid+k]-x {
			fmt.Printf("x-arr[%d]>arr[%d+%d]-%d\n", mid, mid, k, x)
			start = mid + 1
		} else {
			fmt.Printf("x-arr[%d] <= arr[%d+%d]-%d\n", mid, mid, k, x)
			end = mid
		}
	}
	fmt.Printf("return arr[%d:%d]\n", start, start+k)
	return arr[start : start+k]
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	k := 4
	x := 3
	fmt.Println(findClosestElements(arr, k, x))
}
