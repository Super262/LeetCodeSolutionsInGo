package mathematics

func convertToBase7(num int) string {
	isNeg := false
	if num < 0 {
		isNeg = true
		num = -num
	}
	maxBase := 1
	digitsCount := 1
	for num/maxBase >= 7 {
		maxBase *= 7
		digitsCount++
	}
	var resultTop int
	var result []rune
	if isNeg {
		result = make([]rune, digitsCount+1, digitsCount+1)
		result[0] = '-'
		resultTop = 1
	} else {
		result = make([]rune, digitsCount, digitsCount)
		resultTop = 0
	}
	tempFactor := 0
	for i := 0; i < digitsCount; i++ {
		tempFactor = num / maxBase
		result[resultTop] = rune('0' + tempFactor)
		resultTop++
		num -= tempFactor * maxBase
		maxBase /= 7
	}
	return string(result)
}
