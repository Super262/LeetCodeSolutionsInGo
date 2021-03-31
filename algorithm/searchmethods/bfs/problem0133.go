package bfs

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	visitedCopy := make(map[*Node]*Node)
	q := make([]*Node, 0)
	q = append(q, node)
	visitedCopy[node] = &Node{Val: node.Val}
	var curNode *Node
	for len(q) > 0 {
		curNode = q[0]
		q = q[1:]
		for _, n := range curNode.Neighbors {
			_, existed := visitedCopy[n]
			if !existed {
				visitedCopy[n] = &Node{Val: n.Val}
				q = append(q, n)
			}
			visitedCopy[curNode].Neighbors = append(visitedCopy[curNode].Neighbors, visitedCopy[n])
		}
	}
	return visitedCopy[node]
}
