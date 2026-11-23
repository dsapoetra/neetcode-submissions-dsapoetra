// Problem:    Group Anagrams
// Difficulty: Medium
// Link:       https://neetcode.io/problems/anagram-groups
// Source:     Data Structures & Algorithms/anagram-groups/submission-3.go (the solution below is a copy of it)
//
// THE PROBLEM
// Given a list of words, put the words that are anagrams of each other into
// the same group. The groups (and the words inside a group) may be returned
// in any order.
//
// Example 1:
//
//	Input:  strs = ["act", "pots", "tops", "cat", "stop", "hat"]
//	Output: [["hat"], ["act", "cat"], ["stop", "pots", "tops"]]
//
// Example 2:
//
//	Input:  strs = ["x"]
//	Output: [["x"]]
//
// How to run (from the "local" folder):
//
//	go run ./029-anagram-groups
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"
)

// ===================== Solution =====================

func groupAnagrams(strs []string) [][]string {
	groups := map[string][]string{}
	for _, w := range strs {
		r := []rune(w)
		slices.Sort(r)
		key := string(r)
		groups[key] = append(groups[key], w)
	}
	// {"abc": ["abc", "bca", "cab"]}
	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}
	return result

}

// ===================== Test cases =====================

func main() {
	// The groups may come back in any order, so we sort them before comparing.

	// Example 1 from the problem.
	check("example 1", `sortGroups(groupAnagrams([]string{"act", "pots", "tops", "cat", "stop", "hat"}))`,
		sortGroups(groupAnagrams([]string{"act", "pots", "tops", "cat", "stop", "hat"})), [][]string{{"act", "cat"}, {"hat"}, {"pots", "stop", "tops"}})

	// Example 2 from the problem: one word is one group.
	check("example 2: one word", `sortGroups(groupAnagrams([]string{"x"}))`,
		sortGroups(groupAnagrams([]string{"x"})), [][]string{{"x"}})

	// Example 3 from the problem: the empty string is a valid word.
	check("example 3: empty string", `len(groupAnagrams([]string{""}))`,
		len(groupAnagrams([]string{""})), 1)

	// No word is an anagram of another: every word is alone.
	check("no anagrams", `sortGroups(groupAnagrams([]string{"abc", "def", "ghi"}))`,
		sortGroups(groupAnagrams([]string{"abc", "def", "ghi"})), [][]string{{"abc"}, {"def"}, {"ghi"}})

	// Every word is an anagram of the others: one big group.
	check("all anagrams", `sortGroups(groupAnagrams([]string{"abc", "bca", "cab"}))`,
		sortGroups(groupAnagrams([]string{"abc", "bca", "cab"})), [][]string{{"abc", "bca", "cab"}})

	// The same word twice goes into the same group twice.
	check("same word twice", `sortGroups(groupAnagrams([]string{"aa", "aa"}))`,
		sortGroups(groupAnagrams([]string{"aa", "aa"})), [][]string{{"aa", "aa"}})

	// Same letters but different counts are NOT anagrams.
	check("same letters, different counts", `sortGroups(groupAnagrams([]string{"aab", "abb"}))`,
		sortGroups(groupAnagrams([]string{"aab", "abb"})), [][]string{{"aab"}, {"abb"}})

	// Words of different lengths.
	check("different lengths", `sortGroups(groupAnagrams([]string{"a", "ab", "ba", "abc"}))`,
		sortGroups(groupAnagrams([]string{"a", "ab", "ba", "abc"})), [][]string{{"a"}, {"ab", "ba"}, {"abc"}})

	// No words at all: no groups.
	check("empty input", `len(groupAnagrams([]string{}))`,
		len(groupAnagrams([]string{})), 0)

	report()
}

// ===================== Helpers for the tests =====================

// sortGroups sorts the words inside every group and then sorts the groups.
// We use it because the groups may come back in any order.
func sortGroups(groups [][]string) [][]string {
	for _, group := range groups {
		sort.Strings(group)
	}
	sort.Slice(groups, func(a, b int) bool {
		return strings.Join(groups[a], ",") < strings.Join(groups[b], ",")
	})
	return groups
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
