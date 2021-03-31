package bfs

import "strings"

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordDict := Set0127{}
	wordDict.Constructor(&wordList)
	visited := Set0127{}
	q := make([]string, 0)
	q = append(q, beginWord)
	visited.Add(&beginWord)
	currentLevelSize := 0
	pathLen := 0
	currentNode := ""
	for len(q) > 0 {
		currentLevelSize = len(q)
		pathLen++
		for i := 0; i < currentLevelSize; i++ {
			currentNode = q[0]
			q = q[1:]
			if strings.Compare(currentNode, endWord) == 0 {
				return pathLen
			}
			for _, candidate := range *getCandidates0127(&currentNode) {
				if !wordDict.Contains(&candidate) || visited.Contains(&candidate) {
					continue
				}
				q = append(q, candidate)
				visited.Add(&candidate)
			}
		}
	}
	return 0
}

func getCandidates0127(word *string) *[]string {
	results := make([]string, 0)
	dict := []rune("abcdefghijklmnopqrstuvwxyz")
	target := []rune(*word)
	var prev_val rune
	for si := range target {
		for di := range dict {
			if target[si] == dict[di] {
				continue
			}
			prev_val = target[si]
			target[si] = dict[di]
			results = append(results, string(target))
			target[si] = prev_val
		}
	}
	return &results
}

type Set0127 struct {
	data map[string]int
}

func (s *Set0127) Constructor(wordList *[]string) {
	for i := range *wordList {
		s.Add(&(*wordList)[i])
	}
}

func (s *Set0127) Contains(key *string) bool {
	if s.data == nil {
		s.data = make(map[string]int)
	}
	_, ok := s.data[*key]
	return ok
}

func (s *Set0127) Remove(key *string) {
	if s.data == nil {
		s.data = make(map[string]int)
	}
	delete(s.data, *key)
}

func (s *Set0127) Add(key *string) {
	if s.data == nil {
		s.data = make(map[string]int)
	}
	s.data[*key] = 1
}
