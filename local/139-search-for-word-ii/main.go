// Problem:    Word Search II
// Difficulty: Hard
// Link:       https://neetcode.io/problems/search-for-word-ii
// Source:     Data Structures & Algorithms/search-for-word-ii/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Given a grid of letters and a list of words, return all words that can be
// found in the grid by moving up, down, left or right from letter to
// letter. The same cell cannot be used twice within one word.
//
// Example 1:
//
//	Input:  board = [
//	          ["a","b","c","d"],
//	          ["s","a","a","t"],
//	          ["a","c","k","e"],
//	          ["a","c","d","n"]
//	        ], words = ["bat", "cat", "back", "backend", "stack"]
//	Output: ["cat", "back", "backend"]
//
// Example 2:
//
//	Input:  board = [["x","o"], ["x","o"]], words = ["xoxo"]
//	Output: []
//
// How to run (from the "local" folder):
//
//	go run ./139-search-for-word-ii
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
	"sort"
)

// ===================== Solution =====================

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

// ===================== Test cases =====================

func main() {
	// byteGrid("ab", "cd") builds the board: each string is one row.
	// The words may come back in any order, so we sort them before comparing.

	// Example 1 from the problem.
	check("example 1", `sortedStrings(findWords(byteGrid("abcd", "saat", "acke", "acdn"), []string{"bat", "cat", "back", "backend", "stack"}))`,
		sortedStrings(findWords(byteGrid("abcd", "saat", "acke", "acdn"), []string{"bat", "cat", "back", "backend", "stack"})), []string{"back", "backend", "cat"})

	// Example 2 from the problem: "xoxo" would need a cell twice.
	check("example 2: nothing found", `sortedStrings(findWords(byteGrid("xo", "xo"), []string{"xoxo"}))`,
		sortedStrings(findWords(byteGrid("xo", "xo"), []string{"xoxo"})), []string{})

	// The classic example.
	check("classic", `sortedStrings(findWords(byteGrid("oaan", "etae", "ihkr", "iflv"), []string{"oath", "pea", "eat", "rain"}))`,
		sortedStrings(findWords(byteGrid("oaan", "etae", "ihkr", "iflv"), []string{"oath", "pea", "eat", "rain"})), []string{"eat", "oath"})

	// A board with one cell.
	check("one cell", `sortedStrings(findWords(byteGrid("a"), []string{"a", "b", "aa"}))`,
		sortedStrings(findWords(byteGrid("a"), []string{"a", "b", "aa"})), []string{"a"})

	// The word can be found in two places, but must be listed only once.
	check("found twice, listed once", `sortedStrings(findWords(byteGrid("aa"), []string{"a"}))`,
		sortedStrings(findWords(byteGrid("aa"), []string{"a"})), []string{"a"})

	// "aba" would need the same 'a' twice.
	check("cell cannot be reused", `sortedStrings(findWords(byteGrid("ab"), []string{"aba"}))`,
		sortedStrings(findWords(byteGrid("ab"), []string{"aba"})), []string{})

	// One word is the beginning of another: both must be found.
	check("word inside another word", `sortedStrings(findWords(byteGrid("abc"), []string{"ab", "abc", "a"}))`,
		sortedStrings(findWords(byteGrid("abc"), []string{"ab", "abc", "a"})), []string{"a", "ab", "abc"})

	// No words to look for.
	check("no words", `sortedStrings(findWords(byteGrid("ab", "cd"), []string{}))`,
		sortedStrings(findWords(byteGrid("ab", "cd"), []string{})), []string{})

	// The word goes around a corner: a -> b -> d -> c.
	check("around the square", `sortedStrings(findWords(byteGrid("ab", "cd"), []string{"abdc", "abcd"}))`,
		sortedStrings(findWords(byteGrid("ab", "cd"), []string{"abdc", "abcd"})), []string{"abdc"})

	report()
}

// ===================== Helpers for the tests =====================

// sortedStrings sorts the words. We use it when the answer may be in any order.
func sortedStrings(words []string) []string {
	sort.Strings(words)
	return words
}

// byteGrid turns rows written as strings into a [][]byte grid.
// Example: byteGrid("110", "001") is a grid with 2 rows and 3 columns.
func byteGrid(rows ...string) [][]byte {
	grid := [][]byte{}
	for _, row := range rows {
		grid = append(grid, []byte(row))
	}
	return grid
}

// passed and failed count the test results.
var passed, failed int

// check runs one test case. It prints the input, the result we got and the
// result we expected, with PASS or FAIL in front.
//
//	name:  a short description of the case
//	input: the call we are testing, written as text so it can be printed
//	got:   the answer the solution returned
//	want:  the answer we expect
func check(name string, input string, got interface{}, want interface{}) {
	status := "PASS"
	if fmt.Sprint(got) == fmt.Sprint(want) {
		passed++
	} else {
		failed++
		status = "FAIL"
	}
	fmt.Printf("%s  %s\n", status, name)
	fmt.Printf("      input:    %s\n", input)
	fmt.Printf("      result:   %s\n", show(got))
	fmt.Printf("      expected: %s\n\n", show(want))
}

// show turns a value into text. Strings get quotes, so an empty string is visible.
func show(value interface{}) string {
	if text, ok := value.(string); ok {
		return fmt.Sprintf("%q", text)
	}
	return fmt.Sprint(value)
}

// report prints the total. The program exits with an error if a test failed.
func report() {
	fmt.Printf("%d passed, %d failed\n", passed, failed)
	if failed > 0 {
		os.Exit(1)
	}
}
