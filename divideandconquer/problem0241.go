package divideandconquer

import "strconv"

func getResult0241(expression *[]rune, start int, end int) []int {
	result := make([]int, 0)
	i := start
	for i < end {
		if (*expression)[i] == '+' || (*expression)[i] == '-' || (*expression)[i] == '*' {
			break
		}
		i++
	}
	if i >= end {
		res, _ := strconv.Atoi(string((*expression)[start:end]))
		result = append(result, res)
	} else {
		i = start
		var left []int
		var right []int
		for i < end {
			if (*expression)[i] == '+' {
				left = getResult0241(expression, start, i)
				right = getResult0241(expression, i+1, end)
				for _, a := range left {
					for _, b := range right {
						result = append(result, a+b)
					}
				}
			} else if (*expression)[i] == '-' {
				left = getResult0241(expression, start, i)
				right = getResult0241(expression, i+1, end)
				for _, a := range left {
					for _, b := range right {
						result = append(result, a-b)
					}
				}
			} else if (*expression)[i] == '*' {
				left = getResult0241(expression, start, i)
				right = getResult0241(expression, i+1, end)
				for _, a := range left {
					for _, b := range right {
						result = append(result, a*b)
					}
				}
			}
			i++
		}

	}
	return result
}

func diffWaysToCompute(input string) []int {
	expression := []rune(input)
	result := make([]int, 0)
	expLen := len(expression)
	if expLen != 0 {
		result = append(result, getResult0241(&expression, 0, expLen)...)
	}
	return result
}
