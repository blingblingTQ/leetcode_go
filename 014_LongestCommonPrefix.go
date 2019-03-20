package main

func sameAtIndex(strs []string, index int) bool {
	size := len(strs)
	if size == 0 || size == 1 {
		return true
	}
	c := strs[0][index]
	for i := 1; i < size; i++ {
		if len(strs[i]) <= index {
			return false
		}
		if strs[i][index] != c {
			return false
		}
	}
	return true
}

func longestCommonPrefix(strs []string) string {
	var longestStr string
	size := len(strs)
	if size == 0 {
		return longestStr
	}
	firstLen := len(strs[0])
	i := 0
	for i = 0; i < firstLen; i++ {
		if !sameAtIndex(strs, i) {
			break
		}
	}
	return strs[0][:i]
}
