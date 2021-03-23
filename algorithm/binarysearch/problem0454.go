package binarysearch

func fourSumCount(A []int, B []int, C []int, D []int) int {
	if A == nil || B == nil || C == nil || D == nil {
		return 0
	}
	if len(A) == 0 || len(B) == 0 || len(C) == 0 || len(D) == 0 {
		return 0
	}
	sumOfAB := make(map[int]int)
	for _, a := range A {
		for _, b := range B {
			_, ok := sumOfAB[a+b]
			if ok {
				sumOfAB[a+b] += 1
			} else {
				sumOfAB[a+b] = 1
			}
		}
	}
	result := 0
	for _, c := range C {
		for _, d := range D {
			count, ok := sumOfAB[-c-d]
			if ok {
				result += count
			}
		}
	}
	return result
}
