package longestsubsequence

import "sort"

func findLongestChain(pairs [][]int) int {
	if pairs == nil {
		return 0
	}
	pairsLen := len(pairs)
	if pairsLen == 0 {
		return 0
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})
	dp := make([]int, pairsLen, pairsLen)
	dp[0] = 1
	result := dp[0]
	preMax := 0
	for i := 1; i < pairsLen; i++ {
		preMax = 0
		for j := 0; j < i; j++ {
			if pairs[j][1] < pairs[i][0] {
				if dp[j] > preMax {
					preMax = dp[j]
				}
			}
		}
		dp[i] = preMax + 1
		if dp[i] > result {
			result = dp[i]
		}
	}
	return result
}
