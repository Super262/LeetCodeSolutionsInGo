package doublepointers

func reverseVowels(s string) string {
	sArray := []rune(s)
	i := 0
	j := len(sArray) - 1
	var temp rune
	for i < j {
		for i < j {
			if isVowel(sArray[i]) {
				break
			}
			i++
		}
		for i < j {
			if isVowel(sArray[j]) {
				break
			}
			j--
		}
		if i < j {
			temp = sArray[i]
			sArray[i] = sArray[j]
			sArray[j] = temp
			i++
			j--
		}
	}
	return string(sArray)
}
func isVowel(ch rune) bool {
	if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' || ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U' {
		return true
	}
	return false
}
