package greedystrategy

func canJump(nums []int) bool {
	if nums == nil {
		return false
	}
	currentLimit := 0
	start := 0
	nextStart := 0
	for start < len(nums) {
		currentLimit = start + nums[start]

		// 不能改变下面的两个if语句块的顺序，应对特殊情况：nums = [1]
		if currentLimit >= len(nums)-1 {
			return true
		}
		if currentLimit <= start {
			return false
		}

		nextStart = start + 1
		for i := start + 2; i <= currentLimit; i++ {
			if i+nums[i] >= nextStart+nums[nextStart] {
				nextStart = i
			}
		}
		start = nextStart
	}
	return false
}
