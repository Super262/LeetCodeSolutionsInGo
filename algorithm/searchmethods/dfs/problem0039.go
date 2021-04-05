package dfs

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	candidatesCopied := make([]int, len(candidates), len(candidates))
	copy(candidatesCopied, candidates)
	sort.Ints(candidatesCopied)
	results := make([][]int, 0)
	combination := make([]int, 0, len(candidatesCopied))
	dfs0039(&candidatesCopied, &combination, 0, target, &results)
	return results
}

func dfs0039(candidates, combination *[]int, startIndex, target int, results *[][]int) {
	if target < 0 {
		return
	}
	if target == 0 {
		combinationCopied := make([]int, len(*combination), len(*combination))
		copy(combinationCopied, *combination)
		*results = append(*results, combinationCopied)
		return
	}
	for i := startIndex; i < len(*candidates); i++ {
		if target < (*candidates)[i] {
			return
		}
		*combination = append(*combination, (*candidates)[i])
		dfs0039(candidates, combination, i, target-(*candidates)[i], results)
		*combination = (*combination)[0 : len(*combination)-1]
	}
}
