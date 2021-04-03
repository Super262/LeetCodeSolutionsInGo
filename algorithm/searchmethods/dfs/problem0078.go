package dfs

func getCombinations0078(nums *[]int, startIndex int, endIndex int, tempResult *[]int, result *[][]int) {
	if startIndex < endIndex {
		for i := startIndex; i < endIndex; i++ {
			*tempResult = append(*tempResult, (*nums)[i])
			*result = append(*result, append([]int{}, *tempResult...))
			getCombinations0078(nums, i+1, endIndex, tempResult, result)
			*tempResult = (*tempResult)[0 : len(*tempResult)-1]
		}
	}
}

func subsets(nums []int) [][]int {
	result := make([][]int, 1)
	if nums == nil {
		return result
	}
	numsLen := len(nums)
	if numsLen == 0 {
		return result
	}
	tempResult := make([]int, 0)
	getCombinations0078(&nums, 0, numsLen, &tempResult, &result)
	return result
}
