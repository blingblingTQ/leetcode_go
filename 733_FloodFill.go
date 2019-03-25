package main

import "fmt"

func dfs(image [][]int, sr int, sc int, newColor int) {
	if image[sr][sc] == newColor {
		return
	}
	prevColor := image[sr][sc]
	image[sr][sc] = newColor

	if sr > 0 && image[sr-1][sc] == prevColor {
		dfs(image, sr-1, sc, newColor)
	}
	if sr <= len(image)-2 && image[sr+1][sc] == prevColor {
		dfs(image, sr+1, sc, newColor)
	}
	if sc > 0 && image[sr][sc-1] == prevColor {
		dfs(image, sr, sc-1, newColor)
	}
	if sc <= len(image[0])-2 && image[sr][sc+1] == prevColor {
		dfs(image, sr, sc+1, newColor)
	}
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	dfs(image, sr, sc, newColor)
	return image
}

func main() {
	image := [][]int{[]int{0, 0, 0}, []int{0, 1, 1}}
	newColor := 1
	sr, sc := 1, 1

	floodFill(image, sr, sc, newColor)
}
