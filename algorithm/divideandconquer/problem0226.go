package divideandconquer

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	invertedLeft := invertTree(root.Left)
	invertedRight := invertTree(root.Right)
	root.Left = invertedRight
	root.Right = invertedLeft
	return root
}
