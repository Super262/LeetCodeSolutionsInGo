package hashtable

func groupAnagrams(strs []string) [][]string {
	wordGroups := map[string][]string{}
	for _, s := range strs {
		strSorted := countAndSort0049(&s)
		wordGroups[strSorted] = append(wordGroups[strSorted], s)
	}
	results := make([][]string, 0)
	for _, g := range wordGroups {
		results = append(results, g)
	}
	return results
}

func countAndSort0049(s *string) string {
	if s == nil || len(*s) == 0 {
		return ""
	}
	chTable := make([]int, 26, 26)
	for _, ch := range *s {
		chTable[ch-'a']++
	}
	result := make([]byte, 0, len(*s))
	for offset := range chTable {
		for i := 0; i < chTable[offset]; i++ {
			result = append(result, byte('a'+offset))
		}
	}
	return string(result)
}
