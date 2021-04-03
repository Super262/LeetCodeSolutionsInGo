package divideandconquer

func flatten(root *TreeNode) {
	flattenAndReturnLastNode(root)
}

func flattenAndReturnLastNode(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	leftLastNode := flattenAndReturnLastNode(root.Left)
	rightLastNode := flattenAndReturnLastNode(root.Right)
	if leftLastNode != nil {
		leftLastNode.Right = root.Right
		root.Right = root.Left
		root.Left = nil
	}
	if rightLastNode != nil {
		return rightLastNode
	}
	if leftLastNode != nil {
		return leftLastNode
	}
	return root
}
