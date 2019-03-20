package main

func nextGreatestLetter(letters []byte, target byte) byte {
	start, end := 0, len(letters)-1
	for start < end {
		mid := start + (end-start)/2
		if letters[mid] <= target {
			start = mid + 1
		} else {
			end = mid
		}
	}
	if letters[end] > target {
		return letters[end]
	}
	return letters[0]
}
