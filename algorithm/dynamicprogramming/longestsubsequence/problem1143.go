package longestsubsequence

import "strings"

func longestCommonSubsequence(text1 string, text2 string) int {
	if strings.Compare(text1, "") == 0 || strings.Compare(text2, "") == 0 {
		return 0
	}
	t1Array := []rune(text1)
	t2Array := []rune(text2)
	t1Len := len(t1Array)
	t2Len := len(t2Array)
	dp := make([][]int, t1Len+1, t1Len+1)
	for i := range dp {
		dp[i] = make([]int, t2Len+1, t2Len+1)
	}
	for l1 := 1; l1 <= t1Len; l1++ {
		for l2 := 1; l2 <= t2Len; l2++ {
			if t1Array[l1-1] == t2Array[l2-1] {
				dp[l1][l2] = dp[l1-1][l2-1] + 1
			} else {
				if dp[l1][l2-1] > dp[l1-1][l2] {
					dp[l1][l2] = dp[l1][l2-1]
				} else {
					dp[l1][l2] = dp[l1-1][l2]
				}
			}
		}
	}
	return dp[t1Len][t2Len]
}
