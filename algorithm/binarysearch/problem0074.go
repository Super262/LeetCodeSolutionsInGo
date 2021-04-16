package binarysearch

func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 || matrix[0] == nil || len(matrix[0]) == 0 {
		return false
	}
	start := 0
	end := len(matrix)*len(matrix[0]) - 1
	mid := 0
	pivot := 0
	for start+1 < end {
		mid = start + (end-start)/2
		pivot = getNumInMatrix0074(&matrix, mid)
		if target == pivot {
			return true
		} else if target < pivot {
			end = mid
		} else {
			start = mid
		}
	}
	if getNumInMatrix0074(&matrix, start) == target {
		return true
	}
	if getNumInMatrix0074(&matrix, end) == target {
		return true
	}
	return false
}

func getNumInMatrix0074(matrix *[][]int, index int) int {
	return (*matrix)[index/len((*matrix)[0])][index%len((*matrix)[0])]
}
