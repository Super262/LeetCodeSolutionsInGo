package hashtable

func longestConsecutive(nums []int) int {
	numDict := make(map[int]bool)
	for _, n := range nums {
		numDict[n] = true
	}
	maxLen := 0
	for _, n := range nums {
		_, existed := numDict[n-1]
		if existed {
			continue
		}
		nextN := n + 1
		_, existed = numDict[nextN]
		for existed {
			nextN++
			_, existed = numDict[nextN]
		}
		if maxLen < nextN-n {
			maxLen = nextN - n
		}
	}
	return maxLen
}
