package knapsackproblem

func coinChange(coins []int, amount int) int {
	if amount == 0 || coins == nil {
		return 0
	}
	dp := make([]int, amount+1, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0
	temp := 0
	for change := 1; change <= amount; change++ {
		for _, coin := range coins {
			if change >= coin {
				temp = dp[change-coin] + 1
				if temp < dp[change] {
					dp[change] = temp
				}
			}
		}
	}
	if dp[amount] == amount+1 {
		return -1
	} else {
		return dp[amount]
	}
}
