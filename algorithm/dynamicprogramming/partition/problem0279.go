package partition

func generateSquares0279(n int) []int {
	result := make([]int, 0)
	tempRes := 0
	for i := 1; i <= n; i++ {
		tempRes = i * i
		if tempRes <= n {
			result = append(result, tempRes)
		} else {
			break
		}
	}
	return result
}

func numSquares(n int) int {
	prevSquares := generateSquares0279(n)
	dp := make([]int, n+1, n+1)
	for i := range dp {
		dp[i] = i
	}
	dp[0] = 1
	dp[1] = 1
	temp := 0
	for num := 2; num <= n; num++ {
		for _, s := range prevSquares {
			if s > num {
				break
			} else if s == num {
				dp[num] = 1
			} else {
				temp = dp[num-s] + 1
				if temp < dp[num] {
					dp[num] = temp
				}
			}
		}
	}
	return dp[n]
}
