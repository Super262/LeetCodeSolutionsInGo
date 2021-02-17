package greedystrategy

func maxSubArray(nums []int) int {
	if nums == nil {
		return 0
	}
	numsLen := len(nums)
	if numsLen == 0 {
		return 0
	}
	preSum := nums[0]
	result := preSum
	for i := 1; i < numsLen; i++ {
		if preSum <= 0 {
			preSum = nums[i]
		} else {
			preSum += nums[i]
		}
		if preSum > result {
			result = preSum
		}
	}
	return result
}
