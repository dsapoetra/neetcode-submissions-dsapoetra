// Problem:    Sliding Window Maximum
// Difficulty: Hard
// Link:       https://neetcode.io/problems/sliding-window-maximum
// Source:     Data Structures & Algorithms/sliding-window-maximum/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// A window of size k slides over the list from left to right, one step at a
// time. Return the biggest number of every window.
//
// Example 1:
//
//	Input:  nums = [1, 2, 1, 0, 4, 2, 6], k = 3
//	Output: [2, 2, 4, 4, 6]
//
// How to run (from the "local" folder):
//
//	go run ./132-sliding-window-maximum
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea: keep a list ("queue") of INDEXES of numbers that could still be
// the maximum of a window. The numbers at those indexes are always in
// decreasing order, so the front of the queue is the current maximum.
// For every new number:
//  1. remove smaller numbers from the back: they can never be the maximum
//     again, because the new number is bigger AND stays longer
//  2. add the new index at the back
//  3. remove the front if it has slid out of the window
//
// Time: O(n), every index is added once and removed at most once.
func maxSlidingWindow(nums []int, k int) []int {
	res := []int{}
	queue := []int{} // indexes; their numbers are in decreasing order

	for right := 0; right < len(nums); right++ {
		// 1) drop smaller numbers from the back
		for len(queue) > 0 && nums[queue[len(queue)-1]] < nums[right] {
			queue = queue[:len(queue)-1]
		}

		// 2) add the new index
		queue = append(queue, right)

		// 3) the window is [right-k+1 .. right]; drop the front if it is outside
		if queue[0] < right-k+1 {
			queue = queue[1:]
		}

		// once the first full window is reached, the front is its maximum
		if right >= k-1 {
			res = append(res, nums[queue[0]])
		}
	}
	return res
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem.
	check("example 1", `maxSlidingWindow([]int{1, 2, 1, 0, 4, 2, 6}, 3)`,
		maxSlidingWindow([]int{1, 2, 1, 0, 4, 2, 6}, 3), []int{2, 2, 4, 4, 6})

	// The classic example.
	check("classic", `maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)`,
		maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3), []int{3, 3, 5, 5, 6, 7})

	// One number.
	check("one number", `maxSlidingWindow([]int{1}, 1)`,
		maxSlidingWindow([]int{1}, 1), []int{1})

	// Going down: the maximum is always the first number of the window.
	check("descending", `maxSlidingWindow([]int{4, 3, 2, 1}, 2)`,
		maxSlidingWindow([]int{4, 3, 2, 1}, 2), []int{4, 3, 2})

	// Going up: the maximum is always the last number of the window.
	check("ascending", `maxSlidingWindow([]int{1, 2, 3, 4}, 2)`,
		maxSlidingWindow([]int{1, 2, 3, 4}, 2), []int{2, 3, 4})

	// Equal numbers.
	check("all the same", `maxSlidingWindow([]int{5, 5, 5}, 2)`,
		maxSlidingWindow([]int{5, 5, 5}, 2), []int{5, 5})

	// Window of size 1: the list itself.
	check("window size 1", `maxSlidingWindow([]int{1, -1}, 1)`,
		maxSlidingWindow([]int{1, -1}, 1), []int{1, -1})

	// The window is the whole list.
	check("window = whole list", `maxSlidingWindow([]int{9, 1, 2}, 3)`,
		maxSlidingWindow([]int{9, 1, 2}, 3), []int{9})

	// A big number stays the maximum until it leaves the window.
	check("maximum stays", `maxSlidingWindow([]int{-7, -8, 7, 5, 7, 1, 6, 0}, 4)`,
		maxSlidingWindow([]int{-7, -8, 7, 5, 7, 1, 6, 0}, 4), []int{7, 7, 7, 7, 7})

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
