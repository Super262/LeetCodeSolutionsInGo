package divideandconquer

func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return -1
	}
	stack := make([]*TreeNode, 0)
	currentNode := root
	for currentNode != nil {
		stack = append(stack, currentNode)
		currentNode = currentNode.Left
	}
	for i := 0; i < k-1; i++ {
		currentNode = stack[len(stack)-1]
		if currentNode.Right == nil {
			currentNode = stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			for len(stack) > 0 && stack[len(stack)-1].Right == currentNode {
				currentNode = stack[len(stack)-1]
				stack = stack[0 : len(stack)-1]
			}
		} else {
			currentNode = currentNode.Right
			for currentNode != nil {
				stack = append(stack, currentNode)
				currentNode = currentNode.Left
			}
		}
	}
	return stack[len(stack)-1].Val
}
