package editstring

func getMinInt0650(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func minSteps(n int) int {
	dp := make([]int, n+1, n+1)
	for i := range dp {
		dp[i] = 2 * n
	}
	dp[0] = 0
	dp[1] = 0
	for num := 2; num <= n; num++ {
		for i := 1; i < num; i++ {
			if num%i == 0 {
				dp[num] = getMinInt0650(dp[num], dp[i]+(num/i))
			}
		}
	}
	return dp[n]
}
