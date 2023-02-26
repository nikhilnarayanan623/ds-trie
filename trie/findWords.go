package trie

// find all word by its prefix

func (t *trie) FindWordsByPrefix(str string) []string {
	// array of string to store all words
	words := []string{}
	// to store a single word in character array then convert into string
	word := []rune{}
	// start from root
	node := t.root

	// loop through the prefix
	for _, char := range str {
		// if any of prefix character is not in trie connction then return
		if _, ok := node.childerens[char]; !ok {
			return []string{}
		}
		// go that char node
		node = node.childerens[char]
		// also add that charcter into word array
		word = append(word, char)
	}
	// if prefix is in available in the trie then call findWords to get all its words
	t.findWords(node, word, &words)

	return words
}

func (t *trie) findWords(node *trieNode, word []rune, words *[]string) {

	// if this node is the end of word then copy the word array to a new word array and stringify it and store it on words array
	if node.endOfWord {
		copyWord := make([]rune, len(word))
		copy(copyWord, word)

		(*words) = append((*words), string(copyWord))
	}

	// range through the nodes children
	for char, child := range node.childerens {
		// add every char into word
		word = append(word, char)
		// call the same function to get its childrens
		t.findWords(child, word, words)

		// delete the last char to add next char
		word = word[:len(word)-1]
	}
}
