package dfs

func dfsGetArea0695(grid *[][]int, startY int, startX int, direction *[][]int) int {
	if startX < 0 || startY < 0 || startX >= len((*grid)[0]) || startY >= len(*grid) || (*grid)[startY][startX] == 0 {
		return 0
	}
	area := 1
	nextY := 0
	nextX := 0
	(*grid)[startY][startX] = 0
	for _, d := range *direction {
		nextY = d[0] + startY
		nextX = d[1] + startX
		area += dfsGetArea0695(grid, nextY, nextX, direction)
	}
	return area
}

func maxAreaOfIsland(grid [][]int) int {
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
	result := 0
	tempResult := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] == 1 {
				tempResult = dfsGetArea0695(&grid, y, x, &direction)
				if tempResult > result {
					result = tempResult
				}
			}
		}
	}
	return result
}
