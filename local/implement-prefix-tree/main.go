// Problem: Implement Trie (Prefix Tree)
// Link:    https://neetcode.io/problems/implement-prefix-tree
// Source:  Data Structures & Algorithms/implement-prefix-tree/submission-0.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./implement-prefix-tree
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

// ===================== Test cases =====================

func main() {
	trie := Constructor()

	// Example from the problem.
	trie.Insert("dog")
	check("search a word we inserted", trie.Search("dog"), true)
	check("a prefix is not a word", trie.Search("do"), false)
	check("but it IS a prefix", trie.StartsWith("do"), true)

	// Now insert the prefix as its own word.
	trie.Insert("do")
	check("after inserting it, the prefix is a word", trie.Search("do"), true)

	// Words that were never inserted.
	check("unknown word", trie.Search("cat"), false)
	check("unknown prefix", trie.StartsWith("cat"), false)

	// A whole word also counts as a prefix.
	check("whole word as prefix", trie.StartsWith("dog"), true)

	// Longer than anything we inserted.
	check("longer than the word (search)", trie.Search("dogs"), false)
	check("longer than the word (prefix)", trie.StartsWith("dogs"), false)

	// A second word that shares no letters with the first.
	trie.Insert("apple")
	check("second word", trie.Search("apple"), true)
	check("prefix of second word is not a word", trie.Search("app"), false)
	check("prefix of second word", trie.StartsWith("app"), true)

	// Every word starts with the empty prefix.
	check("empty prefix", trie.StartsWith(""), true)

	// But the empty word was never inserted.
	check("empty word", trie.Search(""), false)

	// A new trie knows nothing about the old one.
	empty := Constructor()
	check("fresh trie is empty", empty.Search("dog"), false)

	report()
}

// ===================== Helpers for the tests =====================

// passed and failed count the test results.
var passed, failed int

// check compares the answer we got with the answer we want, and prints
// PASS or FAIL for that test case.
func check(name string, got interface{}, want interface{}) {
	if fmt.Sprint(got) == fmt.Sprint(want) {
		passed++
		fmt.Printf("PASS  %s\n", name)
	} else {
		failed++
		fmt.Printf("FAIL  %s\n      got:  %v\n      want: %v\n", name, got, want)
	}
}

// report prints the total. The program exits with an error if a test failed.
func report() {
	fmt.Printf("\n%d passed, %d failed\n", passed, failed)
	if failed > 0 {
		os.Exit(1)
	}
}
