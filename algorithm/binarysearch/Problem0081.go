package binarysearch

func search0081(nums []int, target int) bool {
	if nums == nil || len(nums) == 0 {
		return false
	}
	start := 0
	end := len(nums) - 1
	for start+1 < end {
		mid := start + (end-start)/2
		if nums[mid] < nums[end] {
			if nums[mid] <= target && target <= nums[end] {
				start = mid
			} else {
				end = mid
			}
		} else if nums[mid] > nums[end] {
			if nums[start] <= target && target <= nums[mid] {
				end = mid
			} else {
				start = mid
			}
		} else {
			end--
		}
	}
	return nums[start] == target || nums[end] == target
}
