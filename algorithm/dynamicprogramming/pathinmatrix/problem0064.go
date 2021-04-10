package pathinmatrix

func minPathSum(grid [][]int) int {
	height := len(grid)
	width := len(grid[0])
	dp := make([][]int, 2, 2)
	for i := range dp {
		dp[i] = make([]int, width, width)
	}
	dp[0][0] = grid[0][0]
	for y := 0; y < height; y++ {
		newY := y % 2
		oldY := (y - 1) % 2
		for x := 0; x < width; x++ {
			if x == 0 && y == 0 {
				continue
			} else if x != 0 && y != 0 {
				dp[newY][x] = getMin0064(dp[newY][x-1], dp[oldY][x]) + grid[y][x]
			} else if x != 0 {
				dp[newY][x] = dp[newY][x-1] + grid[y][x]
			} else {
				dp[newY][x] = dp[oldY][x] + grid[y][x]
			}
		}
	}
	return dp[(height-1)%2][width-1]
}

func getMin0064(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
