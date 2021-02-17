package binarysearch

func mySqrt(x int) int {
	if x < 2 {
		return x
	}
	low := 1
	high := x
	tempResult := 0
	for low <= high {
		mid := low + (high-low)/2
		tempResult = mid * mid
		if tempResult == x {
			return mid
		} else if tempResult > x {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return high
}
