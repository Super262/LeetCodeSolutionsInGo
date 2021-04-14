package problem0251

import "math"

type Vector2D struct {
	vec        [][]int
	currentRow int
	currentCol int
}

func Constructor(vec [][]int) Vector2D {
	iterator := Vector2D{vec: make([][]int, len(vec), len(vec)), currentRow: 0, currentCol: 0}
	copy(iterator.vec, vec)
	return iterator
}

func (vector *Vector2D) Next() int {
	result := math.MinInt32
	if vector.HasNext() {
		result = vector.vec[vector.currentRow][vector.currentCol]
		vector.currentCol++
	}
	return result
}

func (vector *Vector2D) HasNext() bool {
	for vector.currentRow < len(vector.vec) && vector.currentCol >= len(vector.vec[vector.currentRow]) {
		vector.currentRow++
		vector.currentCol = 0
	}
	return vector.currentRow < len(vector.vec)
}
