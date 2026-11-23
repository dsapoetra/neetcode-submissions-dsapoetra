// Problem:    Last Stone Weight
// Difficulty: Easy
// Link:       https://neetcode.io/problems/last-stone-weight
// Source:     Data Structures & Algorithms/last-stone-weight/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Each number is the weight of a stone. Repeat: take the two heaviest
// stones and smash them together. If they weigh the same, both are
// destroyed; otherwise the lighter one is destroyed and the heavier one
// loses the weight of the lighter one. Return the weight of the last stone,
// or 0 if none is left.
//
// Example 1:
//
//	Input:  stones = [2, 3, 6, 2, 4]
//	Output: 1
//
// Example 2:
//
//	Input:  stones = [1, 2]
//	Output: 1
//
// How to run (from the "local" folder):
//
//	go run ./018-last-stone-weight
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

// Idea: do exactly what the problem says. Sort the stones, take the two
// heaviest and smash them. If they are different, the rest goes back.
// Repeat until at most one stone is left.
// Time: O(n^2 log n) because we sort every round. That is fine here, there
// are only a few stones. (A max-heap would be faster, but needs more code.)
func lastStoneWeight(stones []int) int {
	for len(stones) > 1 {
		sort.Ints(stones)

		heaviest := stones[len(stones)-1]
		second := stones[len(stones)-2]
		stones = stones[:len(stones)-2] // remove both

		if heaviest != second {
			stones = append(stones, heaviest-second) // what is left of the heavier one
		}
	}

	if len(stones) == 0 {
		return 0
	}
	return stones[0]
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: 6-4=2, 3-2=1, 2-2=0, one stone (1) is left.
	check("example 1", `lastStoneWeight([]int{2, 3, 6, 2, 4})`,
		lastStoneWeight([]int{2, 3, 6, 2, 4}), 1)

	// Example 2 from the problem.
	check("example 2", `lastStoneWeight([]int{1, 2})`,
		lastStoneWeight([]int{1, 2}), 1)

	// The classic example.
	check("classic", `lastStoneWeight([]int{2, 7, 4, 1, 8, 1})`,
		lastStoneWeight([]int{2, 7, 4, 1, 8, 1}), 1)

	// One stone.
	check("one stone", `lastStoneWeight([]int{1})`,
		lastStoneWeight([]int{1}), 1)

	// Two equal stones destroy each other: nothing is left.
	check("two equal stones", `lastStoneWeight([]int{3, 3})`,
		lastStoneWeight([]int{3, 3}), 0)

	// Two different stones.
	check("two different stones", `lastStoneWeight([]int{10, 4})`,
		lastStoneWeight([]int{10, 4}), 6)

	// Three equal stones: two destroy each other, one is left.
	check("three equal stones", `lastStoneWeight([]int{5, 5, 5})`,
		lastStoneWeight([]int{5, 5, 5}), 5)

	// Four equal stones: all gone.
	check("four equal stones", `lastStoneWeight([]int{1, 1, 1, 1})`,
		lastStoneWeight([]int{1, 1, 1, 1}), 0)

	// 3-2 = 1, then 1-1 = 0.
	check("3, 2, 1", `lastStoneWeight([]int{3, 2, 1})`,
		lastStoneWeight([]int{3, 2, 1}), 0)

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
