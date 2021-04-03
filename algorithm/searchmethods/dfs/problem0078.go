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
	subsetCopied := make([]int, len(*subset), len(*subset))
	copy(subsetCopied, *subset)
	*results = append(*results, subsetCopied)
	for i := index; i < len(*nums); i++ {
		*subset = append(*subset, (*nums)[i])
		helper0078(nums, i+1, subset, results)
		*subset = (*subset)[0 : len(*subset)-1]
	}
}
