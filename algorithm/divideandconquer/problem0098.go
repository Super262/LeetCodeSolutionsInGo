package divideandconquer

import "math"

func isValidBST(root *TreeNode) bool {
	// 根据题目输入，注意边界值的选取：-2^31 <= Node.val <= 2^31 - 1
	return helper0098(root, math.MinInt64, math.MaxInt64)
}

func helper0098(root *TreeNode, lowerBound, upperBound int64) bool {
	if root == nil {
		return true
	}
	if int64(root.Val) >= upperBound || int64(root.Val) <= lowerBound {
		return false
	}
	return helper0098(root.Left, lowerBound, int64(root.Val)) && helper0098(root.Right, int64(root.Val), upperBound)
}
