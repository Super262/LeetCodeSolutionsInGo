package divideandconquer

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createBST(start int, end int) []*TreeNode {
	if start > end {
		return nil
	} else {
		result := make([]*TreeNode, 0)
		for root := start; root <= end; root++ {
			leftChild := createBST(start, root-1)
			rightChild := createBST(root+1, end)
			if leftChild != nil && rightChild != nil {
				for _, l := range leftChild {
					for _, r := range rightChild {
						rootNode := TreeNode{root, l, r}
						result = append(result, &rootNode)
					}
				}
			} else {
				if leftChild == nil && rightChild == nil {
					rootNode := TreeNode{root, nil, nil}
					result = append(result, &rootNode)
				} else {
					if leftChild != nil {
						for _, l := range leftChild {
							rootNode := TreeNode{root, l, nil}
							result = append(result, &rootNode)
						}
					} else if rightChild != nil {
						for _, r := range rightChild {
							rootNode := TreeNode{root, nil, r}
							result = append(result, &rootNode)
						}
					}
				}
			}
		}
		return result
	}
}

func generateTrees(n int) []*TreeNode {
	return createBST(1, n)
}
