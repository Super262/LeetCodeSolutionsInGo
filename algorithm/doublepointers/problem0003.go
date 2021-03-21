package doublepointers

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	existedChars := make(map[rune]int)
	sArray := []rune(s)
	maxLen := 0
	left := 0
	right := 0
	for right < len(sArray) {
		chCount, _ := existedChars[sArray[right]]
		for chCount > 0 && left <= right {
			existedChars[sArray[left]] = 0
			left++
			chCount, _ = existedChars[sArray[right]]
		}
		existedChars[sArray[right]] = 1
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
		right++
	}
	return maxLen
}
