package main

func strStr(haystack string, needle string) int {
	hLen := len(haystack)
	nLen := len(needle)
	if hLen < nLen {
		return -1
	}

	for i := 0; i <= hLen-nLen; i++ {
		j := 0
		for j = 0; j < nLen; j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		if j == nLen {
			return i
		}
	}
	return -1
}
