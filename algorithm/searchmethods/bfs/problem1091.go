package bfs

import "container/list"

func shortestPathBinaryMatrix(grid [][]int) int {
	if grid == nil {
		return -1
	}
	height := len(grid)
	if height == 0 {
		return -1
	}
	width := len(grid[0])
	if width == 0 {
		return -1
	}
	direction := [][]int{{1, -1}, {1, 0}, {1, 1}, {0, -1}, {0, 1}, {-1, -1}, {-1, 0}, {-1, 1}}
	pathLen := 0
	curlevelSize := 1
	nextlevelSize := 0
	q := list.New()
	q.PushBack([]int{0, 0})
	var curPoint []int
	var nextY, nextX int
	for q.Len() > 0 {
		pathLen++
		for curlevelSize > 0 {
			curPoint = q.Front().Value.([]int)
			q.Remove(q.Front())
			if grid[curPoint[0]][curPoint[1]] == 0 {
				grid[curPoint[0]][curPoint[1]] = 1
				if curPoint[0] == height-1 && curPoint[1] == width-1 {
					return pathLen
				} else {
					for _, d := range direction {
						nextY = curPoint[0] + d[0]
						nextX = curPoint[1] + d[1]
						if nextY >= 0 && nextY < height && nextX >= 0 && nextX < width {
							q.PushBack([]int{nextY, nextX})
							nextlevelSize++
						}
					}
				}
			}
			curlevelSize--
		}
		curlevelSize = nextlevelSize
		nextlevelSize = 0
	}
	return -1
}
