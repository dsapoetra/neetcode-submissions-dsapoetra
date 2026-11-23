// Idea: store the words in a trie (one node per letter).
// Searching a normal letter follows one child.
// Searching a '.' can be ANY letter, so we try every child.

type wordNode struct {
	children map[byte]*wordNode
	isWord   bool // true if a word ends at this node
}

type WordDictionary struct {
	root *wordNode
}

func newWordNode() *wordNode {
	return &wordNode{children: map[byte]*wordNode{}}
}

func Constructor() WordDictionary {
	return WordDictionary{root: newWordNode()}
}

// AddWord walks down letter by letter, creating nodes that do not exist yet.
func (this *WordDictionary) AddWord(word string) {
	cur := this.root
	for i := 0; i < len(word); i++ {
		letter := word[i]
		if cur.children[letter] == nil {
			cur.children[letter] = newWordNode()
		}
		cur = cur.children[letter]
	}
	cur.isWord = true
}

func (this *WordDictionary) Search(word string) bool {
	return searchNode(this.root, word, 0)
}

// searchNode returns true if word[index:] can be matched starting at node.
func searchNode(node *wordNode, word string, index int) bool {
	// all letters matched: it only counts if a word ends here
	if index == len(word) {
		return node.isWord
	}

	letter := word[index]

	if letter == '.' {
		// '.' matches any letter: try every child
		for _, child := range node.children {
			if searchNode(child, word, index+1) {
				return true
			}
		}
		return false
	}

	// normal letter: there is only one child to follow
	child := node.children[letter]
	if child == nil {
		return false
	}
	return searchNode(child, word, index+1)
}
