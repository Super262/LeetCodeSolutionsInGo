package bfs

func sequenceReconstruction(org []int, seqs [][]int) bool {
	if (org == nil || len(org) == 0) && (seqs == nil || len(seqs) == 0 || len(seqs[0]) == 0) {
		return true
	}
	if org == nil || len(org) == 0 {
		return false
	}
	if seqs == nil || len(seqs) == 0 || len(seqs[0]) == 0 {
		return false
	}
	inDegree := make(map[int]int)
	neighbors := make(map[int][]int)
	for _, edge := range seqs {
		for _, v := range edge {
			inDegree[v] = 0
			neighbors[v] = make([]int, 0)
		}
	}
	for _, edge := range seqs {
		for i, v := range edge {
			if i == 0 {
				continue
			}
			inDegree[v]++
			neighbors[edge[i-1]] = append(neighbors[edge[i-1]], v)
		}
	}
	queue := make([]int, 0)
	for v, c := range inDegree {
		if c == 0 {
			queue = append(queue, v)
		}
	}
	curNode := 0
	order := make([]int, 0)
	for len(queue) > 0 {
		if len(queue) > 1 {
			return false
		}
		curNode = queue[0]
		queue = queue[1:]
		order = append(order, curNode)
		for _, n := range neighbors[curNode] {
			inDegree[n]--
			if inDegree[n] == 0 {
				queue = append(queue, n)
			}
		}
	}
	if len(order) != len(inDegree) || len(order) != len(org) {
		return false
	}
	for i := range order {
		if order[i] != org[i] {
			return false
		}
	}
	return true
}
