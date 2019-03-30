package main

import "fmt"
import "unicode"

func decodeString(s string) string {
	var res string
	if len(s) == 0 {
		return s
	}

	i := 0
	strStack := []string{}
	digitStack := []int{}
	for i < len(s) {
		if unicode.IsDigit(rune(s[i])) {
			digit := 0
			for unicode.IsDigit(rune(s[i])) {
				digit = digit*10 + int(s[i]-'0')
				i++
			}
			digitStack = append(digitStack, digit)
		} else if s[i] == '[' {
			strStack = append(strStack, res)
			res = ""
			i++
		} else if s[i] == ']' {
			preStr := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]
			digit := digitStack[len(digitStack)-1]
			digitStack = digitStack[:len(digitStack)-1]
			for k := 0; k < digit; k++ {
				preStr += res
			}
			res = preStr
			i++
		} else {
			res += s[i : i+1]
			i++
		}
	}
	return res
}

func main() {
	s := "3[a2[c]]"
	fmt.Println(decodeString(s))
}
