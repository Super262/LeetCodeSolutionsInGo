package twopointers

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	if s == "" || k == 0 {
		return 0
	}
	sArray := []rune(s)
	chFreq := make([]int, 256, 256)
	result := 0
	uniqueCount := 0
	left := 0
	right := 0
	for right < len(sArray) {
		chFreq[int(sArray[right])]++
		if chFreq[int(sArray[right])] == 1 {
			uniqueCount++
		}
		right++
		for left < right && uniqueCount > k {
			chFreq[int(sArray[left])]--
			if chFreq[int(sArray[left])] == 0 {
				uniqueCount--
			}
			left++
		}
		if right-left > result {
			result = right - left
		}
	}
	return result
}
