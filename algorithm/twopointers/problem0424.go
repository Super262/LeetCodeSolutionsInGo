package twopointers

import "math"

func characterReplacement(s string, k int) int {
	if len(s) <= k {
		return len(s)
	}
	counter := make([]int, 26, 26)
	maxFreq := 0
	answer := 0
	right := 0
	for left := range s {
		right = Max0424(right, left)
		for right < len(s) && right-left-maxFreq <= k {
			counter[s[right]-'A']++
			maxFreq = Max0424(maxFreq, counter[s[right]-'A'])
			right++
		}
		if right-left-maxFreq > k {
			answer = Max0424(answer, right-left-1)
		} else {
			answer = Max0424(answer, right-left)
		}
		counter[s[left]-'A']--
		maxFreq = MaxInArray0424(&counter)
	}
	return answer
}

func Max0424(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaxInArray0424(nums *[]int) int {
	result := math.MinInt32
	for _, v := range *nums {
		if v > result {
			result = v
		}
	}
	return result
}
