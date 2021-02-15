package knapsackproblem

func change(amount int, coins []int) int {
	if amount == 0 || coins == nil {
		return 0
	}
	coinTypes := len(coins)
	dp := make([]int, amount+1, amount+1)
	dp[0] = 1
	for ct := 0; ct < coinTypes; ct++ {
		for a := coins[ct]; a <= amount; a++ {
			dp[a] += dp[a-coins[ct]]
		}
	}
	return dp[amount]
}
