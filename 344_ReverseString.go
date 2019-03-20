package main

func reverseString(s []byte) {
	size := len(s)
	lo := 0
	hi := size - 1
	for lo < hi {
		s[lo], s[hi] = s[hi], s[lo]
		lo++
		hi--
	}
}
