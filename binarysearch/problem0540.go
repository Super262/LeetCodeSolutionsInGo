package binarysearch

func singleNonDuplicate(nums []int) int {
	numsLen := len(nums)
	if numsLen == 1 {
		return nums[0]
	}
	low := 0
	high := numsLen - 1
	mid := 0
	for low < high {
		mid = low + (high-low)/2
		if mid%2 == 0 {
			if mid > 0 && nums[mid] == nums[mid-1] {
				high = mid
			} else {
				low = mid
			}
		} else {
			if mid > 0 && nums[mid] == nums[mid-1] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return nums[low]
}
