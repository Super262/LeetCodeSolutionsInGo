package trie

func wordSquares(words []string) [][]string {
	result := make([][]string, 0)
	if words == nil {
		return result
	}
	wordsArray := make([][]rune, len(words), len(words))
	for i := 0; i < len(words); i++ {
		wordsArray[i] = []rune(words[i])
	}
	var temp [][]rune
	maxStep := len(wordsArray[0])
	trieRoot := buildTrie0425(&wordsArray)
	for i := range wordsArray {
		temp = make([][]rune, 0)
		temp = append(temp, wordsArray[i])
		backTracking0425(&wordsArray, 1, maxStep, trieRoot, &temp, &result)
	}
	return result
}

func buildTrie0425(words *[][]rune) *Node0425 {
	root := &Node0425{}
	var currentP *Node0425
	var chOrder int
	for wi := 0; wi < len(*words); wi++ {
		currentP = root
		currentP.wordsIndice = append(currentP.wordsIndice, wi)
		for ci := 0; ci < len((*words)[wi]); ci++ {
			chOrder = int((*words)[wi][ci]) - 'a'
			if currentP.sons[chOrder] == nil {
				currentP.sons[chOrder] = &Node0425{}
			}
			currentP = currentP.sons[chOrder]
			currentP.wordsIndice = append(currentP.wordsIndice, wi)
		}
	}
	return root
}

func backTracking0425(words *[][]rune, currentStep int, maxStep int, root *Node0425, temp *[][]rune, result *[][]string) {
	if currentStep == maxStep {
		newOne := make([]string, 0, len(*temp))
		for i := range *temp {
			newOne = append(newOne, string((*temp)[i]))
		}
		*result = append(*result, newOne)
		return
	}
	prefix := make([]rune, 0, len(*temp))
	for i := range *temp {
		prefix = append(prefix, (*temp)[i][currentStep])
	}
	wordsSet := getWordsWithPrefix0425(root, words, &prefix)
	for i := range *wordsSet {
		*temp = append(*temp, (*wordsSet)[i])
		backTracking0425(words, currentStep+1, maxStep, root, temp, result)
		*temp = (*temp)[0 : len(*temp)-1]
	}
}

func getWordsWithPrefix0425(root *Node0425, words *[][]rune, prefix *[]rune) *[][]rune {
	currentP := root
	var chOrder int
	for ci := 0; ci < len(*prefix); ci++ {
		chOrder = int((*prefix)[ci]) - 'a'
		if currentP.sons[chOrder] == nil {
			return &[][]rune{}
		}
		currentP = currentP.sons[chOrder]
	}
	result := make([][]rune, 0, len(currentP.wordsIndice))
	for _, index := range currentP.wordsIndice {
		result = append(result, (*words)[index])
	}
	return &result
}

type Node0425 struct {
	sons        [26]*Node0425
	wordsIndice []int
}
