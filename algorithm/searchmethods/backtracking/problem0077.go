package backtracking

func getCombinations0077(n int, start int, targetLen int, result *[][]int, tempPath *[]int) {
	if targetLen == 0 {
		*result = append(*result, append([]int{}, *tempPath...))
	} else {
		for i := start; i <= n-targetLen+1; i++ {
			*tempPath = append(*tempPath, i)
			getCombinations0077(n, i+1, targetLen-1, result, tempPath)
			*tempPath = (*tempPath)[0 : len(*tempPath)-1]
		}
	}
}

func combine(n int, k int) [][]int {
	result := make([][]int, 0)
	tempPath := make([]int, 0, k)
	getCombinations0077(n, 1, k, &result, &tempPath)
	return result
}
