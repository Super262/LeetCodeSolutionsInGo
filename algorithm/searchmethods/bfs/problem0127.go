package bfs

import "strings"

func createMap(wordList []string) map[int][]int {
	result := make(map[int][]int)
	wordsArrays := make([][]rune, len(wordList), len(wordList))
	for i := 0; i < len(wordList); i++ {
		result[i] = make([]int, 0)
		wordsArrays[i] = []rune(wordList[i])
	}
	dis := 0
	for p1 := 0; p1 < len(wordList); p1++ {
		for p2 := p1 + 1; p2 < len(wordList); p2++ {
			dis = 0
			if len(wordsArrays[p1]) == len(wordsArrays[p2]) {
				for i := 0; i < len(wordsArrays[p1]); i++ {
					if wordsArrays[p1][i] != wordsArrays[p2][i] {
						dis++
					}
				}
				if dis == 1 {
					result[p1] = append(result[p1], p2)
					result[p2] = append(result[p2], p1)
				}
			}
		}
	}
	return result
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordList = append(wordList, beginWord)
	wordMap := createMap(wordList)
	q := make([]int, 1)
	q[0] = len(wordList) - 1
	visited := make([]bool, len(wordList), len(wordList))
	count := 0
	curLevelSize := 1
	nextLevelSize := 0
	for len(q) > 0 {
		count++
		for curLevelSize > 0 {
			curNode := q[0]
			q = q[1:]
			curLevelSize--
			if !visited[curNode] {
				visited[curNode] = true
				if strings.Compare(wordList[curNode], endWord) == 0 {
					return count
				} else {
					nextLevelSize += len(wordMap[curNode])
					q = append(q, wordMap[curNode]...)
				}
			}
		}
		curLevelSize = nextLevelSize
		nextLevelSize = 0
	}
	return 0
}
