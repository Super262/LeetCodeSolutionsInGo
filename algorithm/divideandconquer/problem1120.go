package divideandconquer

import "math"

func maximumAverageSubtree(root *TreeNode) float64 {
	_, _, result := helper1120(root)
	return result
}

func helper1120(root *TreeNode) (int, int, float64) {
	if root == nil {
		return 0, 0, -1
	}
	leftNodesCount, leftValuesSum, leftAvgMax := helper1120(root.Left)
	rightNodesCount, rightValuesSum, rightAvgMax := helper1120(root.Right)
	rootNodesCount := 1 + leftNodesCount + rightNodesCount
	rootValuesSum := root.Val + leftValuesSum + rightValuesSum
	rootAvg := float64(rootValuesSum) / float64(rootNodesCount)
	rootAvgMax := math.Max(rootAvg, math.Max(leftAvgMax, rightAvgMax))
	return rootNodesCount, rootValuesSum, rootAvgMax
}
