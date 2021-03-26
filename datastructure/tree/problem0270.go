package tree

import "math"

func closestValue(root *TreeNode, target float64) int {
	if root == nil {
		return math.MinInt32
	}
	upper := root
	lower := root
	for root != nil {
		if float64(root.Val) > target {
			upper = root
			root = root.Left
		} else if float64(root.Val) < target {
			lower = root
			root = root.Right
		} else {
			return root.Val
		}
	}
	if math.Abs(target-float64(lower.Val)) < math.Abs(float64(upper.Val)-target) {
		return lower.Val
	} else {
		return upper.Val
	}
}
