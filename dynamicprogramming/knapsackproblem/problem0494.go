package knapsackproblem

func findTargetSumWays(nums []int, S int) int {
	if nums == nil {
		return 0
	}
	numsLen := len(nums)
	numsSum := 0
	for i := 0; i < numsLen; i++ {
		numsSum += nums[i]
	}
	tempSum := S + numsSum
	if tempSum%2 != 0 || S > numsSum {
		return 0
	}
	W := tempSum / 2
	dp := make([]int, W+1, W+1)
	dp[0] = 1
	for i := 0; i < numsLen; i++ {
		for w := W; w >= nums[i]; w-- {
			dp[w] = dp[w] + dp[w-nums[i]]
		}
	}
	return dp[W]
}
