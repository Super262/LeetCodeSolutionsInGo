package bfs

func countComponents(n int, edges [][]int) int {
	if edges == nil || len(edges) == 0 {
		return n
	}
	neighbors := make(map[int][]int)
	for i := 0; i < n; i++ {
		neighbors[i] = make([]int, 0)
	}
	for _, e := range edges {
		neighbors[e[0]] = append(neighbors[e[0]], e[1])
		neighbors[e[1]] = append(neighbors[e[1]], e[0])
	}
	visited := make([]bool, n, n)
	count := 0
	for i := 0; i < n; i++ {
		if (visited)[i] {
			continue
		}
		bfs(i, &neighbors, &visited)
		count++
	}
	return count
}

func bfs(root int, neighbors *map[int][]int, visited *[]bool) {
	queue := make([]int, 0, len(*visited))
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		for _, n := range (*neighbors)[node] {
			if (*visited)[n] {
				continue
			}
			queue = append(queue, n)
			(*visited)[n] = true
		}
	}
}
