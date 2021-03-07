package longestsubsequence

func longestPalindromeSubseq(s string) int {
	if s == "" {
		return 0
	}
	sArray := []rune(s)
	sLen := len(sArray)
	seqLen := make([][]int, sLen, sLen)
	for start := range seqLen {
		seqLen[start] = make([]int, sLen, sLen)
		seqLen[start][start] = 1
	}
	for start := sLen - 1; start >= 0; start-- {
		for end := start + 1; end < sLen; end++ {
			if sArray[start] == sArray[end] {
				seqLen[start][end] = seqLen[start+1][end-1] + 2
			} else {
				seqLen[start][end] = getMax0516(seqLen[start+1][end], seqLen[start][end-1])
			}
		}
	}
	return seqLen[0][sLen-1]
}

func getMax0516(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
