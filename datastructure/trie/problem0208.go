package trie

import "strings"

type Trie0208 struct {
	sons    [26]*Trie0208
	isWord  bool
	wordVal string
}

/** Initialize your data structure here. */
func Constructor() Trie0208 {
	return Trie0208{}
}

/** Inserts a word into the trie. */
func (this *Trie0208) Insert(word string) {
	charArray := []rune(word)
	arrLen := len(charArray)
	chOrd := 0
	p := this
	for i := 0; i < arrLen; i++ {
		chOrd = int(charArray[i]) - 'a'
		if p.sons[chOrd] == nil {
			p.sons[chOrd] = &Trie0208{}
		}
		p = p.sons[chOrd]
	}
	p.wordVal = word
	p.isWord = true
}

/** Returns if the word is in the trie. */
func (this *Trie0208) Search(word string) bool {
	charArray := []rune(word)
	arrLen := len(charArray)
	chOrd := 0
	p := this
	for i := 0; i < arrLen; i++ {
		chOrd = int(charArray[i]) - 'a'
		if p.sons[chOrd] == nil {
			return false
		}
		p = p.sons[chOrd]
	}
	return strings.Compare(p.wordVal, word) == 0 && p.isWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie0208) StartsWith(prefix string) bool {
	charArray := []rune(prefix)
	arrLen := len(charArray)
	chOrd := 0
	p := this
	for i := 0; i < arrLen; i++ {
		chOrd = int(charArray[i]) - 'a'
		if p.sons[chOrd] == nil {
			return false
		}
		p = p.sons[chOrd]
	}
	return true
}

/**
 * Your Trie0208 object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
