package main

import "container/list"
import "fmt"

//时间复杂度O(mn), O(1)空间复杂度
func isSubsequence(s string, t string) bool {
	sb := []byte(s)
	tb := []byte(t)
	start := 0
	for _, c := range sb {
		i := 0
		for i = start; i < len(tb); i++ {
			if tb[i] == c {
				start = i + 1
				break
			}
		}
		if i == len(tb) {
			return false
		}
	}
	return true
}

//两个栈
func isSubsequence(s string, t string) bool {
	stack1, stack2 := list.New(), list.New()
	for _, c := range []byte(s) {
		stack1.PushBack(c)
	}
	for _, c := range []byte(t) {
		stack2.PushBack(c)
	}

	for stack2.Len() > 0 {
		if stack1.Len() == 0 {
			return true
		}
		c1 := stack1.Back()
		c2 := stack2.Back()
		if c1.Value == c2.Value {
			stack1.Remove(c1)
			stack2.Remove(c2)
		} else {
			stack2.Remove(c2)
		}
	}
	return stack1.Len() == 0
}

func main() {
	s := "abc"
	t := "ahbgdc"
	fmt.Println(isSubsequence(s, t))
}
