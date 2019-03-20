package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func largestRectangleArea(heights []int) int {
	size := len(heights)
	if size == 0 {
		return 0
	}
	area := 0
	stack := make([]int, 0)
	newHeights := append(heights, 0)
	for i := 0; i < len(newHeights); {
		if len(stack) == 0 || newHeights[i] >= newHeights[stack[len(stack)-1]] {
			stack = append(stack, i)
			i++
		} else {
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			length := 0
			if len(stack) == 0 {
				length = i
			} else {
				length = i - stack[len(stack)-1] - 1
			}
			area = max(area, newHeights[index]*length)
		}
	}
	return area
}
