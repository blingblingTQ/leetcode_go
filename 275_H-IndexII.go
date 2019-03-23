package main

func hIndex(citations []int) int {
	size := len(citations)
	start, end := 0, size-1
	for start <= end {
		mid := start + (end-start)/2
		if citations[mid] == size-mid {
			return size - mid
		} else if citations[mid] < size-mid {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return size - start
}
