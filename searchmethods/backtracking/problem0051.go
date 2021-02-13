package backtracking

func dfsToFill0051(board *[][]rune, colUsed *[]bool, dia45Used *[]bool, dia135Used *[]bool, result *[][]string, curRow int, maxRow int) {
	if curRow == maxRow {
		tempResult := make([]string, 0, maxRow)
		for _, r := range *board {
			tempResult = append(tempResult, string(r))
		}
		*result = append(*result, tempResult)
	} else {
		for col := 0; col < maxRow; col++ {
			if !(*colUsed)[col] && !(*dia45Used)[curRow+col] && !(*dia135Used)[maxRow-1-(curRow-col)] {
				(*colUsed)[col] = true
				(*dia45Used)[curRow+col] = true
				(*dia135Used)[maxRow-1-(curRow-col)] = true
				(*board)[curRow][col] = 'Q'
				dfsToFill0051(board, colUsed, dia45Used, dia135Used, result, curRow+1, maxRow)
				(*colUsed)[col] = false
				(*dia45Used)[curRow+col] = false
				(*dia135Used)[maxRow-1-(curRow-col)] = false
				(*board)[curRow][col] = '.'
			}
		}
	}
}

func solveNQueens(n int) [][]string {
	board := make([][]rune, n, n)
	for i := 0; i < n; i++ {
		board[i] = make([]rune, n, n)
		for j := 0; j < n; j++ {
			board[i][j] = '.'
		}
	}
	colUsed := make([]bool, n, n)
	dia45Used := make([]bool, 2*n-1, 2*n-1)
	dia135Used := make([]bool, 2*n-1, 2*n-1)
	result := make([][]string, 0, n)
	dfsToFill0051(&board, &colUsed, &dia45Used, &dia135Used, &result, 0, n)
	return result
}
