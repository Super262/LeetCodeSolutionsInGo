package bfs

func shortestPathBinaryMatrix(grid [][]int) int {
	if grid == nil || len(grid) == 0 || grid[0] == nil || len(grid[0]) == 0 {
		return -1
	}
	maxX := len(grid)
	maxY := len(grid[0])
	if grid[0][0] == 1 || grid[maxX-1][maxY-1] == 1 {
		return -1
	}
	if maxX == 1 && maxY == 1 {
		return 1
	}
	directions := [][]int{{1, -1}, {1, 0}, {1, 1}, {0, -1}, {0, 1}, {-1, -1}, {-1, 0}, {-1, 1}}
	forwardVisited := make([][]bool, maxX, maxX)
	backwardVisited := make([][]bool, maxX, maxX)
	for i := range forwardVisited {
		forwardVisited[i] = make([]bool, maxY, maxY)
		backwardVisited[i] = make([]bool, maxY, maxY)
	}
	forwardQueue := make([]*Point1091, 0, maxX*maxY)
	backwardQueue := make([]*Point1091, 0, maxX*maxY)
	forwardQueue = append(forwardQueue, &Point1091{x: 0, y: 0})
	backwardQueue = append(backwardQueue, &Point1091{x: maxX - 1, y: maxY - 1})
	forwardVisited[0][0] = true
	backwardVisited[maxX-1][maxY-1] = true
	distance := 1
	for len(forwardQueue) > 0 && len(backwardQueue) > 0 {
		distance++
		if extendQueue1091(&forwardQueue, &grid, &directions, &forwardVisited, &backwardVisited) {
			return distance
		}
		distance++
		if extendQueue1091(&backwardQueue, &grid, &directions, &backwardVisited, &forwardVisited) {
			return distance
		}
	}
	return -1
}

func extendQueue1091(queue *[]*Point1091, grid *[][]int, directions *[][]int, currentVisited *[][]bool, oppositeVisited *[][]bool) bool {
	currentSize := len(*queue)
	for i := 0; i < currentSize; i++ {
		p := (*queue)[0]
		*queue = (*queue)[1:]
		for _, d := range *directions {
			nextP := &Point1091{x: p.x + d[0], y: p.y + d[1]}
			if !isAvailable1091(nextP, grid, currentVisited) {
				continue
			}
			if (*oppositeVisited)[nextP.x][nextP.y] {
				return true
			}
			*queue = append(*queue, nextP)
			(*currentVisited)[nextP.x][nextP.y] = true
		}
	}
	return false
}

func isAvailable1091(p *Point1091, grid *[][]int, visited *[][]bool) bool {
	if p.x < 0 || p.x >= len(*grid) {
		return false
	}
	if p.y < 0 || p.y >= len((*grid)[0]) {
		return false
	}
	if (*grid)[p.x][p.y] == 1 {
		return false
	}
	return !(*visited)[p.x][p.y]
}

type Point1091 struct {
	x int
	y int
}
