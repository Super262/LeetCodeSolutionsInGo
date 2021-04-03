package dfs

import "sort"

func getCombinations0090(nums *[]int, startIndex int, endIndex int, isCaptured *[]bool, tempResult *[]int, result *[][]int) {
	if startIndex < endIndex {
		for i := startIndex; i < endIndex; i++ {
			if i > 0 && (*nums)[i] == (*nums)[i-1] && !(*isCaptured)[i-1] {
				continue
			} else {
				*tempResult = append(*tempResult, (*nums)[i])
				(*isCaptured)[i] = true
				*result = append(*result, append([]int{}, *tempResult...))
				getCombinations0090(nums, i+1, endIndex, isCaptured, tempResult, result)
				(*isCaptured)[i] = false
				*tempResult = (*tempResult)[0 : len(*tempResult)-1]
			}
		}
	}
}

func subsetsWithDup(nums []int) [][]int {
	result := make([][]int, 1, 1)
	if nums == nil {
		return result
	}
	numsLen := len(nums)
	if numsLen == 0 {
		return result
	}
	sort.Ints(nums)
	isCaptured := make([]bool, numsLen, numsLen)
	tempResult := make([]int, 0, numsLen)
	getCombinations0090(&nums, 0, numsLen, &isCaptured, &tempResult, &result)
	return result
}
