package queue_stack

func NumIslands(grid [][]byte) int {
	if 0 == len(grid) {
		return 0
	}

	var visited [][]byte
	num := 0
	columnLength := len(grid[0])

	for range grid {
		visited = append(visited, make([]byte, columnLength))
	}

	for row, values := range grid {
		for column, value := range values {
			if value == '1' && visited[row][column] == 0 {
				num++
				bfs(grid, row, column, visited)
			}
		}
	}

	return num
}

func bfs(grid [][]byte, row, column int, visited [][]byte) {
	visited[row][column] = 1

	if row > 0 {
		if visited[row-1][column] == 0 && grid[row-1][column] == '1' {
			bfs(grid, row-1, column, visited)
		}
	}

	if row < len(grid)-1 {
		if visited[row+1][column] == 0 && grid[row+1][column] == '1' {
			bfs(grid, row+1, column, visited)
		}
	}

	if column > 0 {
		if visited[row][column-1] == 0 && grid[row][column-1] == '1' {
			bfs(grid, row, column-1, visited)
		}
	}

	if column < len(grid[0])-1 {
		if visited[row][column+1] == 0 && grid[row][column+1] == '1' {
			bfs(grid, row, column+1, visited)
		}
	}
}
