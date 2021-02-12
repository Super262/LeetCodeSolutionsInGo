package backtracking

import "strconv"

func getPath0257(prefix *[]rune, root *TreeNode, result *[]string) {
	if root == nil {
		return
	}
	curPreLen := len(*prefix)
	if root.Left == nil && root.Right == nil {
		*prefix = append(*prefix, []rune(strconv.Itoa(root.Val))...)
		*result = append(*result, string(*prefix))
		*prefix = (*prefix)[0:curPreLen]
	}
	if root.Left != nil {
		*prefix = append(*prefix, []rune(strconv.Itoa(root.Val)+"->")...)
		getPath0257(prefix, root.Left, result)
		*prefix = (*prefix)[0:curPreLen]
	}
	if root.Right != nil {
		*prefix = append(*prefix, []rune(strconv.Itoa(root.Val)+"->")...)
		getPath0257(prefix, root.Right, result)
		*prefix = (*prefix)[0:curPreLen]
	}
}

func binaryTreePaths(root *TreeNode) []string {
	result := make([]string, 0)
	if root == nil {
		return result
	}
	prefix := make([]rune, 0, 0)
	getPath0257(&prefix, root, &result)
	return result
}
