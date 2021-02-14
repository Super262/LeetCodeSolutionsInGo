package partition

func integerBreak(n int) int {
	if n < 2 {
		return 0
	}
	dp := make([]int, n+1, n+1)
	dp[2] = 1
	temp1 := 0
	temp2 := 0
	tempMax := 0
	for num := 3; num <= n; num++ {
		for i := 1; i < num; i++ {
			temp1 = i * (num - i)
			temp2 = dp[i] * (num - i)
			if temp1 > temp2 {
				tempMax = temp1
			} else {
				tempMax = temp2
			}
			if tempMax > dp[num] {
				dp[num] = tempMax
			}
		}
	}
	return dp[n]
}
