package binarysearch

import "math"

func findPeakElement(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return math.MinInt32
	}
	start := 0
	end := len(nums) - 1
	mid := 0
	for start+1 < end {
		mid = start + (end-start)/2
		if mid > 0 && nums[mid] < nums[mid-1] {
			end = mid
		} else if mid < len(nums)-1 && nums[mid] < nums[mid+1] {
			start = mid
		} else {
			return mid
		}
	}
	if nums[start] > nums[end] {
		return start
	} else {
		return end
	}
}
