package main

import (
	"fmt"
	"unicode"
	"math"
)


func myAtoi(str string) int {
	i := 0
	sign  := 1
	if len(str) == 0 {
		return 0
	}

	//跳过空格
	for i < len(str) && str[i] == ' ' {
		i++
	}

	//判断符号
	if i < len(str) && str[i] == '+' {
		sign = 1
		i++
	} else if i < len(str) && str[i] == '-' {
		sign = -1
		i++
	}

	//计算数值
	digit := 0
	for i < len(str) && unicode.IsDigit(rune(str[i])) {
		digit = 10*digit + int(str[i]-'0')
		if digit*sign >= math.MaxInt32 {
			return math.MaxInt32
		} else if digit*sign <= math.MinInt32 {
			return math.MinInt32
		}
		i++
	}
	//计算最终结果
	return int(digit * sign)
}

func main() {
	str := "-91283472332"
	fmt.Println(myAtoi(str))
}
