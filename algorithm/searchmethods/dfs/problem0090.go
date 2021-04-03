package dfs

import "sort"

func helper0090(nums *[]int, startIndex int, subset *[]int, results *[][]int) {
	subsetCopied := make([]int, len(*subset), len(*subset))
	copy(subsetCopied, *subset)
	*results = append(*results, subsetCopied)
	for i := startIndex; i < len(*nums); i++ {
		if i > startIndex && (*nums)[i-1] == (*nums)[i] {
			continue
		}
		*subset = append(*subset, (*nums)[i])
		helper0090(nums, i+1, subset, results)
		*subset = (*subset)[0 : len(*subset)-1]
	}
}

func subsetsWithDup(nums []int) [][]int {
	numsCopied := make([]int, len(nums), len(nums))
	copy(numsCopied, nums)
	sort.Ints(numsCopied)
	results := make([][]int, 0)
	tempResult := make([]int, 0, len(numsCopied))
	helper0090(&numsCopied, 0, &tempResult, &results)
	return results
}
