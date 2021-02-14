package pathinmatrix

func minPathSum(grid [][]int) int {
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
	dp := make([][]int, height, height)
	for i := range dp {
		dp[i] = make([]int, width, width)
	}
	dp[0][0] = grid[0][0]
	for y := 1; y < height; y++ {
		dp[y][0] = dp[y-1][0] + grid[y][0]
	}
	for x := 1; x < width; x++ {
		dp[0][x] = dp[0][x-1] + grid[0][x]
	}
	for y := 1; y < height; y++ {
		for x := 1; x < width; x++ {
			if dp[y-1][x] > dp[y][x-1] {
				dp[y][x] = dp[y][x-1] + grid[y][x]
			} else {
				dp[y][x] = dp[y-1][x] + grid[y][x]
			}
		}
	}
	return dp[height-1][width-1]
}
