package tree

import "math"

func closestKValues(root *TreeNode, target float64, k int) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	lowerStack := initializeStack(root, target)
	upperStack := append([]*TreeNode{}, lowerStack...)
	if float64(lowerStack[len(lowerStack)-1].Val) > target {
		moveLower(&lowerStack)
	} else {
		moveUpper(&upperStack)
	}
	for i := 0; i < k; i++ {
		if isLowerCloser(&lowerStack, &upperStack, target) {
			result = append(result, lowerStack[len(lowerStack)-1].Val)
			moveLower(&lowerStack)
		} else {
			result = append(result, upperStack[len(upperStack)-1].Val)
			moveUpper(&upperStack)
		}
	}
	return result
}

func isLowerCloser(lowerStack *[]*TreeNode, upperStack *[]*TreeNode, target float64) bool {
	if len(*lowerStack) == 0 {
		return false
	}
	if len(*upperStack) == 0 {
		return true
	}
	return math.Abs(float64((*lowerStack)[len(*lowerStack)-1].Val)-target) < math.Abs(float64((*upperStack)[len(*upperStack)-1].Val)-target)
}

func initializeStack(root *TreeNode, target float64) []*TreeNode {
	stack := make([]*TreeNode, 0)
	for root != nil {
		stack = append(stack, root)
		if float64(root.Val) > target {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return stack
}

func moveUpper(stack *[]*TreeNode) {
	node := (*stack)[len(*stack)-1]
	if node.Right == nil {
		*stack = (*stack)[0 : len(*stack)-1]
		for len(*stack) > 0 && (*stack)[len(*stack)-1].Right == node {
			node = (*stack)[len(*stack)-1]
			*stack = (*stack)[0 : len(*stack)-1]
		}
	} else {
		node = node.Right
		for node != nil {
			*stack = append(*stack, node)
			node = node.Left
		}
	}
}

func moveLower(stack *[]*TreeNode) {
	node := (*stack)[len(*stack)-1]
	if node.Left == nil {
		*stack = (*stack)[0 : len(*stack)-1]
		for len(*stack) > 0 && (*stack)[len(*stack)-1].Left == node {
			node = (*stack)[len(*stack)-1]
			*stack = (*stack)[0 : len(*stack)-1]
		}
	} else {
		node = node.Left
		for node != nil {
			*stack = append(*stack, node)
			node = node.Right
		}
	}
}
