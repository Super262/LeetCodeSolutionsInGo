package greedystrategy

func canJump(nums []int) bool {
	if nums == nil || len(nums) == 0 {
		return false
	}
	globalBound := 0
	currentBound := 0
	for i := 0; i < len(nums); i++ {
		if i > globalBound {
			continue
		}
		currentBound = i + nums[i]
		if currentBound > globalBound {
			globalBound = currentBound
		}
		if globalBound >= len(nums)-1 {
			return true
		}
	}
	return false
}
