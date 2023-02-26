package main

import (
	"fmt"

	"github.com/nikhilnarayanan623/ds-trie/trie"
)

func main() {

	t := trie.GetTrieNode()

	fmt.Println("Words finder")

	var (
		word  string
		limit int
	)

	fmt.Print("How many word enter: ")
	fmt.Scan(&limit)

	for i := 1; i <= limit; i++ {
		fmt.Print("Enter word: ")
		fmt.Scan(&word)
		t.Insert(word)
		word = ""
	}

	fmt.Print("\nEnter a word to remove: ")
	fmt.Scan(&word)

	t.Remove(word)

	fmt.Print("\nEnter a prefix to find its word: ")
	fmt.Scan(&word)

	words := t.FindWordsByPrefix(word)

	fmt.Println("\nWord founded ")

	for i, word := range words {
		fmt.Println(i+1, ": ", word)
	}

}
