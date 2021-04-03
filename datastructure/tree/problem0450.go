package tree

import "math"

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		if root.Left == nil && root.Right == nil {
			root = nil
		} else if root.Left != nil && root.Right != nil {
			root.Val = findMax0450(root.Left)
			root.Left = deleteNode(root.Left, root.Val)
		} else if root.Left != nil {
			root = root.Left
		} else {
			root = root.Right
		}
	} else if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}

func findMax0450(root *TreeNode) int {
	if root == nil {
		return math.MinInt32
	}
	if root.Right == nil {
		return root.Val
	}
	return findMax0450(root.Right)
}
