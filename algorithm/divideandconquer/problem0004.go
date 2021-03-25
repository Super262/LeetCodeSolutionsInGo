package divideandconquer

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1) + len(nums2)
	if n%2 == 0 {
		return float64(findKth0004(&nums1, 0, &nums2, 0, n/2)+findKth0004(&nums1, 0, &nums2, 0, n/2+1)) / 2.0
	} else {
		return float64(findKth0004(&nums1, 0, &nums2, 0, n/2+1))
	}
}

func findKth0004(sortedArr1 *[]int, start1 int, sortedArr2 *[]int, start2 int, k int) int {
	if *sortedArr1 == nil || len(*sortedArr1) == 0 || start1 >= len(*sortedArr1) {
		return (*sortedArr2)[start2+k-1]
	}
	if *sortedArr2 == nil || len(*sortedArr2) == 0 || start2 >= len(*sortedArr2) {
		return (*sortedArr1)[start1+k-1]
	}
	if k == 1 {
		if (*sortedArr1)[start1] < (*sortedArr2)[start2] {
			return (*sortedArr1)[start1]
		}
		return (*sortedArr2)[start2]
	}
	halfKOfArr1 := 0
	if start1+k/2-1 < len(*sortedArr1) {
		halfKOfArr1 = (*sortedArr1)[start1+k/2-1]
	} else {
		halfKOfArr1 = math.MaxInt32
	}
	halfKOfArr2 := 0
	if start2+k/2-1 < len(*sortedArr2) {
		halfKOfArr2 = (*sortedArr2)[start2+k/2-1]
	} else {
		halfKOfArr2 = math.MaxInt32
	}
	if halfKOfArr1 > halfKOfArr2 {
		return findKth0004(sortedArr1, start1, sortedArr2, start2+k/2, k-k/2)
	} else {
		return findKth0004(sortedArr1, start1+k/2, sortedArr2, start2, k-k/2)
	}
}
