package pathinmatrix

func uniquePaths(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	dp := make([][]int, m, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n, n)
	}
	for x := 0; x < n; x++ {
		dp[0][x] = 1
	}
	for y := 1; y < m; y++ {
		dp[y][0] = 1
	}
	for y := 1; y < m; y++ {
		for x := 1; x < n; x++ {
			dp[y][x] = dp[y-1][x] + dp[y][x-1]
		}
	}
	return dp[m-1][n-1]
}
