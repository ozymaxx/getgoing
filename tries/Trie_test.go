package ozymaxx

import "testing"

func TestTrie(t *testing.T) {
	trie := (&Trie{}).Init()
	t.Run("empty value already exists", func(t *testing.T) {
		if !trie.ContainsWord("") {
			t.Error("empty word must exist in the trie")
		}
	})
	t.Run("when 'a' is added, it must exist in the trie", func(t *testing.T) {
		word := "a"
		trie.AddWord(word)
		if !trie.ContainsWord(word) {
			t.Errorf("the word '%s' must exist in the trie", word)
		}
	})
	t.Run("when 'abc' is added, 'abc' must exist in the trie but 'ab' mustn't", func(t *testing.T) {
		word := "abc"
		trie.AddWord(word)
		if !trie.ContainsWord(word) {
			t.Errorf("the word '%s' must exist in the trie", word)
		}
		anotherWord := "ab"
		if trie.ContainsWord(anotherWord) {
			t.Errorf("the word '%s' must not exist in the trie", anotherWord)
		}
	})
	t.Run("when 'ozan' is added, 'ozan' must exist in the trie but 'ozanc' mustn't", func(t *testing.T) {
		word := "ozan"
		trie.AddWord(word)
		if !trie.ContainsWord(word) {
			t.Errorf("the word '%s' must exist in the trie", word)
		}
		anotherWord := "ozanc"
		if trie.ContainsWord(anotherWord) {
			t.Errorf("the word '%s' must not exist in the trie", anotherWord)
		}
	})
}
