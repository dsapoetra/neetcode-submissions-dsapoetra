// Problem:    Hand of Straights
// Difficulty: Medium
// Link:       https://neetcode.io/problems/hand-of-straights
// Source:     Data Structures & Algorithms/hand-of-straights/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Given a hand of cards (numbers) and a group size, return true if all
// cards can be arranged into groups of exactly that size, where each group
// is made of consecutive numbers (like 3, 4, 5).
//
// Example 1:
//
//	Input:  hand = [1, 2, 4, 2, 3, 5, 3, 4], groupSize = 4
//	Output: true          ([1, 2, 3, 4] and [2, 3, 4, 5])
//
// Example 2:
//
//	Input:  hand = [1, 2, 3, 3, 4, 5, 6, 7], groupSize = 4
//	Output: false
//
// How to run (from the "local" folder):
//
//	go run ./114-hand-of-straights
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

// Idea (greedy): the SMALLEST card left must be the start of a group,
// because nothing smaller can come before it.
// So sort the cards and count how many of each value we have. Go through
// the cards from small to big: if a card is still available, it starts a
// group, and we must also take value+1, value+2, ... If one of them is
// missing, it is impossible.
// Time: O(n log n + n * groupSize).
func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false // the cards cannot be split evenly
	}

	sort.Ints(hand)
	count := map[int]int{}
	for _, card := range hand {
		count[card]++
	}

	for _, card := range hand {
		if count[card] == 0 {
			continue // already used in an earlier group
		}

		// card is the smallest card left: build card, card+1, ...
		for value := card; value < card+groupSize; value++ {
			if count[value] == 0 {
				return false // a card is missing
			}
			count[value]--
		}
	}
	return true
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: [1,2,3,4] and [2,3,4,5].
	check("example 1", `isNStraightHand([]int{1, 2, 4, 2, 3, 5, 3, 4}, 4)`,
		isNStraightHand([]int{1, 2, 4, 2, 3, 5, 3, 4}, 4), true)

	// Example 2 from the problem: after [1,2,3,4] the cards 3,5,6,7 are left.
	check("example 2", `isNStraightHand([]int{1, 2, 3, 3, 4, 5, 6, 7}, 4)`,
		isNStraightHand([]int{1, 2, 3, 3, 4, 5, 6, 7}, 4), false)

	// The classic examples.
	check("classic true", `isNStraightHand([]int{1, 2, 3, 6, 2, 3, 4, 7, 8}, 3)`,
		isNStraightHand([]int{1, 2, 3, 6, 2, 3, 4, 7, 8}, 3), true)
	check("classic false: 5 cards, groups of 4", `isNStraightHand([]int{1, 2, 3, 4, 5}, 4)`,
		isNStraightHand([]int{1, 2, 3, 4, 5}, 4), false)

	// Groups of 1 always work.
	check("groups of 1", `isNStraightHand([]int{5}, 1)`,
		isNStraightHand([]int{5}, 1), true)
	check("groups of 1, any cards", `isNStraightHand([]int{1, 9, 4}, 1)`,
		isNStraightHand([]int{1, 9, 4}, 1), true)

	// Two equal groups.
	check("two equal groups", `isNStraightHand([]int{1, 1, 2, 2}, 2)`,
		isNStraightHand([]int{1, 1, 2, 2}, 2), true)

	// [1,2] works, but then 1 and 3 are left.
	check("gap left over", `isNStraightHand([]int{1, 1, 2, 3}, 2)`,
		isNStraightHand([]int{1, 1, 2, 3}, 2), false)

	// The cards are not consecutive.
	check("not consecutive", `isNStraightHand([]int{8, 10, 12}, 3)`,
		isNStraightHand([]int{8, 10, 12}, 3), false)

	// Not sorted.
	check("unsorted", `isNStraightHand([]int{3, 2, 1}, 3)`,
		isNStraightHand([]int{3, 2, 1}, 3), true)

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
