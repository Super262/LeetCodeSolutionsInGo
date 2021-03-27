package unionfind

func numIslands2(m int, n int, positions [][]int) []int {
	father := make([]int, m*n, m*n)
	matrix := make([]int, m*n, m*n)
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	result := make([]int, 0, len(positions))
	tempSize := 0
	rootA := 0
	rootB := 0
	nextY := 0
	nextX := 0
	for _, p := range positions {
		if matrix[pointToInt0305(p[0], p[1], n)] == 1 {
			result = append(result, result[len(result)-1])
			continue
		}
		tempSize = 1
		if len(result) > 0 {
			tempSize += result[len(result)-1]
		}
		connect0305(&father, pointToInt0305(p[0], p[1], n), pointToInt0305(p[0], p[1], n))
		matrix[pointToInt0305(p[0], p[1], n)] = 1
		for _, d := range directions {
			nextY = p[0] + d[0]
			nextX = p[1] + d[1]
			if nextY >= m || nextY < 0 || nextX < 0 || nextX >= n || matrix[pointToInt0305(nextY, nextX, n)] == 0 {
				continue
			}
			rootA = findAndCompress0305(&father, pointToInt0305(p[0], p[1], n))
			rootB = findAndCompress0305(&father, pointToInt0305(nextY, nextX, n))
			if rootA == rootB {
				continue
			}
			connect0305(&father, rootA, rootB)
			tempSize--
		}
		result = append(result, tempSize)
	}
	return result
}

func connect0305(father *[]int, aIndex int, bIndex int) {
	(*father)[bIndex] = aIndex
}

func findAndCompress0305(father *[]int, nodeIndex int) int {
	path := make([]int, 0)
	for (*father)[nodeIndex] != nodeIndex {
		path = append(path, nodeIndex)
		nodeIndex = (*father)[nodeIndex]
	}
	for _, p := range path {
		(*father)[p] = nodeIndex
	}
	return nodeIndex
}

func pointToInt0305(y int, x int, n int) int {
	return y*n + x
}
