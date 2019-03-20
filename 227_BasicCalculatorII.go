package main

import (
	"fmt"
	"unicode"
)

func calculate(s string) int {
	var stack []int
	lastSign := '+'
	for i := 0; i < len(s); {
		if unicode.IsDigit(rune(s[i])) {
			digit := 0
			for i < len(s) && unicode.IsDigit(rune(s[i])) {
				digit = 10*digit + int(s[i]-'0')
				i++
			}
			if lastSign == '+' {
				stack = append(stack, digit)
			} else if lastSign == '-' {
				stack = append(stack, -1*digit)
			} else if lastSign == '*' {
				newValue := stack[len(stack)-1] * digit
				stack[len(stack)-1] = newValue
			} else if lastSign == '/' {
				newValue := stack[len(stack)-1] / digit
				stack[len(stack)-1] = newValue
			}
		} else if s[i] != ' ' {
			//sign
			lastSign = rune(s[i])
			i++
		} else {
			i++
		}
	}

	sum := 0
	for _, v := range stack {
		sum += v
	}
	return sum
}

func main() {
	str := "3/2"
	fmt.Println(calculate(str))
}
