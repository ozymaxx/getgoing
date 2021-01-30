package ozymaxx

import "testing"

func TestTrie(t *testing.T) {
	trie := (&Trie{}).Init()
	t.Run("empty value already exists", func(t *testing.T) {
		if !trie.ContainsWord("") {
			t.Error("the empty word must exist in the trie")
		}
	})
	t.Run("when 'ozan' is added, both the empty word and 'ozan' must exist in the trie", func(t *testing.T) {
		word := "ozan"
		if trie.ContainsWord(word) {
			t.Errorf("the word '%s' must not exist in the trie yet because it has not been added yet", word)
		}
		trie.AddWord(word)
		if !trie.ContainsWord("") {
			t.Error("the empty word must exist in the trie")
		}
		if !trie.ContainsWord(word) {
			t.Errorf("the word '%s' must exist in the trie", word)
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
	t.Run("when 'abc', 'abcd' and 'ab' is added in order, both must exist in the trie", func(t *testing.T) {
		word1 := "abc"
		word2 := "abcd"
		word3 := "ab"
		trie.AddWord(word1)
		trie.AddWord(word2)
		if !trie.ContainsWord(word1) {
			t.Errorf("the word '%s' must exist in the trie", word1)
		}
		if !trie.ContainsWord(word2) {
			t.Errorf("the word '%s' must exist in the trie", word2)
		}
		if trie.ContainsWord(word3) {
			t.Errorf("the word '%s' must not exist in the trie yet because it has not been added yet", word3)
		}
		trie.AddWord(word3)
		if !trie.ContainsWord(word3) {
			t.Errorf("the word '%s' must exist in the trie", word3)
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
