package mathematics

import "math"

func maximumProduct(nums []int) int {
	min1 := math.MaxInt32
	min2 := math.MaxInt32
	max1 := math.MinInt32
	max2 := math.MinInt32
	max3 := math.MinInt32
	for _, num := range nums {
		if num > max1 {
			max3 = max2
			max2 = max1
			max1 = num
		} else if num > max2 {
			max3 = max2
			max2 = num
		} else if num > max3 {
			max3 = num
		}
		if num < min1 {
			min2 = min1
			min1 = num
		} else if num < min2 {
			min2 = num
		}
	}
	return getMax0628(max1*max2*max3, min1*min2*max1)
}

func getMax0628(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
