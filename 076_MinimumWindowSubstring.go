package main

import (
	"fmt"
	"math"
)

func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}
	dict := make([]int, 128)
	for _, c := range t {
		dict[c]++
	}

	begin, end := 0, 0
	head := 0
	minRes := math.MaxInt32
	counter := len(t)

	for end < len(s) {
		if dict[s[end]] > 0 {
			counter--
		}
		dict[s[end]] -= 1
		end++
		for counter == 0 {
			if end-begin < minRes {
				minRes = end - begin
				head = begin
			}
			if dict[s[begin]] == 0 {
				counter++
			}
			dict[s[begin]] += 1
			begin++
		}
	}
	if minRes == math.MaxInt32 {
		return ""
	}
	return s[head : head+minRes]
}

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Println(minWindow(s, t))
}
