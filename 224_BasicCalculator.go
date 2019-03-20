package main

import (
	"fmt"
	"unicode"
)

func calculate(s string) int {
	stack := make([]int, 0)
	lastSign := 1
	result := 0

	for i := 0; i < len(s); {
		if unicode.IsDigit(rune(s[i])) {
			//如果是数字，则计算出这个数
			digit := 0
			for i < len(s) && unicode.IsDigit(rune(s[i])) {
				digit = 10*digit + int(s[i]-'0')
				i++
			}
			result += lastSign * digit
		} else if s[i] == '+' {
			lastSign = 1
			i++
		} else if s[i] == '-' {
			lastSign = -1
			i++
		} else if s[i] == '(' {
			stack = append(stack, result)
			stack = append(stack, lastSign)
			result = 0
			lastSign = 1
			i++
		} else if s[i] == ')' {
			sign := stack[len(stack)-1]
			prevRes := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			result = result*sign + prevRes
			i++
		} else {
			i++
		}
	}
	return result
}

func main() {
	str := "(2+3)"
	fmt.Println(calculate(str))
}
