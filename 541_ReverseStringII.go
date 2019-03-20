package main

import (
	"fmt"
)

func reverse(s []byte) {
	size := len(s)
	for lo, hi := 0, size-1; lo < hi; {
		s[lo], s[hi] = s[hi], s[lo]
		lo++
		hi--
	}
}

func reverseStr(s string, k int) string {
	strByte := []byte(s)
	size := len(strByte)
	for i := 0; i < size; i += 2 * k {
		//反转从i开始，长度为k的数组
		if i+k > size {
			reverse(strByte[i:])
		} else {
			fmt.Printf("reverse [%d:%d]\n", i, i+k)
			reverse(strByte[i : i+k])
		}
	}
	return string(strByte)
}

func main() {
	s := "a"
	fmt.Println(reverseStr(s, 2))
}
