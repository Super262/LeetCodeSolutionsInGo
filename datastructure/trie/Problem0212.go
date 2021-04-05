package trie

func findWords(board [][]byte, words []string) []string {
	trieRoot := BuildTrie0212(&words)
	results := make([]string, 0)
	visited := make([][]bool, len(board), len(board))
	directions := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	for i := range visited {
		visited[i] = make([]bool, len(board[0]), len(board[0]))
	}
	for y := range board {
		for x := range board[y] {
			_, existed := trieRoot.children[board[y][x]]
			if !existed {
				continue
			}
			visited[y][x] = true
			DfsToFindAndPrune(trieRoot.children[board[y][x]], &board, y, x, &directions, &visited, &results)
			visited[y][x] = false
		}
	}
	return results
}

func DfsToFindAndPrune(root *Trie0212, board *[][]byte, startY, startX int, directions *[][]int, visited *[][]bool, results *[]string) bool {
	if root.isWord {
		*results = append(*results, ""+root.wordValue)
		root.isWord = false
		if len(root.children) == 0 {
			return true
		}
	}
	for _, d := range *directions {
		nextY := startY + d[0]
		nextX := startX + d[1]
		if nextY < 0 || nextY >= len(*board) || nextX < 0 || nextX >= len((*board)[0]) || (*visited)[nextY][nextX] {
			continue
		}
		_, existed := root.children[(*board)[nextY][nextX]]
		if !existed {
			continue
		}
		(*visited)[nextY][nextX] = true
		if DfsToFindAndPrune(root.children[(*board)[nextY][nextX]], board, nextY, nextX, directions, visited, results) {
			delete(root.children, (*board)[nextY][nextX])
		}
		(*visited)[nextY][nextX] = false
	}
	return len(root.children) == 0
}

func BuildTrie0212(words *[]string) *Trie0212 {
	root := &Trie0212{isWord: false, wordValue: "", children: make(map[byte]*Trie0212)}
	for _, w := range *words {
		root.AddWord0212(&w)
	}
	return root
}

func (root *Trie0212) AddWord0212(word *string) {
	if root == nil || word == nil || len(*word) == 0 {
		return
	}
	currentNode := root
	for _, ch := range []byte(*word) {
		_, existed := currentNode.children[ch]
		if !existed {
			currentNode.children[ch] = &Trie0212{isWord: false, wordValue: "", children: make(map[byte]*Trie0212)}
		}
		currentNode = currentNode.children[ch]
	}
	currentNode.isWord = true
	currentNode.wordValue = "" + *word
}

type Trie0212 struct {
	isWord    bool
	wordValue string
	children  map[byte]*Trie0212
}
