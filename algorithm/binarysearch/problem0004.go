package binarysearch

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1) + len(nums2)
	if n%2 == 0 {
		return float64(findKth0004(&nums1, &nums2, n/2)+findKth0004(&nums1, &nums2, n/2+1)) / 2.0
	} else {
		return float64(findKth0004(&nums1, &nums2, n/2+1))
	}
}

func findKth0004(sortedArr1 *[]int, sortedArr2 *[]int, k int) int {
	if (*sortedArr1) == nil || len(*sortedArr1) == 0 {
		return (*sortedArr2)[k-1]
	}
	if (*sortedArr2) == nil || len(*sortedArr2) == 0 {
		return (*sortedArr1)[k-1]
	}
	minValue := 0
	if (*sortedArr1)[0] > (*sortedArr2)[0] {
		minValue = (*sortedArr2)[0]
	} else {
		minValue = (*sortedArr1)[0]
	}
	maxValue := 0
	if (*sortedArr1)[len(*sortedArr1)-1] < (*sortedArr2)[len(*sortedArr2)-1] {
		maxValue = (*sortedArr2)[len(*sortedArr2)-1]
	} else {
		maxValue = (*sortedArr1)[len(*sortedArr1)-1]
	}
	midValue := 0
	for minValue+1 < maxValue {
		midValue = minValue + (maxValue-minValue)/2
		if countSmallerAndEqual0004(sortedArr1, midValue)+countSmallerAndEqual0004(sortedArr2, midValue) >= k {
			maxValue = midValue
		} else {
			minValue = midValue
		}
	}
	if countSmallerAndEqual0004(sortedArr1, minValue)+countSmallerAndEqual0004(sortedArr2, minValue) >= k {
		return minValue
	}
	return maxValue
}

func countSmallerAndEqual0004(sortedArr *[]int, target int) int {
	if *sortedArr == nil || len(*sortedArr) == 0 {
		return 0
	}
	start := 0
	end := len(*sortedArr) - 1
	mid := 0
	for start+1 < end {
		mid = start + (end-start)/2
		if (*sortedArr)[mid] > target {
			end = mid
		} else {
			start = mid
		}
	}
	if (*sortedArr)[start] > target {
		return start
	}
	if (*sortedArr)[end] > target {
		return end
	}
	return len(*sortedArr)
}
