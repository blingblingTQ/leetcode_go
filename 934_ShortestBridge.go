package main

import "fmt"

type Node struct {
	Row   int
	Col   int
	Depth int
}

func dfs(A [][]int, rows int, cols int, i int, j int) {
	if i < 0 || i >= rows || j < 0 || j >= cols {
		return
	}
	if A[i][j] == 1 {
		A[i][j] = 2
		dfs(A, rows, cols, i-1, j)
		dfs(A, rows, cols, i+1, j)
		dfs(A, rows, cols, i, j-1)
		dfs(A, rows, cols, i, j+1)
	}
}

func neighbors(A [][]int, rows int, cols int, i int, j int) []Node {
	nodes := []Node{}
	if i-1 >= 0 {
		nodes = append(nodes, Node{Row: i - 1, Col: j, Depth: 0})
	}
	if i+1 < rows {
		nodes = append(nodes, Node{Row: i + 1, Col: j, Depth: 0})
	}
	if j-1 >= 0 {
		nodes = append(nodes, Node{Row: i, Col: j - 1, Depth: 0})
	}
	if j+1 < cols {
		nodes = append(nodes, Node{Row: i, Col: j + 1, Depth: 0})
	}
	return nodes
}

func shortestBridge(A [][]int) int {
	var rows, cols int
	rows = len(A)
	if rows > 0 {
		cols = len(A[0])
	}

	//use dfs to search one island
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if A[i][j] == 1 {
				dfs(A, rows, cols, i, j)
				goto FIND
			}
		}
	}
FIND:
	queue := []Node{}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if A[i][j] == 2 {
				node := Node{
					Row:   i,
					Col:   j,
					Depth: 0,
				}
				queue = append(queue, node)
			}
		}
	}

	for len(queue) != 0 {
		for k := 0; k < len(queue); k++ {
			node := queue[0]
			queue = queue[1:]

			neighborNodes := neighbors(A, rows, cols, node.Row, node.Col)
			for _, neighbor := range neighborNodes {
				if A[neighbor.Row][neighbor.Col] == 0 {
					A[neighbor.Row][neighbor.Col] = 2
					neighbor.Depth = node.Depth + 1
					queue = append(queue, neighbor)

				} else if A[neighbor.Row][neighbor.Col] == 1 {
					//find
					return node.Depth
				}
			}
		}
	}
	return 0
}

func printMap(A [][]int) {
	rows := len(A)
	cols := len(A[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("%d ", A[i][j])
		}
		fmt.Println()
	}
}

func main() {
	A := [][]int{[]int{0, 1}, []int{1, 0}}
	fmt.Println(shortestBridge(A))
}
