package ozymaxx

import "testing"

func TestInit(t *testing.T) {
	trieNode := (&TrieNode{}).Init()
	if rune2ChildNode := trieNode.Rune2ChildNode; rune2ChildNode == nil {
		t.Error("the rune -> children node map must be initialized")
	}
	if trieNode.IsEndOfWord {
		t.Error("the IsEndOfWord property must initially be false")
	}
}

func TestMarkEndOfWord(t *testing.T) {
	trieNode := (&TrieNode{}).Init()
	trieNode.MarkEndOfWord()
	if !trieNode.IsEndOfWord {
		t.Error("the IsEndOfWord property must be true")
	}
}

func TestAddChild(t *testing.T) {
	trieNode := (&TrieNode{}).Init()
	t.Run("add a child A", func(t *testing.T) {
		testRune := 'A'
		if childNode := trieNode.AddChild(testRune); childNode == nil {
			t.Errorf("a child node must be created and must be pointed by the rune %v", testRune)
		}
	})
}

func TestGetChild(t *testing.T) {
	trieNode := (&TrieNode{}).Init()
	t.Run("retrieve the child A, which has already been added", func(t *testing.T) {
		testRune := 'A'
		childNode := trieNode.AddChild(testRune)
		if childNodeRetrieved := trieNode.GetChild(testRune); !(childNode == childNodeRetrieved && childNode != nil) {
			t.Error("the child node could not be retrieved properly")
		}
	})
	t.Run("try to retrieve the child B, which has never been added", func(t *testing.T) {
		anotherTestRune := 'B'
		if anotherChildNode := trieNode.GetChild(anotherTestRune); anotherChildNode != nil {
			t.Errorf("the child node pointed by the rune %v must not exist", anotherTestRune)
		}
	})
}
