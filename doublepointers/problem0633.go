package doublepointers

import "math"

func judgeSquareSum(c int) bool {
	if c >= 0 {
		i := 0
		j := int(math.Sqrt(float64(c)))
		var tempResult int
		for i <= j {
			tempResult = i*i + j*j
			if tempResult == c {
				return true
			} else if tempResult > c {
				j--
			} else {
				i++
			}
		}
	}
	return false
}
