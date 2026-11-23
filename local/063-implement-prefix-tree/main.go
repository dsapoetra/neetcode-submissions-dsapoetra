// Problem:    Implement Trie (Prefix Tree)
// Difficulty: Medium
// Link:       https://neetcode.io/problems/implement-prefix-tree
// Source:     Data Structures & Algorithms/implement-prefix-tree/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Build a prefix tree (trie) with three operations:
//
//	Insert(word)        stores the word
//	Search(word)        true if exactly this word was inserted
//	StartsWith(prefix)  true if some inserted word starts with the prefix
//
// Example:
//
//	Input:  Insert("dog"), Search("dog"), Search("do"), StartsWith("do"),
//	        Insert("do"), Search("do")
//	Output: -, true, false, true, -, true
//
// How to run (from the "local" folder):
//
//	go run ./063-implement-prefix-tree
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea: a trie is a tree where every node is one letter.
// Walking from the root down spells a word. Each node has:
//   - children: the next possible letters
//   - isWord:   true if a word ends exactly here

type prefixNode struct {
	// next letter -> the node for that letter
	children map[byte]*prefixNode
	isWord   bool
}

type PrefixTree struct {
	root *prefixNode
}

// newPrefixNode makes an empty node with a ready-to-use children map.
func newPrefixNode() *prefixNode {
	return &prefixNode{children: map[byte]*prefixNode{}}
}

// Constructor makes an empty tree: only a root, which holds no letter.
func Constructor() PrefixTree {
	return PrefixTree{root: newPrefixNode()}
}

// Insert walks down letter by letter, creating nodes that do not exist yet.
func (this *PrefixTree) Insert(word string) {
	// cur is the node we are standing on; start at the root
	cur := this.root
	for i := 0; i < len(word); i++ {
		letter := word[i]
		// no node for this letter yet: create it
		if cur.children[letter] == nil {
			cur.children[letter] = newPrefixNode()
		}
		// step down to the node of this letter
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
	// step down one letter; a missing letter gives nil
	for i := 0; i < len(text); i++ {
		cur = cur.children[text[i]]
		if cur == nil {
			return nil
		}
	}
	// the whole text was found; cur is the node of its last letter
	return cur
}

// ===================== Test cases =====================

func main() {
	// The trie remembers what was inserted before, so the lines starting
	// with "--" show every Insert in the output.
	trie := Constructor()

	// Example from the problem.
	trie.Insert("dog")
	fmt.Println(`-- trie.Insert("dog")`)
	check("search a word we inserted", `trie.Search("dog")`,
		trie.Search("dog"), true)
	check("a prefix is not a word", `trie.Search("do")`,
		trie.Search("do"), false)
	check("but it IS a prefix", `trie.StartsWith("do")`,
		trie.StartsWith("do"), true)

	// Now insert the prefix as its own word.
	trie.Insert("do")
	fmt.Println(`-- trie.Insert("do")`)
	check("after inserting it, the prefix is a word", `trie.Search("do")`,
		trie.Search("do"), true)

	// Words that were never inserted.
	check("unknown word", `trie.Search("cat")`,
		trie.Search("cat"), false)
	check("unknown prefix", `trie.StartsWith("cat")`,
		trie.StartsWith("cat"), false)

	// A whole word also counts as a prefix.
	check("whole word as prefix", `trie.StartsWith("dog")`,
		trie.StartsWith("dog"), true)

	// Longer than anything we inserted.
	check("longer than the word (search)", `trie.Search("dogs")`,
		trie.Search("dogs"), false)
	check("longer than the word (prefix)", `trie.StartsWith("dogs")`,
		trie.StartsWith("dogs"), false)

	// A second word that shares no letters with the first.
	trie.Insert("apple")
	fmt.Println(`-- trie.Insert("apple")`)
	check("second word", `trie.Search("apple")`,
		trie.Search("apple"), true)
	check("prefix of second word is not a word", `trie.Search("app")`,
		trie.Search("app"), false)
	check("prefix of second word", `trie.StartsWith("app")`,
		trie.StartsWith("app"), true)

	// Every word starts with the empty prefix.
	check("empty prefix", `trie.StartsWith("")`,
		trie.StartsWith(""), true)

	// But the empty word was never inserted.
	check("empty word", `trie.Search("")`,
		trie.Search(""), false)

	// A new trie knows nothing about the old one.
	empty := Constructor()
	fmt.Println(`-- empty := Constructor()   (a new, empty trie)`)
	check("fresh trie is empty", `empty.Search("dog")`,
		empty.Search("dog"), false)

	report()
}

// ===================== Helpers for the tests =====================

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
