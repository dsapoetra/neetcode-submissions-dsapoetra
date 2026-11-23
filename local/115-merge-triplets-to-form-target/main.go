// Problem:    Merge Triplets to Form Target Triplet
// Difficulty: Medium
// Link:       https://neetcode.io/problems/merge-triplets-to-form-target
// Source:     Data Structures & Algorithms/merge-triplets-to-form-target/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// You have a list of triplets [a, b, c]. You may merge two triplets: the
// result takes the bigger value at each of the three positions. Return true
// if the target triplet can be produced by merging (or is already in the list).
//
// Example 1:
//
//	Input:  triplets = [[1, 2, 3], [7, 1, 1]], target = [7, 2, 3]
//	Output: true
//
// Example 2:
//
//	Input:  triplets = [[2, 5, 6], [1, 4, 4], [5, 7, 5]], target = [5, 4, 6]
//	Output: false
//
// How to run (from the "local" folder):
//
//	go run ./115-merge-triplets-to-form-target
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea: merging takes the MAXIMUM of every position, so values can only
// grow. A triplet with any value bigger than the target would break the
// target forever, so we may never use it.
// All the other triplets are safe to merge. So: look only at the safe
// triplets and check that each of the three target values appears in at
// least one of them at the right position.
// Time: O(n).
func mergeTriplets(triplets [][]int, target []int) bool {
	found := [3]bool{} // found[i] = some safe triplet has target[i] at position i

	for _, triplet := range triplets {
		// one value too big: this triplet can never be used
		if triplet[0] > target[0] || triplet[1] > target[1] || triplet[2] > target[2] {
			continue
		}

		for i := 0; i < 3; i++ {
			if triplet[i] == target[i] {
				found[i] = true
			}
		}
	}
	return found[0] && found[1] && found[2]
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: merging both gives [7,2,3].
	check("example 1", `mergeTriplets([][]int{{1, 2, 3}, {7, 1, 1}}, []int{7, 2, 3})`,
		mergeTriplets([][]int{{1, 2, 3}, {7, 1, 1}}, []int{7, 2, 3}), true)

	// Example 2 from the problem: the 4 in the middle only comes with a 1 or a too-big value.
	check("example 2", `mergeTriplets([][]int{{2, 5, 6}, {1, 4, 4}, {5, 7, 5}}, []int{5, 4, 6})`,
		mergeTriplets([][]int{{2, 5, 6}, {1, 4, 4}, {5, 7, 5}}, []int{5, 4, 6}), false)

	// The classic examples.
	check("classic true", `mergeTriplets([][]int{{2, 5, 3}, {1, 8, 4}, {1, 7, 5}}, []int{2, 7, 5})`,
		mergeTriplets([][]int{{2, 5, 3}, {1, 8, 4}, {1, 7, 5}}, []int{2, 7, 5}), true)
	check("classic false", `mergeTriplets([][]int{{3, 4, 5}, {4, 5, 6}}, []int{3, 2, 5})`,
		mergeTriplets([][]int{{3, 4, 5}, {4, 5, 6}}, []int{3, 2, 5}), false)
	check("classic three pieces", `mergeTriplets([][]int{{2, 5, 3}, {2, 3, 4}, {1, 2, 5}, {5, 2, 3}}, []int{5, 5, 5})`,
		mergeTriplets([][]int{{2, 5, 3}, {2, 3, 4}, {1, 2, 5}, {5, 2, 3}}, []int{5, 5, 5}), true)

	// The target itself is in the list.
	check("target is in the list", `mergeTriplets([][]int{{1, 1, 1}}, []int{1, 1, 1})`,
		mergeTriplets([][]int{{1, 1, 1}}, []int{1, 1, 1}), true)

	// Values can only grow by merging, never reach something that is not there.
	check("values too small", `mergeTriplets([][]int{{1, 1, 1}}, []int{2, 2, 2})`,
		mergeTriplets([][]int{{1, 1, 1}}, []int{2, 2, 2}), false)

	// The only triplet with the right 5 also has a 9, which is too big.
	check("right value, but poisoned", `mergeTriplets([][]int{{5, 1, 9}, {1, 1, 1}}, []int{5, 1, 1})`,
		mergeTriplets([][]int{{5, 1, 9}, {1, 1, 1}}, []int{5, 1, 1}), false)

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
