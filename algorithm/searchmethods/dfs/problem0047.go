package dfs

import "sort"

func permuteUnique(nums []int) [][]int {
	numsCopied := make([]int, len(nums), len(nums))
	permutation := make([]int, 0, len(nums))
	visited := make([]bool, len(nums), len(nums))
	copy(numsCopied, nums)
	results := make([][]int, 0)
	sort.Ints(numsCopied)
	dfs0047(&numsCopied, &permutation, &visited, &results)
	return results
}

func dfs0047(nums, permutation *[]int, visited *[]bool, results *[][]int) {
	if len(*permutation) == len(*nums) {
		permutationCopied := make([]int, len(*permutation), len(*permutation))
		copy(permutationCopied, *permutation)
		*results = append(*results, permutationCopied)
		return
	}
	for i := range *nums {
		if (*visited)[i] || (i > 0 && (*nums)[i-1] == (*nums)[i] && !(*visited)[i-1]) {
			continue
		}
		(*visited)[i] = true
		*permutation = append(*permutation, (*nums)[i])
		dfs0047(nums, permutation, visited, results)
		*permutation = (*permutation)[0 : len(*permutation)-1]
		(*visited)[i] = false
	}
}
