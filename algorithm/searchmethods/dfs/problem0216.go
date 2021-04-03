package dfs

func getCombinations0216(targetSum int, minAddend int, maxAddend int, currentCount int, targetCount int, tempResult *[]int, result *[][]int) {
	if currentCount == targetCount {
		for a := minAddend; a <= maxAddend; a++ {
			if a == targetSum {
				*tempResult = append(*tempResult, a)
				*result = append(*result, append([]int{}, *tempResult...))
				*tempResult = (*tempResult)[0 : len(*tempResult)-1]
				break
			}
		}
	} else {
		for a := minAddend; a <= maxAddend; a++ {
			if a < targetSum {
				*tempResult = append(*tempResult, a)
				getCombinations0216(targetSum-a, a+1, maxAddend, currentCount+1, targetCount, tempResult, result)
				*tempResult = (*tempResult)[0 : len(*tempResult)-1]
			}
		}
	}
}

func combinationSum3(k int, n int) [][]int {
	result := make([][]int, 0, 0)
	tempResult := make([]int, 0, k)
	getCombinations0216(n, 1, 9, 1, k, &tempResult, &result)
	return result
}
