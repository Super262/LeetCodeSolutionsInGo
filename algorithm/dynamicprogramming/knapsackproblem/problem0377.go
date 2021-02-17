package knapsackproblem

import "sort"

func combinationSum4(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		if target == 0 {
			return 1
		} else {
			return 0
		}
	}
	sort.Ints(nums)
	numsLen := len(nums)
	dp := make([]int, target+1, target+1)
	dp[0] = 1
	for t := 1; t <= target; t++ {
		for i := 0; i < numsLen; i++ {
			if t >= nums[i] {
				dp[t] += dp[t-nums[i]]
			}
		}
	}
	return dp[target]
}
