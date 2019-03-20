package main

func peakIndexInMountainArray(A []int) int {
	start, end := 0, len(A)-1
	for start <= end {
		mid := start + (end-start)/2
		if A[mid] > A[mid+1] && A[mid] > A[mid-1] {
			return mid
		} else if A[mid] < A[mid+1] {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return start
}
