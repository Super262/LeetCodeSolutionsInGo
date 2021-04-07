package partition

func minimumTotal(triangle [][]int) int {
	memo := make([][]int, len(triangle), len(triangle))
	maxWidth := len(triangle[len(triangle)-1])
	for i := range memo {
		memo[i] = make([]int, maxWidth, maxWidth)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	return dpToGetMinimal0120(&triangle, &memo, 0, 0)
}

func dpToGetMinimal0120(triangle, memo *[][]int, y, x int) int {
	if y == len(*triangle) {
		return 0
	}
	if (*memo)[y][x] != -1 {
		return (*memo)[y][x]
	}
	currentSum := (*triangle)[y][x] + getMin0120(dpToGetMinimal0120(triangle, memo, y+1, x), dpToGetMinimal0120(triangle, memo, y+1, x+1))
	(*memo)[y][x] = currentSum
	return currentSum
}

func getMin0120(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
