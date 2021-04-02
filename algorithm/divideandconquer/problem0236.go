package divideandconquer

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || q == nil || p == nil {
		return nil
	}
	pathP := make([]*TreeNode, 0)
	if !findPath0236(root, p, &pathP) {
		return nil
	}
	pathQ := make([]*TreeNode, 0)
	if !findPath0236(root, q, &pathQ) {
		return nil
	}
	i := 1
	for i < len(pathP) && i < len(pathQ) {
		if pathP[i] != pathQ[i] {
			break
		}
		i++
	}
	return pathP[i-1]
}

func findPath0236(root, target *TreeNode, path *[]*TreeNode) bool {
	if root == nil || target == nil {
		return false
	}
	*path = append(*path, root)
	if root == target {
		return true
	}
	if findPath0236(root.Left, target, path) || findPath0236(root.Right, target, path) {
		return true
	}
	*path = (*path)[0 : len(*path)-1]
	return false
}
