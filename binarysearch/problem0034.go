package binarysearch

func findFirst(nums *[]int, target int) int {
	low := 0
	high := len(*nums)
	mid := 0
	for low < high {
		mid = low + (high-low)/2
		if (*nums)[mid] >= target {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

func searchRange(nums []int, target int) []int {
	result := make([]int, 2, 2)
	result[0] = -1
	result[1] = -1
	numsLen := len(nums)
	if nums == nil || len(nums) == 0 {
		return result
	}
	firstIndex := findFirst(&nums, target)
	if firstIndex != numsLen && nums[firstIndex] == target {
		result[0] = firstIndex
		result[1] = firstIndex
	}
	if result[0] != -1 {
		result[1] = findFirst(&nums, target+1) - 1
	}
	return result
}
