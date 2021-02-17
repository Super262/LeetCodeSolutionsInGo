package greedystrategy

func checkPossibility(nums []int) bool {
	if nums == nil {
		return true
	}
	numsLen := len(nums)
	if numsLen < 2 {
		return true
	}
	count := 0
	for i := 1; i < numsLen; i++ {
		if nums[i] < nums[i-1] {
			count++
			if count > 1 {
				return false
			}
			if i-2 >= 0 && nums[i-2] > nums[i] {
				nums[i] = nums[i-1]
			} else {
				nums[i-1] = nums[i]
			}
		}
	}
	return true
}
