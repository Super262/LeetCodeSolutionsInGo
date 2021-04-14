package matching

func isMatch(s string, p string) bool {
	memo := make([][]bool, len(p), len(p))
	visited := make([][]bool, len(p), len(p))
	for i := 0; i < len(p); i++ {
		memo[i] = make([]bool, len(s), len(s))
		visited[i] = make([]bool, len(s), len(s))
	}
	return helper0044(&s, 0, &p, 0, &memo, &visited)
}

func helper0044(s *string, sStart int, p *string, pStart int, memo *[][]bool, visited *[][]bool) bool {
	if pStart == len(*p) {
		return sStart == len(*s)
	}
	if sStart == len(*s) {
		return allStars0044(p, pStart)
	}
	if (*visited)[pStart][sStart] {
		return (*memo)[pStart][sStart]
	}
	currentResult := false
	if (*p)[pStart] == '*' {
		currentResult = helper0044(s, sStart+1, p, pStart, memo, visited) || helper0044(s, sStart, p, pStart+1, memo, visited)
	} else {
		currentResult = twoCharsMatch0044((*s)[sStart], (*p)[pStart]) && helper0044(s, sStart+1, p, pStart+1, memo, visited)
	}
	(*visited)[pStart][sStart] = true
	(*memo)[pStart][sStart] = currentResult
	return currentResult
}

func allStars0044(p *string, pIndex int) bool {
	for i := pIndex; i < len(*p); i++ {
		if (*p)[i] != '*' {
			return false
		}
	}
	return true
}

func twoCharsMatch0044(s byte, p byte) bool {
	return s == p || p == '?'
}
