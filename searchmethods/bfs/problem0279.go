package bfs

func generateSquares(n int) []int {
	result := make([]int, 0)
	tempRes := 0
	for i := 1; i <= n; i++ {
		tempRes = i * i
		if tempRes <= n {
			result = append(result, tempRes)
		} else {
			break
		}
	}
	return result
}

func numSquares(n int) int {
	squaresPrev := generateSquares(n)
	visited := make([]bool, n+1, n+1)
	q := make([]int, 0, n)
	count := 0
	q = append(q, n)
	curNode := 0
	nextNode := 0
	curLevelSize := 1
	nextLevelSize := 0
	for len(q) != 0 {
		count++
		for curLevelSize > 0 {
			curNode = q[0]
			q = q[1:]
			curLevelSize--
			if !visited[curNode] {
				visited[curNode] = true
				for _, sq := range squaresPrev {
					if sq > curNode {
						break
					} else {
						nextNode = curNode - sq
						if nextNode == 0 {
							return count
						} else {
							q = append(q, nextNode)
							nextLevelSize++
						}
					}
				}
			}
		}
		curLevelSize = nextLevelSize
		nextLevelSize = 0
	}
	return n
}
