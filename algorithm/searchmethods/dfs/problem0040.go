package dfs

import "sort"

func getCombinations0040(nums *[]int, numsLen int, startIndex int, target int, isCaptured *[]bool, tempResult *[]int, result *[][]int) {
	for i := startIndex; i < numsLen; i++ {
		if i > 0 && (*nums)[i] == (*nums)[i-1] && !(*isCaptured)[i-1] {
			continue
		} else {
			if (*nums)[i] == target {
				*tempResult = append(*tempResult, (*nums)[i])
				*result = append(*result, append([]int{}, *tempResult...))
				*tempResult = (*tempResult)[0 : len(*tempResult)-1]
			} else if (*nums)[i] < target {
				(*isCaptured)[i] = true
				*tempResult = append(*tempResult, (*nums)[i])
				getCombinations0040(nums, numsLen, i+1, target-(*nums)[i], isCaptured, tempResult, result)
				*tempResult = (*tempResult)[0 : len(*tempResult)-1]
				(*isCaptured)[i] = false
			}
		}
	}
}

func combinationSum2(candidates []int, target int) [][]int {
	result := make([][]int, 0, 0)
	if candidates == nil || target == 0 {
		return result
	}
	sort.Ints(candidates)
	numsLen := len(candidates)
	if numsLen == 0 {
		return result
	}
	if target < candidates[0] {
		return result
	}
	tempResult := make([]int, 0, 0)
	isCaptured := make([]bool, numsLen, numsLen)
	getCombinations0040(&candidates, numsLen, 0, target, &isCaptured, &tempResult, &result)
	return result
}
