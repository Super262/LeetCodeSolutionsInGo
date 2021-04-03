package dfs

func swap0046(nums *[]int, a int, b int) {
	t := (*nums)[a]
	(*nums)[a] = (*nums)[b]
	(*nums)[b] = t
}

func getPermutations0046(nums *[]int, start int, end int, result *[][]int) {
	if start == end-1 {
		temp := append([]int{}, *nums...)
		*result = append(*result, temp)
	} else {
		for i := start; i < end; i++ {
			swap0046(nums, i, start)
			getPermutations0046(nums, start+1, end, result)
			swap0046(nums, i, start)
		}
	}
}

func permute(nums []int) [][]int {
	result := make([][]int, 0)
	if nums == nil {
		return result
	}
	numsLen := len(nums)
	if numsLen == 0 {
		return result
	}
	getPermutations0046(&nums, 0, numsLen, &result)
	return result
}
