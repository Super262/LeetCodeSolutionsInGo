package dfs

import (
	"math"
	"strings"
)

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	wordDict := make(map[string]bool)
	wordDict[beginWord] = true
	for _, w := range wordList {
		wordDict[w] = true
	}
	_, endExisted := wordDict[endWord]
	if !endExisted {
		return [][]string{}
	}
	path := make([]string, 0)
	path = append(path, beginWord)
	results := make([][]string, 0)
	indices := buildIndices0126(&wordDict)
	distance := bfsGetDistanceFromEnd0126(&endWord, indices, &wordDict)
	dfsGetPathFromStart0126(&beginWord, &endWord, indices, distance, &path, &results)
	return results
}

func dfsGetPathFromStart0126(beginWord, endWord *string, indices *map[string]map[string]bool, distance *map[string]int, path *[]string, results *[][]string) {
	if strings.Compare(*beginWord, *endWord) == 0 {
		pathCopied := make([]string, len(*path), len(*path))
		copy(pathCopied, *path)
		*results = append(*results, pathCopied)
		return
	}
	for _, n := range *getNeighbors0126(beginWord, indices) {
		if (*distance)[n] != (*distance)[*beginWord]-1 {
			continue
		}
		*path = append(*path, n)
		dfsGetPathFromStart0126(&n, endWord, indices, distance, path, results)
		*path = (*path)[0 : len(*path)-1]
	}
}

func bfsGetDistanceFromEnd0126(end *string, indices *map[string]map[string]bool, wordDict *map[string]bool) *map[string]int {
	distance := make(map[string]int)
	distance[*end] = 0
	queue := make([]string, 0)
	queue = append(queue, *end)
	for len(queue) > 0 {
		word := queue[0]
		queue = queue[1:]
		for _, n := range *getNeighbors0126(&word, indices) {
			_, ok := distance[n]
			if ok {
				continue
			}
			distance[n] = distance[word] + 1
			queue = append(queue, n)
		}
	}
	for w := range *wordDict {
		_, ok := distance[w]
		if ok {
			continue
		}
		distance[w] = math.MaxInt32
	}
	return &distance
}

func buildIndices0126(wordDict *map[string]bool) *map[string]map[string]bool {
	indices := make(map[string]map[string]bool)
	for w := range *wordDict {
		for i := range w {
			key := w[0:i] + "%" + w[i+1:]
			_, existed := indices[key]
			if !existed {
				indices[key] = make(map[string]bool)
			}
			indices[key][w] = true
		}
	}
	return &indices
}

func getNeighbors0126(word *string, indices *map[string]map[string]bool) *[]string {
	neighbors := make([]string, 0)
	for i := range *word {
		key := (*word)[0:i] + "%" + (*word)[i+1:]
		for n := range (*indices)[key] {
			neighbors = append(neighbors, n)
		}
	}
	return &neighbors
}
