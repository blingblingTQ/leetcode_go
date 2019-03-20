package main

func firstUniqChar(s string) int {
	m := make(map[rune]int)
	for _, c := range s {
		m[c] += 1
	}
	for i, c := range s {
		if m[c] == 1 {
			return i
		}
	}
	return -1
}
