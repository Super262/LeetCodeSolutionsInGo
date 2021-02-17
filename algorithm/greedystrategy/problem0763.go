package greedystrategy

func partitionLabels(S string) []int {
	lastIndex := make([]int, 26, 26)
	sArray := []rune(S)
	for i, ch := range sArray {
		lastIndex[int(ch)-int('a')] = i
	}
	result := make([]int, 0)
	curEnd := lastIndex[int(sArray[0])-int('a')]
	curStart := 0
	for curP, ch := range sArray {
		if curP > curEnd {
			result = append(result, curEnd-curStart+1)
			curStart = curEnd + 1
		}
		if lastIndex[int(ch)-int('a')] > curEnd {
			curEnd = lastIndex[int(ch)-int('a')]
		}
	}
	result = append(result, curEnd-curStart+1)
	return result
}
