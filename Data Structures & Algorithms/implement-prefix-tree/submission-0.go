// Idea: a trie is a tree where every node is one letter.
// Walking from the root down spells a word. Each node has:
//   - children: the next possible letters
//   - isWord:   true if a word ends exactly here

type prefixNode struct {
	children map[byte]*prefixNode
	isWord   bool
}

type PrefixTree struct {
	root *prefixNode
}

func newPrefixNode() *prefixNode {
	return &prefixNode{children: map[byte]*prefixNode{}}
}

func Constructor() PrefixTree {
	return PrefixTree{root: newPrefixNode()}
}

// Insert walks down letter by letter, creating nodes that do not exist yet.
func (this *PrefixTree) Insert(word string) {
	cur := this.root
	for i := 0; i < len(word); i++ {
		letter := word[i]
		if cur.children[letter] == nil {
			cur.children[letter] = newPrefixNode()
		}
		cur = cur.children[letter]
	}
	cur.isWord = true // mark the end of the word
}

// Search is true only if the path exists AND a word ends there.
func (this *PrefixTree) Search(word string) bool {
	node := this.find(word)
	return node != nil && node.isWord
}

// StartsWith is true if the path exists at all.
func (this *PrefixTree) StartsWith(prefix string) bool {
	return this.find(prefix) != nil
}

// find walks down the letters of text and returns the last node,
// or nil if the path does not exist.
func (this *PrefixTree) find(text string) *prefixNode {
	cur := this.root
	for i := 0; i < len(text); i++ {
		cur = cur.children[text[i]]
		if cur == nil {
			return nil
		}
	}
	return cur
}
