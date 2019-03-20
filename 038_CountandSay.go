package main

func next(s string) string {
	size := len(s)
	count := 1
	var nextStr string
	for i := 1; i < size; i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			nextStr += strconv.Itoa(count) + s[i-1:i]
			count = 1
		}
	}
	nextStr += strconv.Itoa(count) + s[size-1:]
	return nextStr
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	var ret = "1"
	for i := 1; i < n; i++ {
		ret = next(ret)
	}
	return ret
}
