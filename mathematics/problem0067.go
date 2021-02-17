package mathematics

func addBinary(a string, b string) string {
	aArray := []rune(a)
	bArray := []rune(b)
	i := len(aArray) - 1
	j := len(bArray) - 1
	result := make([]rune, 0)
	hasCarry := false
	tempSum := 0
	for i >= 0 && j >= 0 {
		tempSum = (int(aArray[i]) - int('0')) + (int(bArray[j]) - int('0'))
		hasCarry = appendDigit0067(hasCarry, tempSum, &result)
		i--
		j--
	}
	for i >= 0 {
		tempSum = int(aArray[i]) - int('0')
		hasCarry = appendDigit0067(hasCarry, tempSum, &result)
		i--
	}
	for j >= 0 {
		tempSum = int(bArray[j]) - int('0')
		hasCarry = appendDigit0067(hasCarry, tempSum, &result)
		j--
	}
	if hasCarry {
		result = append(result, '1')
	}
	for i, j = 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return string(result)
}

func appendDigit0067(hasCarry bool, tempSum int, result *[]rune) bool {
	if hasCarry {
		tempSum++
	}
	if tempSum == 3 {
		*result = append(*result, '1')
		return true
	} else if tempSum == 2 {
		*result = append(*result, '0')
		return true
	} else {
		*result = append(*result, rune('0'+tempSum))
		return false
	}
}
