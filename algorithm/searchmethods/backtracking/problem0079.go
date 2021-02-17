package backtracking

import "strings"

func dfsToFind0079(board *[][]byte, startY int, maxY int, startX int, maxX int, direction *[][]int, word *[]rune, startIndex int, endIndex int) bool {
	if startY < 0 || startY >= maxY || startX < 0 || startX >= maxX || startIndex >= endIndex || rune((*board)[startY][startX]) != (*word)[startIndex] {
		return false
	}
	currentValue := (*board)[startY][startX]
	(*board)[startY][startX] = '0'
	if startIndex == endIndex-1 {
		return true
	}
	nextY := 0
	nextX := 0
	for _, d := range *direction {
		nextY = d[0] + startY
		nextX = d[1] + startX
		if dfsToFind0079(board, nextY, maxY, nextX, maxX, direction, word, startIndex+1, endIndex) {
			return true
		}
	}
	(*board)[startY][startX] = currentValue
	return false
}

func exist(board [][]byte, word string) bool {
	if board == nil || strings.Compare(word, "") == 0 {
		return false
	}
	height := len(board)
	if height == 0 {
		return false
	}
	width := len(board[0])
	if width == 0 {
		return false
	}
	wordArray := []rune(word)
	direction := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if rune(board[y][x]) == wordArray[0] {
				if dfsToFind0079(&board, y, height, x, width, &direction, &wordArray, 0, len(wordArray)) {
					return true
				}
			}
		}
	}
	return false
}
