package editstring

func isInterleave(s1 string, s2 string, s3 string) bool {
	s1Array := []rune(s1)
	s2Array := []rune(s2)
	s3Array := []rune(s3)
	s1Len := len(s1Array)
	s2Len := len(s2Array)
	if s1Len+s2Len != len(s3Array) {
		return false
	}
	dp := make([][]bool, 2, 2)
	for i := range dp {
		dp[i] = make([]bool, s2Len+1, s2Len+1)
	}
	dp[0][0] = true
	for l := 1; l <= s2Len; l++ {
		dp[0][l] = s2Array[l-1] == s3Array[l-1] && dp[0][l-1]
	}
	for l1 := 1; l1 <= s1Len; l1++ {
		newL1 := l1 % 2
		oldL1 := (l1 - 1) % 2
		dp[newL1][0] = s3Array[l1-1] == s1Array[l1-1] && dp[oldL1][0]
		for l2 := 1; l2 <= s2Len; l2++ {
			dp[newL1][l2] = (dp[oldL1][l2] && s3Array[l1+l2-1] == s1Array[l1-1]) || (dp[newL1][l2-1] && s3Array[l1+l2-1] == s2Array[l2-1])
		}
	}
	return dp[s1Len%2][s2Len]
}
