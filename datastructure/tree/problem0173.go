package tree

type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	iter := BSTIterator{make([]*TreeNode, 0, 0)}
	iter.FindMostLeft(root)
	return iter
}

func (this *BSTIterator) Next() int {
	curNode := this.stack[len(this.stack)-1]
	this.stack = this.stack[0 : len(this.stack)-1]
	this.FindMostLeft(curNode.Right)
	return curNode.Val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0
}

func (this *BSTIterator) FindMostLeft(root *TreeNode) {
	for root != nil {
		this.stack = append(this.stack, root)
		root = root.Left
	}
}
