// Problem:    Design Add and Search Words Data Structure
// Difficulty: Medium
// Link:       https://neetcode.io/problems/design-word-search-data-structure
// Source:     Data Structures & Algorithms/design-word-search-data-structure/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Build a word dictionary with two operations:
//
//	AddWord(word)   stores the word
//	Search(word)    true if a stored word matches. In the search word a
//	                '.' matches any one letter.
//
// Example:
//
//	Input:  AddWord("day"), AddWord("bay"), AddWord("may"),
//	        Search("say"), Search("day"), Search(".ay"), Search("b..")
//	Output: -, -, -, false, true, true, true
//
// How to run (from the "local" folder):
//
//	go run ./064-design-word-search-data-structure
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

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

// ===================== Test cases =====================

func main() {
	// The dictionary remembers what was added before, so the lines starting
	// with "--" show every AddWord in the output.
	dict := Constructor()

	// Example from the problem.
	dict.AddWord("day")
	dict.AddWord("bay")
	dict.AddWord("may")
	fmt.Println(`-- dict.AddWord("day"), dict.AddWord("bay"), dict.AddWord("may")`)
	check("word that was never added", `dict.Search("say")`,
		dict.Search("say"), false)
	check("word that was added", `dict.Search("day")`,
		dict.Search("day"), true)
	check("dot as first letter", `dict.Search(".ay")`,
		dict.Search(".ay"), true)
	check("two dots", `dict.Search("b..")`,
		dict.Search("b.."), true)

	// A dot in the middle.
	check("dot in the middle", `dict.Search("d.y")`,
		dict.Search("d.y"), true)

	// Only dots: any 3-letter word matches.
	check("three dots", `dict.Search("...")`,
		dict.Search("..."), true)

	// Only dots, but no word has 4 letters.
	check("four dots", `dict.Search("....")`,
		dict.Search("...."), false)

	// Only dots, but no word has 2 letters.
	check("two dots only", `dict.Search("..")`,
		dict.Search(".."), false)

	// "da" is only the beginning of "day", not a whole word.
	check("prefix is not a word", `dict.Search("da")`,
		dict.Search("da"), false)

	// The dot matches, but the last letter does not.
	check("dot matches, rest does not", `dict.Search(".ax")`,
		dict.Search(".ax"), false)

	// Add a one-letter word.
	dict.AddWord("a")
	fmt.Println(`-- dict.AddWord("a")`)
	check("one-letter word", `dict.Search("a")`,
		dict.Search("a"), true)
	check("one dot", `dict.Search(".")`,
		dict.Search("."), true)
	check("other one-letter word", `dict.Search("b")`,
		dict.Search("b"), false)

	// A new dictionary is empty.
	empty := Constructor()
	fmt.Println(`-- empty := Constructor()   (a new, empty dictionary)`)
	check("fresh dictionary is empty", `empty.Search(".")`,
		empty.Search("."), false)

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
