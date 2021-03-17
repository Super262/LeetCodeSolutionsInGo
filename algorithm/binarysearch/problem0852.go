package binarysearch

func peakIndexInMountainArray(arr []int) int {
	if arr == nil {
		return -1
	}
	start := 0
	end := len(arr) - 1
	if end == -1 {
		return -1
	}
	var mid int
	var temp int
	for start < end {
		mid = start + (end-start)/2
		temp = mid + 1
		if temp <= end && arr[temp] >= arr[mid] {
			start = temp
			continue
		}
		temp = mid - 1
		if temp >= start && arr[temp] >= arr[mid] {
			end = temp
			continue
		}
		return mid
	}
	return start
}
