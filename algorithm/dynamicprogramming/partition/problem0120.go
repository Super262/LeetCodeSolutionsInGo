package partition

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([][]int, n, n)
	for i := range dp {
		dp[i] = make([]int, i+1, i+1)
	}
	for i := 0; i < n; i++ {
		dp[n-1][i] = triangle[n-1][i]
	}
	for i := n - 2; i >= 0; i-- {
		for j := range dp[i] {
			dp[i][j] = getMin0120(dp[i+1][j], dp[i+1][j+1]) + triangle[i][j]
		}
	}
	return dp[0][0]
}

func getMin0120(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
