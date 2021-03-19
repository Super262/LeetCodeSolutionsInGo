package greedystrategy

func jump(nums []int) int {
	if nums == nil {
		return 0
	}
	result := 0
	start := 0
	nextStart := 0
	currentBound := 0
	for start < len(nums)-1 {
		result++
		currentBound = start + nums[start]
		if currentBound >= len(nums)-1 {
			return result
		}
		nextStart = start
		for i := start + 1; i <= currentBound; i++ {
			if i+nums[i] >= nextStart+nums[nextStart] {
				nextStart = i
			}
		}
		start = nextStart
	}
	return result
}
