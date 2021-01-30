package ozymaxx

import "unicode/utf8"

// Trie - a struct that represents a trie structure where
// a rune points to a child node
type Trie struct {
	Root *TrieNode
}

// Init - a method of the Trie class initialising the root node
func (trie *Trie) Init() *Trie {
	trie.Root = (&TrieNode{}).Init()
	trie.Root.MarkAsTerminal()
	return trie
}

// ContainsWord - a Trie method checking if a given word is
// contained in the trie or not
func (trie *Trie) ContainsWord(word string) bool {
	currentNode := trie.Root
	for _, ch := range word {
		currentNodeChild := currentNode.GetChild(ch)
		if currentNodeChild == nil {
			return false
		}
		currentNode = currentNodeChild
	}
	return currentNode.IsTerminal
}

// AddWord - a Trie method that adds a given word
func (trie *Trie) AddWord(word string) {
	i := 0
	currentNode := trie.Root
	for _, ch := range word {
		currentNodeChild := currentNode.GetChild(ch)
		if currentNodeChild == nil {
			break
		}
		i++
		currentNode = currentNodeChild
	}
	numRunes := utf8.RuneCountInString(word)
	if i == numRunes {
		currentNode.MarkAsTerminal()
		return
	}
	wordAsRuneArray := []rune(word)
	for ; i < numRunes; i++ {
		currentNode = currentNode.AddChild(wordAsRuneArray[i])
	}
	currentNode.MarkAsTerminal()
}
