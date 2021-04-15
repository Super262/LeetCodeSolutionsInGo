package twopointers

func numSub(s string) int {
	if s == "" {
		return 0
	}
	answer := 0
	right := 1
	bound := 1000000007
	for left := 0; left < len(s); left++ {
		if s[left] != '1' {
			continue
		}
		right = getMax1513(right, left+1)
		for right < len(s) && s[right] == '1' {
			right++
		}
		answer += right - left
		answer %= bound
	}
	return answer
}

func getMax1513(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
