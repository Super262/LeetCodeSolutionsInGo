package dfs

func dfsToClean(board [][]byte, startY int, maxY int, startX int, maxX int, direction *[][]int) {
	if startX < 0 || startX >= maxX || startY < 0 || startY >= maxY || board[startY][startX] != 'O' {
		return
	}
	board[startY][startX] = 'A'
	for _, d := range *direction {
		dfsToClean(board, startY+d[0], maxY, startX+d[1], maxX, direction)
	}
}

func solve(board [][]byte) {
	if board == nil {
		return
	}
	height := len(board)
	if height == 0 {
		return
	}
	width := len(board[0])
	if width == 0 {
		return
	}
	direction := [][]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}
	for y := 0; y < height; y++ {
		dfsToClean(board, y, height, 0, width, &direction)
		dfsToClean(board, y, height, width-1, width, &direction)
	}
	for x := 0; x < width; x++ {
		dfsToClean(board, 0, height, x, width, &direction)
		dfsToClean(board, height-1, height, x, width, &direction)
	}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if board[y][x] == 'O' {
				board[y][x] = 'X'
			} else if board[y][x] == 'A' {
				board[y][x] = 'O'
			}
		}
	}
}
