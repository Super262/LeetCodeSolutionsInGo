package tree

type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	iter := BSTIterator{make([]*TreeNode, 0, 0)}
	for root != nil {
		iter.stack = append(iter.stack, root)
		root = root.Left
	}
	return iter
}

func (this *BSTIterator) Next() int {
	curNode := this.stack[len(this.stack)-1]
	this.stack = this.stack[0 : len(this.stack)-1]
	nextNode := curNode.Right
	for nextNode != nil {
		this.stack = append(this.stack, nextNode)
		nextNode = nextNode.Left
	}
	return curNode.Val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0
}
