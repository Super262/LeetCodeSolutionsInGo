package bfs

func numIslands(grid [][]byte) int {
	if grid == nil || grid[0] == nil || len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	directions := [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	visited := make([][]bool, len(grid), len(grid))
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]bool, len(grid[0]), len(grid[0]))
	}
	connectedCount := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == '0' || visited[y][x] {
				continue
			}
			bfs0200(&grid, &visited, &directions, y, x)
			connectedCount++
		}
	}
	return connectedCount
}

func bfs0200(grid *[][]byte, visited *[][]bool, directions *[][]int, startY int, startX int) {
	q := make([][]int, 0)
	q = append(q, []int{startY, startX})
	currentSize := 0
	nextY := 0
	nextX := 0
	maxY := len(*grid)
	maxX := len((*grid)[0])
	var currentNode []int
	for len(q) > 0 {
		currentSize = len(q)
		for i := 0; i < currentSize; i++ {
			currentNode = q[0]
			q = q[1:]
			for _, d := range *directions {
				nextY = currentNode[0] + d[0]
				nextX = currentNode[1] + d[1]
				if nextY < 0 || nextY >= maxY || nextX < 0 || nextX >= maxX || (*grid)[nextY][nextX] == '0' || (*visited)[nextY][nextX] {
					continue
				}
				q = append(q, []int{nextY, nextX})
				(*visited)[nextY][nextX] = true
			}
		}
	}
}
