// Idea: searching the board once per word is too slow.
// Instead put ALL the words in a trie, then walk the board once:
// from every cell do a DFS, and move through the trie at the same time.
// If the trie has no child for the current letter, no word continues
// this way, so we stop early. When we reach the end of a word in the
// trie, we found that word.

type boardNode struct {
	children map[byte]*boardNode
	word     string // the full word, if a word ends at this node
}

func findWords(board [][]byte, words []string) []string {
	// build the trie
	root := &boardNode{children: map[byte]*boardNode{}}
	for _, word := range words {
		cur := root
		for i := 0; i < len(word); i++ {
			letter := word[i]
			if cur.children[letter] == nil {
				cur.children[letter] = &boardNode{children: map[byte]*boardNode{}}
			}
			cur = cur.children[letter]
		}
		cur.word = word
	}

	// start a search from every cell
	res := []string{}
	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[0]); c++ {
			res = collectWords(board, r, c, root, res)
		}
	}
	return res
}

// collectWords continues the search at cell (r, c). node is where we are
// in the trie. It returns res with any newly found words added.
func collectWords(board [][]byte, r int, c int, node *boardNode, res []string) []string {
	// outside the board
	if r < 0 || c < 0 || r >= len(board) || c >= len(board[0]) {
		return res
	}

	// follow this letter in the trie. If there is no child, no word
	// continues this way. (A used cell holds '#', which never has a child.)
	letter := board[r][c]
	next := node.children[letter]
	if next == nil {
		return res
	}

	// a word ends here: save it, and clear it so we do not add it twice
	if next.word != "" {
		res = append(res, next.word)
		next.word = ""
	}

	board[r][c] = '#' // mark as used on this path
	res = collectWords(board, r+1, c, next, res)
	res = collectWords(board, r-1, c, next, res)
	res = collectWords(board, r, c+1, next, res)
	res = collectWords(board, r, c-1, next, res)
	board[r][c] = letter // un-mark (backtrack)

	return res
}
