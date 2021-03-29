package trie

import "strings"

type Trie struct {
	sons    [26]*Trie
	isWord  bool
	wordVal string
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	charArray := []rune(word)
	arrLen := len(charArray)
	chOrd := 0
	p := this
	for i := 0; i < arrLen; i++ {
		chOrd = int(charArray[i]) - 'a'
		if p.sons[chOrd] == nil {
			p.sons[chOrd] = &Trie{}
		}
		p = p.sons[chOrd]
	}
	p.wordVal = word
	p.isWord = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
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
func (this *Trie) StartsWith(prefix string) bool {
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
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
