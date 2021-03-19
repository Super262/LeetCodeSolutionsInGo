package greedystrategy

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	nums1Len := len(nums1)
	nums2Len := len(nums2)
	var part1Len int
	var part2Len int
	if k-nums2Len > 0 {
		part1Len = k - nums2Len
	} else {
		part1Len = 0
	}
	var temp *[]int
	var maxPart1 *[]int
	var maxPart2 *[]int
	result := make([]int, k, k)
	for part1Len <= nums1Len && part1Len <= k {
		part2Len = k - part1Len
		maxPart1 = getMaxSubSeq(&nums1, part1Len)
		maxPart2 = getMaxSubSeq(&nums2, part2Len)
		temp = getLargestSeq0321(maxPart1, maxPart2)
		if compareSeqsFromHead0321(temp, 0, k, &result, 0, k) > 0 {
			copy(result, *temp)
		}
		part1Len++
	}
	return result
}

func getMaxSubSeq(seq *[]int, k int) *[]int {
	seqLen := len(*seq)
	remainsCount := seqLen - k
	stack := make([]int, k, k)
	stackTop := -1
	for i := 0; i < seqLen; i++ {
		for stackTop >= 0 && remainsCount > 0 && stack[stackTop] < (*seq)[i] {
			stackTop--
			remainsCount--
		}
		if stackTop < k-1 {
			stackTop++
			stack[stackTop] = (*seq)[i]
		} else {
			remainsCount--
		}
	}
	return &stack
}

func getLargestSeq0321(seq1 *[]int, seq2 *[]int) *[]int {
	end1 := len(*seq1)
	end2 := len(*seq2)
	result := make([]int, end1+end2, end1+end2)
	p1 := 0
	p2 := 0
	pr := 0
	for p1 < end1 && p2 < end2 {
		if compareSeqsFromHead0321(seq1, p1, end1, seq2, p2, end2) > 0 {
			result[pr] = (*seq1)[p1]
			p1++
		} else {
			result[pr] = (*seq2)[p2]
			p2++
		}
		pr++
	}
	for p1 < end1 {
		result[pr] = (*seq1)[p1]
		p1++
		pr++
	}
	for p2 < end2 {
		result[pr] = (*seq2)[p2]
		p2++
		pr++
	}
	return &result
}

func compareSeqsFromHead0321(seq1 *[]int, start1 int, end1 int, seq2 *[]int, start2 int, end2 int) int {
	l1 := end1 - start1
	l2 := end2 - start2
	commonL := 0
	if l1 <= l2 {
		commonL = l1
	} else {
		commonL = l2
	}
	for i := 0; i < commonL; i++ {
		if (*seq1)[start1+i] != (*seq2)[start2+i] {
			return (*seq1)[start1+i] - (*seq2)[start2+i]
		}
	}
	return l1 - l2
}
