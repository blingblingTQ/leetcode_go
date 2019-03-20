package main

import (
	"fmt"
)

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
func lengthOfLongestSubstring(s string) int {
	dict := [256]int{}
	for i := 0; i < 256; i++ {
		dict[i] = -1
	}
	start := -1
	maxLen := 0
	for i := 0; i < len(s); i++ {
		if dict[s[i]] > start {
			start = dict[s[i]]
		}
		dict[s[i]] = i
		maxLen = max(maxLen, i-start)
	}
	return maxLen
}

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}
