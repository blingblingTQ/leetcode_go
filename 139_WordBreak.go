package main

type word []string

func (w word) find(str string) bool {
	for _, s := range w {
		if s == str {
			return true
		}
	}
	return false
}

func wordBreak(s string, wordDict []string) bool {
	size := len(s)

	dp := make([]bool, size+1)
	dp[0] = true
	for i := 1; i <= size; i++ {
		for j := i - 1; j >= 0; j-- {
			if dp[j] && word(wordDict).find(s[j:i]) {
				dp[i] = true
				break
			}

		}
	}
	return dp[size]
}
