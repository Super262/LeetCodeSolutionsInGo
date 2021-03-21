package doublepointers

func minSubArrayLen(target int, nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	minLen := 0
	left := 0
	right := 0
	currentSum := 0
	for right < len(nums) {
		currentSum += nums[right]
		for left <= right && currentSum >= target {
			if minLen <= 0 || right-left+1 < minLen {
				minLen = right - left + 1
			}
			currentSum -= nums[left]
			left++
		}
		right++
	}
	return minLen
}
