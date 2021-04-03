package dfs

import "sort"

func subsets(nums []int) [][]int {
	results := make([][]int, 0)
	subset := make([]int, 0, len(nums))
	numsCopied := make([]int, len(nums), len(nums))
	copy(numsCopied, nums)
	sort.Ints(numsCopied)
	helper0078(&numsCopied, 0, &subset, &results)
	return results
}

func helper0078(nums *[]int, index int, subset *[]int, results *[][]int) {
	if index == len(*nums) {
		subsetCopied := make([]int, 0)
		subsetCopied = append(subsetCopied, *subset...)
		*results = append(*results, subsetCopied)
		return
	}
	*subset = append(*subset, (*nums)[index])
	helper0078(nums, index+1, subset, results)
	*subset = (*subset)[0 : len(*subset)-1]
	helper0078(nums, index+1, subset, results)
}
