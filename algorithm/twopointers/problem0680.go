package twopointers

func isPalindrome(charArray []rune, start int, end int) bool {
	for start <= end {
		if charArray[start] != charArray[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func validPalindrome(s string) bool {
	charArray := []rune(s)
	i := 0
	j := len(charArray) - 1
	for i < j {
		if charArray[i] != charArray[j] {
			break
		}
		i++
		j--
	}
	if i >= j {
		return true
	} else {
		if isPalindrome(charArray, i, j-1) || isPalindrome(charArray, i+1, j) {
			return true
		} else {
			return false
		}
	}
}
