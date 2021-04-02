package divideandconquer

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || q == nil || p == nil {
		return nil
	}
	if p == q {
		return p
	}
	if root == p || q == root {
		return root
	}
	leftResult := lowestCommonAncestor(root.Left, p, q)
	rightResult := lowestCommonAncestor(root.Right, p, q)
	if leftResult != nil && rightResult != nil {
		return root
	}
	if leftResult != nil {
		return leftResult
	}
	return rightResult
}
