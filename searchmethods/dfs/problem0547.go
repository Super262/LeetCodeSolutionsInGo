package dfs

func dfsExploreArea0547(grid *[][]int, i int, N int, hasVisited *[]bool) {
	(*hasVisited)[i] = true
	for j := 0; j < N; j++ {
		if (*grid)[i][j] == 1 && !(*hasVisited)[j] {
			dfsExploreArea0547(grid, j, N, hasVisited)
		}
	}
}

func findCircleNum(isConnected [][]int) int {
	if isConnected == nil {
		return 0
	}
	height := len(isConnected)
	if height == 0 || height != len(isConnected[0]) {
		return 0
	}
	N := height
	hasVisited := make([]bool, N, N)
	count := 0
	for i := 0; i < N; i++ {
		if !hasVisited[i] {
			dfsExploreArea0547(&isConnected, i, N, &hasVisited)
			count++
		}
	}
	return count
}
