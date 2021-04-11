package interval

func maxCoins(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	n := len(nums) + 2
	values := make([]int, n, n)
	for i := 1; i < n-1; i++ {
		values[i] = nums[i-1]
	}
	values[0] = 1
	values[n-1] = 1
	dp := make([][]int, n, n)
	for i := range dp {
		dp[i] = make([]int, n, n)
	}
	for i := 0; i < n-1; i++ {
		dp[i][i+1] = 0
	}
	for l := 3; l <= n; l++ {
		for i := 0; i < n-l+1; i++ {
			j := i + l - 1
			dp[i][j] = 0
			for k := i + 1; k < j; k++ {
				dp[i][j] = getMax0312(dp[i][j], dp[i][k]+dp[k][j]+values[i]*values[k]*values[j])
			}
		}
	}
	return dp[0][n-1]
}

func getMax0312(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
