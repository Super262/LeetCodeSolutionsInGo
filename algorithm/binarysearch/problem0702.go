package binarysearch

type ArrayReader struct {
}

func (this *ArrayReader) get(index int) int {
	return index
}

func search(reader ArrayReader, target int) int {
	kth := 1
	for reader.get(kth-1) < target {
		kth *= 2
	}
	start := 0
	end := kth - 1
	mid := 0
	for start+1 < end {
		mid = start + (end-start)/2
		if reader.get(mid) < target {
			start = mid
		} else {
			end = mid
		}
	}
	if reader.get(start) == target {
		return start
	}
	if reader.get(end) == target {
		return end
	}
	return -1
}
