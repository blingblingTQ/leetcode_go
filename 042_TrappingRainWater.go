package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//暴力法
func trap(height []int) int {
	area := 0
	size := len(height)
	for i := 1; i < size-1; i++ {
		maxLeft, maxRight := 0, 0
		for j := i; j >= 0; j-- {
			maxLeft = max(maxLeft, height[j])
		}
		for j := i; j < size; j++ {
			maxRight = max(maxRight, height[j])
		}
		area += min(maxLeft, maxRight) - height[i]
	}
	return area
}

// DP
func trap(height []int) int {
	size := len(height)
	if size == 0 {
		return 0
	}
	maxLeft := make([]int, size)
	maxLeft[0] = height[0]
	for i := 1; i < size; i++ {
		maxLeft[i] = max(maxLeft[i-1], height[i])
	}
	maxRight := make([]int, size)
	maxRight[size-1] = height[size-1]
	for i := size - 2; i >= 0; i-- {
		maxRight[i] = max(maxRight[i+1], height[i])
	}
	area := 0
	for i := 1; i < size-1; i++ {
		area += min(maxLeft[i], maxRight[i]) - height[i]
	}
	return area
}

//two pointers
func trap(height []int) int {
	left, right := 0, len(height)-1
	area := 0
	maxLeft, maxRight := 0, 0
	for left < right {
		if height[left] < height[right] {
			if height[left] >= maxLeft {
				maxLeft = height[left]
			} else {
				area += maxLeft - height[left]
			}
			left++
		} else {
			if height[right] >= maxRight {
				maxRight = height[right]
			} else {
				area += maxRight - height[right]
			}
			right--
		}
	}
	return area
}
