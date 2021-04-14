package problem0341

import "math"

type NestedIterator struct {
	stack []*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	iterator := &NestedIterator{stack: make([]*NestedInteger, len(nestedList), len(nestedList))}
	copy(iterator.stack, nestedList)
	return iterator
}

func (iterator *NestedIterator) Next() int {
	if iterator.HasNext() {
		ans := iterator.stack[0].GetInteger()
		iterator.stack = iterator.stack[1:]
		return ans
	}
	return math.MinInt32
}

func (iterator *NestedIterator) HasNext() bool {
	for len(iterator.stack) > 0 && !iterator.stack[0].IsInteger() {
		front := iterator.stack[0].GetList()
		iterator.stack = iterator.stack[1:]
		if len(front) > 0 {
			tmp := []*NestedInteger{}
			tmp = append(tmp, front...)
			tmp = append(tmp, iterator.stack...)
			iterator.stack = tmp
		}
	}
	return len(iterator.stack) > 0
}

type NestedInteger struct {
}

// IsInteger Return true if this NestedInteger holds a single integer, rather than a nested list.
func (this NestedInteger) IsInteger() bool {
	return false
}

// GetInteger Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (this NestedInteger) GetInteger() int {
	return 0
}

// SetInteger Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {

}

// Add Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (this *NestedInteger) Add(elem NestedInteger) {}

// GetList Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (this NestedInteger) GetList() []*NestedInteger {
	return nil
}
