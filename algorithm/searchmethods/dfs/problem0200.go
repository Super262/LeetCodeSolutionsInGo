package dfs

func dfsExploreArea0200(grid *[][]byte, startY int, maxY int, startX int, maxX int, direction *[][]int) {
	if startX < 0 || startY < 0 || startX >= maxX || startY >= maxY || (*grid)[startY][startX] == '0' {
		return
	}
	nextY := 0
	nextX := 0
	(*grid)[startY][startX] = '0'
	for _, d := range *direction {
		nextY = d[0] + startY
		nextX = d[1] + startX
		dfsExploreArea0200(grid, nextY, maxY, nextX, maxX, direction)
	}
}

func numIslands(grid [][]byte) int {
	if grid == nil {
		return 0
	}
	height := len(grid)
	if height == 0 {
		return 0
	}
	width := len(grid[0])
	if width == 0 {
		return 0
	}
	direction := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	count := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] == '1' {
				dfsExploreArea0200(&grid, y, height, x, width, &direction)
				count++
			}
		}
	}
	return count
}
