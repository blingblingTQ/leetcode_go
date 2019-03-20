package main

func romanToInt(s string) int {
	size := len(s)
	m := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	sum := 0
	for i := 0; i < size; i++ {
		sign := 1
		switch s[i] {
		case 'I':
			if i != size-1 && (s[i+1] == 'V' || s[i+1] == 'X') {
				sign = -1
			}
		case 'X':
			if i != size-1 && (s[i+1] == 'L' || s[i+1] == 'C') {
				sign = -1
			}
		case 'C':
			if i != size-1 && (s[i+1] == 'D' || s[i+1] == 'M') {
				sign = -1
			}
		}
		sum += (sign * m[s[i]])
	}
	return sum
}
