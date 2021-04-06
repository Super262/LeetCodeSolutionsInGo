package problem1429

type FirstUnique struct {
	head             *DLinkNode
	tail             *DLinkNode
	valuesToNode     map[int]*DLinkNode
	duplicatedValues map[int]bool
}

func Constructor(nums []int) FirstUnique {
	fu := FirstUnique{head: &DLinkNode{value: -1}, tail: &DLinkNode{value: -1}, valuesToNode: make(map[int]*DLinkNode), duplicatedValues: make(map[int]bool)}
	fu.head.next = fu.tail
	fu.tail.prev = fu.head
	for _, n := range nums {
		fu.Add(n)
	}
	return fu
}

func (fu *FirstUnique) ShowFirstUnique() int {
	return fu.head.next.value
}

func (fu *FirstUnique) Add(value int) {
	_, existed := fu.duplicatedValues[value]
	if existed {
		return
	}
	_, existed = fu.valuesToNode[value]
	if existed {
		RemoveNode(fu.valuesToNode[value])
		fu.duplicatedValues[value] = true
		delete(fu.valuesToNode, value)
	} else {
		node := &DLinkNode{value: value}
		fu.PushBack(node)
		fu.valuesToNode[value] = node
	}
}

func RemoveNode(node *DLinkNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
	node.prev = nil
	node.next = nil
}

func (fu *FirstUnique) PushBack(node *DLinkNode) {
	node.next = fu.tail
	node.prev = fu.tail.prev
	fu.tail.prev = node
	node.prev.next = node
}

type DLinkNode struct {
	value int
	prev  *DLinkNode
	next  *DLinkNode
}

/**
 * Your FirstUnique object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.ShowFirstUnique();
 * obj.Add(value);
 */
