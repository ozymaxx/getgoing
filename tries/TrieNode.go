package ozymaxx

// TrieNode - a structure that represents a trie node
type TrieNode struct {
	Rune2ChildNode map[rune]*TrieNode
	IsTerminal     bool
}

// Init - a method initialising the properties of a TrieNode
func (trieNode *TrieNode) Init() *TrieNode {
	trieNode.Rune2ChildNode = make(map[rune]*TrieNode)
	trieNode.IsTerminal = false
	return trieNode
}

// MarkAsTerminal - a method setting IsTerminal to true
func (trieNode *TrieNode) MarkAsTerminal() {
	trieNode.IsTerminal = true
}

// GetChild - a method that returns the children node pointed by a string
// consisting of a rune
func (trieNode *TrieNode) GetChild(rn rune) *TrieNode {
	if val, ok := trieNode.Rune2ChildNode[rn]; ok {
		return val
	}
	return nil
}

// AddAndGetChild - a method that adds a child pointed by a rune
func (trieNode *TrieNode) AddAndGetChild(rn rune) *TrieNode {
	if val, ok := trieNode.Rune2ChildNode[rn]; ok {
		return val
	}
	newChild := (&TrieNode{}).Init()
	trieNode.Rune2ChildNode[rn] = newChild
	return newChild
}
