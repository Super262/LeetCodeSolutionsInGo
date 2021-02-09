package binarysearch

func isBadVersion(version int) bool {
	return version == 2
}

func firstBadVersion(n int) int {
	low := 1
	high := n
	mid := 0
	for low < high {
		mid = low + (high-low)/2
		if isBadVersion(mid) {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}
