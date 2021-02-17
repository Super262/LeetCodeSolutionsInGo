package binarysearch

func findMin(nums []int) int {
	numsLen := len(nums)
	high := numsLen - 1
	low := 0
	mid := 0
	for low < high {
		mid = low + (high-low)/2
		if nums[mid] >= nums[low] && nums[mid] >= nums[high] {
			low = mid + 1
		} else if nums[mid] < nums[low] && nums[mid] < nums[high] {
			high = mid
		} else {
			high = mid
		}
	}
	return nums[low]
}
