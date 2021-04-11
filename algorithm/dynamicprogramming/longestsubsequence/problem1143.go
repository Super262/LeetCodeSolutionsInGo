package longestsubsequence

func longestCommonSubsequence(text1 string, text2 string) int {
	t1Array := []rune(text1)
	t2Array := []rune(text2)
	t1Len := len(t1Array)
	t2Len := len(t2Array)
	if t1Len == 0 || t2Len == 0 {
		return 0
	}
	dp := make([][]int, 2, 2)
	for i := range dp {
		dp[i] = make([]int, t2Len+1, t2Len+1)
	}
	for l1 := 1; l1 <= t1Len; l1++ {
		newL1 := l1 % 2
		oldL1 := (l1 - 1) % 2
		for l2 := 1; l2 <= t2Len; l2++ {
			if t1Array[l1-1] == t2Array[l2-1] {
				dp[newL1][l2] = dp[oldL1][l2-1] + 1
			} else {
				dp[newL1][l2] = getMax1143(dp[oldL1][l2], dp[newL1][l2-1])
			}
		}
	}
	return dp[t1Len%2][t2Len]
}

func getMax1143(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
