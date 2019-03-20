package main

import (
	"fmt"
)

func core(result *[]string, oneSolution string, index, maxIndex int, dicts []string) {
	if index == maxIndex {
		*result = append(*result, oneSolution)
		return
	}

	for i := 0; i < len(dicts[index]); i++ {
		core(result, oneSolution+dicts[index][i:i+1], index+1, maxIndex, dicts)
	}
}

func letterCombinations(digits string) []string {
	dict := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	dicts := make([]string, len(digits))
	if(len(digits) == 0) {
		return []string{}
	}
	for i := 0; i < len(digits); i++ {
		dicts[i] = dict[digits[i]-'0']
	}
	var result []string
	var oneSolution string
	core(&result, oneSolution, 0, len(digits), dicts)
	return result
}

func main() {
	digits := ""
	res := letterCombinations(digits)
	fmt.Println(res)
}
