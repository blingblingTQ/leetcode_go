package main

func dfs(rooms [][]int, visited *[]bool, roomId int) {
	if (*visited)[roomId] {
		return
	}
	(*visited)[roomId] = true

	for i := 0; i < len(rooms[roomId]); i++ {
		dfs(rooms, visited, rooms[roomId][i])
	}
}

func canVisitAllRooms(rooms [][]int) bool {
	nums := len(rooms)
	visited := make([]bool, nums)
	dfs(rooms, &visited, 0)

	for i := 0; i < nums; i++ {
		if !visited[i] {
			return false
		}
	}
	return true
}
