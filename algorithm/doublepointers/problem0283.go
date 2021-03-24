package doublepointers

func moveZeroes(nums []int) {
	// Minimize the total number of operations!
	if nums == nil || len(nums) == 0 {
		return
	}
	left := 0
	right := 0
	for right < len(nums) {
		if nums[right] != 0 {
			nums[left] = nums[right]
			left++
		}
		right++
	}
	for left < len(nums) {
		nums[left] = 0
		left++
	}
}
