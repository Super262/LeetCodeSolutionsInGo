package tree

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	currentNode := root
	for true {
		if currentNode.Val == val {
			break
		} else if currentNode.Val > val {
			if currentNode.Left == nil {
				currentNode.Left = &TreeNode{Val: val}
				break
			}
			currentNode = currentNode.Left
		} else {
			if currentNode.Right == nil {
				currentNode.Right = &TreeNode{Val: val}
				break
			}
			currentNode = currentNode.Right
		}
	}
	return root
}
