package unionfind

func validTree(n int, edges [][]int) bool {
	if len(edges) != n-1 {
		return false
	}
	father := make([]int, n, n)
	for i := range father {
		father[i] = i
	}
	connectedCount := n
	rootA := 0
	rootB := 0
	for _, e := range edges {
		rootA = findAndCompress0261(&father, e[0])
		rootB = findAndCompress0261(&father, e[1])
		if rootB == rootA {
			continue
		}
		father[rootB] = rootA
		connectedCount--
	}
	return connectedCount == 1
}

func findAndCompress0261(father *[]int, nodeIndex int) int {
	path := make([]int, 0)
	for (*father)[nodeIndex] != nodeIndex {
		path = append(path, nodeIndex)
		nodeIndex = (*father)[nodeIndex]
	}
	for _, v := range path {
		(*father)[v] = nodeIndex
	}
	return nodeIndex
}
