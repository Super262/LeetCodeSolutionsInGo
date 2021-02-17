package backtracking

import "strings"

func isPalindrome0131(s *[]rune, start int, end int) bool {
	low := start
	high := end - 1
	for low < high {
		if (*s)[low] == (*s)[high] {
			low++
			high--
		} else {
			return false
		}
	}
	return true
}

func getPartitions0131(s *[]rune, start int, end int, tempResult *[]string, result *[][]string) {
	if start == end {
		*result = append(*result, append([]string{}, *tempResult...))
	} else if start < end {
		for i := start + 1; i <= end; i++ {
			if isPalindrome0131(s, start, i) {
				*tempResult = append(*tempResult, string((*s)[start:i]))
				getPartitions0131(s, i, end, tempResult, result)
				*tempResult = (*tempResult)[0 : len(*tempResult)-1]
			}
		}
	}
}

func partition(s string) [][]string {
	result := make([][]string, 0, 0)
	if strings.Compare(s, "") == 0 {
		return result
	}
	sArray := []rune(s)
	sLen := len(sArray)
	tempResult := make([]string, 0)
	getPartitions0131(&sArray, 0, sLen, &tempResult, &result)
	return result
}
