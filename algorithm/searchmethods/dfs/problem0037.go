package dfs

func dfsToFill0037(allDone *bool, blankAreaList *[][]int, block *[][][]bool, row *[][]bool, column *[][]bool, board *[][]byte, currentIndex int, maxIndex int) {
	if currentIndex >= maxIndex {
		*allDone = true
	} else {
		y := (*blankAreaList)[currentIndex][0]
		x := (*blankAreaList)[currentIndex][1]
		for num := 1; num <= 9; num++ {
			if !(*block)[y/3][x/3][num] && !(*row)[y][num] && !(*column)[x][num] {
				(*block)[y/3][x/3][num] = true
				(*row)[y][num] = true
				(*column)[x][num] = true
				(*board)[y][x] = byte('0' + num)
				dfsToFill0037(allDone, blankAreaList, block, row, column, board, currentIndex+1, maxIndex)
				if *allDone {
					break
				} else {
					(*block)[y/3][x/3][num] = false
					(*row)[y][num] = false
					(*column)[x][num] = false
					(*board)[y][x] = byte('.')
				}
			}
		}
	}
}

func solveSudoku(board [][]byte) {
	block := make([][][]bool, 3, 3)
	row := make([][]bool, 9, 9)
	column := make([][]bool, 9, 9)
	for i := 0; i < 3; i++ {
		block[i] = make([][]bool, 3, 3)
		for j := 0; j < 3; j++ {
			block[i][j] = make([]bool, 10, 10)
		}
	}
	for i := 0; i < 9; i++ {
		row[i] = make([]bool, 10, 10)
		column[i] = make([]bool, 10, 10)
	}
	blankAreaList := make([][]int, 0, 81)
	curDigit := 0
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if board[y][x] == byte('.') {
				blankAreaList = append(blankAreaList, []int{y, x})
			} else {
				curDigit = int(rune(board[y][x]) - '0')
				block[y/3][x/3][curDigit] = true
				row[y][curDigit] = true
				column[x][curDigit] = true
			}
		}
	}
	allDone := false
	dfsToFill0037(&allDone, &blankAreaList, &block, &row, &column, &board, 0, len(blankAreaList))
}
