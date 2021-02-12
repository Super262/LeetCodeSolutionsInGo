package backtracking

import (
	"sort"
)

func getCombinations0039(nums *[]int, numsLen int, startIndex int, target int, tempResult *[]int, result *[][]int) {
	for i := startIndex; i < numsLen; i++ {
		if (*nums)[i] == target {
			*tempResult = append(*tempResult, (*nums)[i])
			*result = append(*result, append([]int{}, *tempResult...))
			*tempResult = (*tempResult)[0 : len(*tempResult)-1]
		} else if (*nums)[i] < target {
			*tempResult = append(*tempResult, (*nums)[i])
			getCombinations0039(nums, numsLen, i, target-(*nums)[i], tempResult, result)
			*tempResult = (*tempResult)[0 : len(*tempResult)-1]
		}
	}
}

func combinationSum(candidates []int, target int) [][]int {
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
	getCombinations0039(&candidates, numsLen, 0, target, &tempResult, &result)
	return result
}
