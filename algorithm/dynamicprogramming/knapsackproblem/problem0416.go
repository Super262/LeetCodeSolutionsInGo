package knapsackproblem

func canPartition(nums []int) bool {
	if nums == nil {
		return false
	}
	numsLen := len(nums)
	if numsLen < 2 {
		return false
	}
	numsSum := 0
	for i := 0; i < numsLen; i++ {
		numsSum += nums[i]
	}
	if numsSum%2 != 0 {
		return false
	}
	W := numsSum / 2
	dp := make([]bool, W+1, W+1)
	dp[0] = true
	for i := 0; i < numsLen; i++ {
		for w := W; w >= nums[i]; w-- {
			dp[w] = dp[w] || dp[w-nums[i]]
		}
	}
	return dp[W]
}
