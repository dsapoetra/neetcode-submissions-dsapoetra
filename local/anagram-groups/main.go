// Problem: Group Anagrams
// Link:    https://neetcode.io/problems/anagram-groups
// Source:  Data Structures & Algorithms/anagram-groups/submission-3.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./anagram-groups
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
	check("example 1",
		sortGroups(groupAnagrams([]string{"act", "pots", "tops", "cat", "stop", "hat"})),
		[][]string{{"act", "cat"}, {"hat"}, {"pots", "stop", "tops"}})

	// Example 2 from the problem: one word is one group.
	check("example 2: one word", sortGroups(groupAnagrams([]string{"x"})), [][]string{{"x"}})

	// Example 3 from the problem: the empty string is a valid word.
	check("example 3: empty string", len(groupAnagrams([]string{""})), 1)

	// No word is an anagram of another: every word is alone.
	check("no anagrams",
		sortGroups(groupAnagrams([]string{"abc", "def", "ghi"})),
		[][]string{{"abc"}, {"def"}, {"ghi"}})

	// Every word is an anagram of the others: one big group.
	check("all anagrams",
		sortGroups(groupAnagrams([]string{"abc", "bca", "cab"})),
		[][]string{{"abc", "bca", "cab"}})

	// The same word twice goes into the same group twice.
	check("same word twice", sortGroups(groupAnagrams([]string{"aa", "aa"})), [][]string{{"aa", "aa"}})

	// Same letters but different counts are NOT anagrams.
	check("same letters, different counts",
		sortGroups(groupAnagrams([]string{"aab", "abb"})),
		[][]string{{"aab"}, {"abb"}})

	// Words of different lengths.
	check("different lengths",
		sortGroups(groupAnagrams([]string{"a", "ab", "ba", "abc"})),
		[][]string{{"a"}, {"ab", "ba"}, {"abc"}})

	// No words at all: no groups.
	check("empty input", len(groupAnagrams([]string{})), 0)

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
