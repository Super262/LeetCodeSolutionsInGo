package knapsackproblem

func findMaxForm(strs []string, m int, n int) int {
	if strs == nil {
		return 0
	}
	strsLen := len(strs)
	if strsLen == 0 {
		return 0
	}
	zerosCount := make([]int, strsLen, strsLen)
	onesCount := make([]int, strsLen, strsLen)
	dp := make([][]int, m+1, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1, n+1)
	}
	var tempArray []rune
	for i, s := range strs {
		tempArray = []rune(s)
		for _, ch := range tempArray {
			if ch == '0' {
				zerosCount[i]++
			} else {
				onesCount[i]++
			}
		}
	}
	tempSum := 0
	for i := 0; i < strsLen; i++ {
		for zc := m; zc >= zerosCount[i]; zc-- {
			for oc := n; oc >= onesCount[i]; oc-- {
				tempSum = 1 + dp[zc-zerosCount[i]][oc-onesCount[i]]
				if tempSum > dp[zc][oc] {
					dp[zc][oc] = tempSum
				}
			}
		}
	}
	return dp[m][n]
}
