package doublepointers

func twoSum(numbers []int, target int) []int {
	i := 0
	j := len(numbers) - 1
	var tempSum int
	for i < j {
		tempSum = numbers[i] + numbers[j]
		if tempSum == target {
			break
		} else if tempSum > target {
			j--
		} else {
			i++
		}
	}
	return []int{i + 1, j + 1}
}
