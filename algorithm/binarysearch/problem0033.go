package binarysearch

func search0033(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}
	start := 0
	end := len(nums) - 1
	mid := 0
	for start+1 < end {
		mid = start + (end-start)/2
		if nums[mid] >= nums[start] {
			if nums[start] <= target && target <= nums[mid] {
				end = mid
			} else {
				start = mid
			}
		} else {
			if nums[mid] <= target && target <= nums[end] {
				start = mid
			} else {
				end = mid
			}
		}
	}
	if nums[start] == target {
		return start
	}
	if nums[end] == target {
		return end
	}
	return -1
}
