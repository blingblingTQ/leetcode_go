package main

func core(result *[]string, oneSolution string, lLeft, rLeft int) {
	if lLeft == 0 && rLeft == 0 {
		*result = append(*result, string(oneSolution))
	}
	if lLeft > 0 {
		core(result, oneSolution+"(", lLeft-1, rLeft)
	}
	if rLeft > lLeft {
		core(result, oneSolution+")", lLeft, rLeft-1)
	}
}

func generateParenthesis(n int) []string {
	result := []string{}
	var oneSolution string
	if n == 0 {
		return result
	}
	lLeft, rLeft := n, n
	core(&result, oneSolution, lLeft, rLeft)
	return result
}
