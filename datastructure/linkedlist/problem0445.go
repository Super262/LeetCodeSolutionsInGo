package linkedlist

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l1.Val == 0 {
		return l2
	}
	if l2 == nil || l2.Val == 0 {
		return l1
	}
	len1 := 0
	len2 := 0
	p := l1
	for p != nil {
		len1++
		p = p.Next
	}
	p = l2
	for p != nil {
		len2++
		p = p.Next
	}
	s1 := make([]int, len1, len1)
	s2 := make([]int, len2, len2)
	p = l1
	for i := range s1 {
		s1[i] = p.Val
		p = p.Next
	}
	p = l2
	for i := range s2 {
		s2[i] = p.Val
		p = p.Next
	}
	hasCarry := false
	result := &ListNode{Val: s1[len1-1] + s2[len2-1], Next: nil}
	len1--
	len2--
	if result.Val > 9 {
		result.Val -= 10
		hasCarry = true
	}
	for len1 > 0 && len2 > 0 {
		cur := &ListNode{Val: s1[len1-1] + s2[len2-1], Next: nil}
		len1--
		len2--
		if hasCarry {
			cur.Val++
			hasCarry = false
		}
		if cur.Val > 9 {
			cur.Val -= 10
			hasCarry = true
		}
		cur.Next = result
		result = cur
	}
	for len1 > 0 {
		cur := &ListNode{Val: s1[len1-1], Next: nil}
		len1--
		if hasCarry {
			cur.Val++
			hasCarry = false
		}
		if cur.Val > 9 {
			cur.Val -= 10
			hasCarry = true
		}
		cur.Next = result
		result = cur
	}
	for len2 > 0 {
		cur := &ListNode{Val: s2[len2-1], Next: nil}
		len2--
		if hasCarry {
			cur.Val++
			hasCarry = false
		}
		if cur.Val > 9 {
			cur.Val -= 10
			hasCarry = true
		}
		cur.Next = result
		result = cur
	}
	if hasCarry {
		cur := &ListNode{Val: 1, Next: nil}
		cur.Next = result
		result = cur
	}
	return result
}
