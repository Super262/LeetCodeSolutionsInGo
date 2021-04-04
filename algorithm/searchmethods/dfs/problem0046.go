package dfs

func permute(nums []int) [][]int {
	results := make([][]int, 0, 0)
	permutation := make([]int, 0, len(nums))
	visited := make([]bool, len(nums), len(nums))
	dfs0046(&nums, &permutation, &visited, &results)
	return results
}

func dfs0046(nums, permutation *[]int, visited *[]bool, results *[][]int) {
	if len(*permutation) == len(*nums) {
		permutationCopied := make([]int, len(*nums), len(*nums))
		copy(permutationCopied, *permutation)
		*results = append(*results, permutationCopied)
		return
	}
	for i := range *nums {
		if (*visited)[i] {
			continue
		}
		(*visited)[i] = true
		*permutation = append(*permutation, (*nums)[i])
		dfs0046(nums, permutation, visited, results)
		*permutation = (*permutation)[0 : len(*permutation)-1]
		(*visited)[i] = false
	}
}
