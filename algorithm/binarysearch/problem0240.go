package binarysearch

func searchMatrixII(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 || matrix[0] == nil || len(matrix[0]) == 0 {
		return false
	}
	row := len(matrix) - 1
	col := 0
	for row >= 0 && row < len(matrix) && col >= 0 && col < len(matrix[0]) {
		if target == matrix[row][col] {
			return true
		} else if target > matrix[row][col] {
			col++
		} else {
			row--
		}
	}
	return false
}
