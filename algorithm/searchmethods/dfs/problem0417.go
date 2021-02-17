package dfs

func dfsToExplore0417(heightMatrix *[][]int, canReach *[][]bool, currentY int, maxY int, currentX int, maxX int, direction *[][]int) {
	if (*canReach)[currentY][currentX] {
		return
	}
	(*canReach)[currentY][currentX] = true
	prevY := 0
	prevX := 0
	for _, d := range *direction {
		prevY = currentY + d[1]
		prevX = currentX + d[0]
		if prevX >= 0 && prevX < maxX && prevY >= 0 && prevY < maxY && (*heightMatrix)[prevY][prevX] >= (*heightMatrix)[currentY][currentX] {
			dfsToExplore0417(heightMatrix, canReach, prevY, maxY, prevX, maxX, direction)
		}
	}
}

func pacificAtlantic(matrix [][]int) [][]int {
	result := make([][]int, 0, 0)
	if matrix == nil {
		return result
	}
	height := len(matrix)
	if height == 0 {
		return result
	}
	width := len(matrix[0])
	if width == 0 {
		return result
	}
	canReachP := make([][]bool, height, height)
	canReachA := make([][]bool, height, height)
	for i := 0; i < height; i++ {
		canReachA[i] = make([]bool, width, width)
		canReachP[i] = make([]bool, width, width)
	}
	direction := [][]int{{1, 0}, {0, 1}, {0, -1}, {-1, 0}}
	for x := 0; x < width; x++ {
		dfsToExplore0417(&matrix, &canReachP, 0, height, x, width, &direction)
		dfsToExplore0417(&matrix, &canReachA, height-1, height, x, width, &direction)
	}
	for y := 0; y < height; y++ {
		dfsToExplore0417(&matrix, &canReachP, y, height, 0, width, &direction)
		dfsToExplore0417(&matrix, &canReachA, y, height, width-1, width, &direction)
	}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if canReachA[y][x] && canReachP[y][x] {
				result = append(result, []int{y, x})
			}
		}
	}
	return result
}
