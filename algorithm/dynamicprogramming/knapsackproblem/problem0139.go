package knapsackproblem

func arraySame0139(a *[]rune, startA int, endA int, b *[]rune, startB int, endB int) bool {
	if endA-startA != endB-startB {
		return false
	}
	i := startA
	j := startB
	for i < endA && j < endB {
		if (*a)[i] != (*b)[j] {
			return false
		}
		i++
		j++
	}
	return true
}

func wordBreak(s string, wordDict []string) bool {
	sArray := []rune(s)
	sLen := len(sArray)
	dictLen := len(wordDict)
	wordArrayDict := make([][]rune, dictLen, dictLen)
	wordLenDict := make([]int, dictLen, dictLen)
	for i := 0; i < dictLen; i++ {
		wordArrayDict[i] = []rune(wordDict[i])
		wordLenDict[i] = len(wordArrayDict[i])
	}
	dp := make([]bool, sLen+1, sLen+1)
	dp[0] = true
	for l := 1; l <= sLen; l++ {
		for k, w := range wordArrayDict {
			if l >= wordLenDict[k] && arraySame0139(&sArray, l-wordLenDict[k], l, &w, 0, wordLenDict[k]) {
				dp[l] = dp[l] || dp[l-wordLenDict[k]]
			}
		}
	}
	return dp[sLen]
}
