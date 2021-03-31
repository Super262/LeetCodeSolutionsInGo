package bfs

func findOrder(numCourses int, prerequisites [][]int) []int {
	inDegree := make([]int, numCourses, numCourses)
	neighbors := make(map[int][]int)
	for i := 0; i < numCourses; i++ {
		neighbors[i] = make([]int, 0)
	}
	for _, e := range prerequisites {
		neighbors[e[1]] = append(neighbors[e[1]], e[0])
		inDegree[e[0]]++
	}
	order := make([]int, 0, numCourses)
	queue := make([]int, 0, numCourses)
	for i := range inDegree {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	curNode := 0
	for len(queue) > 0 {
		curNode = queue[0]
		queue = queue[1:]
		order = append(order, curNode)
		for _, n := range neighbors[curNode] {
			inDegree[n] -= 1
			if inDegree[n] == 0 {
				queue = append(queue, n)
			}
		}
	}
	if len(order) == numCourses {
		return order
	}
	return []int{}
}
