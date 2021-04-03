package dfs

import "sort"

func getPermutations0047(nums *[]int, result *[][]int, visited *[]bool, tempPath *[]int, numsLen int) {
	if len(*tempPath) == numsLen {
		*result = append(*result, append([]int{}, *tempPath...))
	} else {
		for i := 0; i < numsLen; i++ {
			if (*visited)[i] || (i > 0 && !(*visited)[i-1] && (*nums)[i] == (*nums)[i-1]) {
				continue
			}
			(*visited)[i] = true
			*tempPath = append(*tempPath, (*nums)[i])
			getPermutations0047(nums, result, visited, tempPath, numsLen)
			*tempPath = (*tempPath)[0 : len(*tempPath)-1]
			(*visited)[i] = false
		}
	}
}

func permuteUnique(nums []int) [][]int {
	result := make([][]int, 0)
	if nums == nil {
		return result
	}
	numsLen := len(nums)
	if numsLen == 0 {
		return result
	}
	visited := make([]bool, numsLen, numsLen)
	tempPath := make([]int, 0)
	sort.Ints(nums)
	getPermutations0047(&nums, &result, &visited, &tempPath, numsLen)
	return result
}
