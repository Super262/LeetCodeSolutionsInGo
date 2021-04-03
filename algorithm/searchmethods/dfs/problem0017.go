package dfs

import "strings"

func getCombinations0017(prefix *[]rune, digits *[]rune, startIndex int, endIndex int, digitToLetter *[][]rune, result *[]string) {
	if startIndex == endIndex {
		*result = append(*result, string(*prefix))
	} else if startIndex < endIndex {
		curDigits := (*digitToLetter)[(*digits)[startIndex]-'0']
		for _, d := range curDigits {
			(*prefix)[startIndex] = d
			getCombinations0017(prefix, digits, startIndex+1, endIndex, digitToLetter, result)
			(*prefix)[startIndex] = '\000'
		}
	}
}

func letterCombinations(digits string) []string {
	result := make([]string, 0)
	if strings.Compare(digits, "") == 0 {
		return result
	}
	digitToLetter := make([][]rune, 10, 10)
	digitToLetter[2] = []rune{'a', 'b', 'c'}
	digitToLetter[3] = []rune{'d', 'e', 'f'}
	digitToLetter[4] = []rune{'g', 'h', 'i'}
	digitToLetter[5] = []rune{'j', 'k', 'l'}
	digitToLetter[6] = []rune{'m', 'n', 'o'}
	digitToLetter[7] = []rune{'p', 'q', 'r', 's'}
	digitToLetter[8] = []rune{'t', 'u', 'v'}
	digitToLetter[9] = []rune{'w', 'x', 'y', 'z'}
	chArray := []rune(digits)
	lenOfChArr := len(chArray)
	prefix := make([]rune, lenOfChArr, lenOfChArr)
	getCombinations0017(&prefix, &chArray, 0, lenOfChArr, &digitToLetter, &result)
	return result
}
