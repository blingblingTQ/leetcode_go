package main

func match(left, right rune) bool {
	if (left == '(' && right == ')') ||
		(left == '[' && right == ']') ||
		(left == '{' && right == '}') {
		return true
	}
	return false
}

func isValid(s string) bool {
	stack := make([]rune, 0)
	for _, c := range s {
		if c == '{' || c == '(' || c == '[' {
			stack = append(stack, c)
		} else {
			size := len(stack)
			if size == 0 {
				return false
			}
			last := stack[len(stack)-1]
			if match(last, c) {
				stack = stack[:size-1]
			} else {
				return false
			}
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
