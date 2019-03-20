package main

import (
	"fmt"
)

func longestPalindrome(s string) string {
	//palindrome[i, j]表示从i开始，j结束的子串，是否是回文串
	size := len(s)
	if size == 0 {
		return s
	}
	maxStart, maxEnd := 0, 0
	maxLen := 0
	var palindrome [1001][1001]bool
	//每个单独的字母肯定是一个回文串
	for i := 0; i < size; i++ {
		palindrome[i][i] = true
		if maxLen < 1 {
			maxLen = 1
			maxStart = i
			maxEnd = i
		}
	}
	//如果相邻的字母相同，肯定也是回文串
	for i := 0; i < size-1; i++ {
		if s[i] == s[i+1] {
			palindrome[i][i+1] = true
			if maxLen < 2 {
				maxLen = 2
				maxStart = i
				maxEnd = i + 1
			}
		}
	}
	//如果palindrome[i+1,j-1]=true，并且s[i]=s[j]，那么palindrome[i,j]=true
	for subLen := 3; subLen <= size; subLen++ {
		//对于每个长度为subLen的子串
		for i := 0; i < size-subLen+1; i++ {
			if palindrome[i+1][i+subLen-2] == true && s[i] == s[i+subLen-1] {
				palindrome[i][i+subLen-1] = true
				if maxLen < subLen {
					maxLen = subLen
					maxStart = i
					maxEnd = i + subLen - 1
				}
			}
		}
	}
	return s[maxStart : maxEnd+1]
}

func main() {
	s := ""
	fmt.Println(longestPalindrome(s))
}
