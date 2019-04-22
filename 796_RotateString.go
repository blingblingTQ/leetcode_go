package main

func rotateString(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}
	if strings.Contains(B+B, A) {
		return true
	}
	return false
}
