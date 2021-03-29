package trie

type WordDictionary struct {
	son     [26]*WordDictionary
	wordVal string
	isWord  bool
}

/** Initialize your data structure here. */
func Constructor0211() WordDictionary {
	return WordDictionary{}
}

func (this *WordDictionary) AddWord(word string) {
	charArray := []rune(word)
	p := this
	chOrder := 0
	for i := 0; i < len(charArray); i++ {
		chOrder = int(charArray[i]) - int('a')
		if p.son[chOrder] == nil {
			p.son[chOrder] = &WordDictionary{}
		}
		p = p.son[chOrder]
	}
	p.wordVal = word
	p.isWord = true
}

func (this *WordDictionary) Search(word string) bool {
	charArray := []rune(word)
	return this.Find0211(&charArray, 0, len(charArray)-1)
}

func (this *WordDictionary) Find0211(word *[]rune, start int, end int) bool {
	chOrder := 0
	if start == end {
		if (*word)[start] == '.' {
			for i := 0; i < len(this.son); i++ {
				if this.son[i] != nil && this.son[i].isWord {
					return true
				}
			}
		} else {
			chOrder = int((*word)[start]) - 'a'
			return this.son[chOrder] != nil && this.son[chOrder].isWord
		}
	} else {
		if (*word)[start] == '.' {
			for i := 0; i < len(this.son); i++ {
				if this.son[i] != nil && this.son[i].Find0211(word, start+1, end) {
					return true
				}
			}
		} else {
			chOrder = int((*word)[start]) - 'a'
			if this.son[chOrder] == nil {
				return false
			} else {
				return this.son[chOrder].Find0211(word, start+1, end)
			}
		}
	}
	return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
