package dfs

import (
	"strconv"
)

func validIpPart(digits *[]rune, start int, end int) bool {
	if end == start {
		return false
	}
	if end-start > 1 {
		if (*digits)[start] == '0' {
			return false
		} else {
			i, _ := strconv.Atoi(string((*digits)[start:end]))
			if i > 255 {
				return false
			}
		}
	}
	return true
}

func getCombinations0093(prefix *[]rune, partIndex int, s *[]rune, startIndex int, endIndex int, result *[]string) {
	if partIndex == 3 {
		if validIpPart(s, startIndex, endIndex) {
			*result = append(*result, string(append(*prefix, (*s)[startIndex:endIndex]...)))
		}
	} else {
		partLen := 0
		for i := startIndex + 1; i < endIndex && i < startIndex+4; i++ {
			partLen = i - startIndex
			if validIpPart(s, startIndex, i) {
				*prefix = append(*prefix, (*s)[startIndex:i]...)
				*prefix = append(*prefix, '.')
				getCombinations0093(prefix, partIndex+1, s, i, endIndex, result)
				*prefix = (*prefix)[0 : len(*prefix)-partLen-1]
			}
		}
	}
}

func restoreIpAddresses(s string) []string {
	result := make([]string, 0)
	sArray := []rune(s)
	sLen := len(sArray)
	if sLen < 4 {
		return result
	}
	prefix := make([]rune, 0, sLen+4)
	getCombinations0093(&prefix, 0, &sArray, 0, sLen, &result)
	return result
}
