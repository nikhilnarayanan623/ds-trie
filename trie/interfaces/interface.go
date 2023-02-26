package interfaces

type TrieInterface interface {
	Insert(str string)                        // to insert a new string into trie
	Contains(str string) bool                 // to check the string is contain in the trie
	Remove(str string)                        // to remove a string from the trie
	FindWordsByPrefix(prefix string) []string // to get all word that begin from the prefix
}
