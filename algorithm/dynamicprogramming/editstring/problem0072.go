package editstring

func getMinInt0072(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func minDistance2(word1 string, word2 string) int {
	word1Array := []rune(word1)
	word2Array := []rune(word2)
	word1Len := len(word1Array)
	word2Len := len(word2Array)
	dp := make([][]int, word1Len+1, word1Len+1)
	for i := range dp {
		dp[i] = make([]int, word2Len+1, word2Len+1)
	}
	for l := 1; l <= word1Len; l++ {
		dp[l][0] = l
	}
	for l := 1; l <= word2Len; l++ {
		dp[0][l] = l
	}
	for l1 := 1; l1 <= word1Len; l1++ {
		for l2 := 1; l2 <= word2Len; l2++ {
			if word1Array[l1-1] == word2Array[l2-1] {
				dp[l1][l2] = getMinInt0072(dp[l1-1][l2-1], getMinInt0072(dp[l1-1][l2]+1, dp[l1][l2-1]+1))
			} else {
				dp[l1][l2] = getMinInt0072(dp[l1-1][l2-1]+1, getMinInt0072(dp[l1-1][l2]+1, dp[l1][l2-1]+1))
			}
		}
	}
	return dp[word1Len][word2Len]
}
