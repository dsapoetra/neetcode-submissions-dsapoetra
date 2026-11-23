// Problem:    Largest Rectangle in Histogram
// Difficulty: Hard
// Link:       https://neetcode.io/problems/largest-rectangle-in-histogram
// Source:     Data Structures & Algorithms/largest-rectangle-in-histogram/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Each number is the height of a bar of width 1 in a histogram. Return the
// area of the largest rectangle that fits completely inside the bars.
//
// Example 1:
//
//	Input:  heights = [7, 1, 7, 2, 2, 4]
//	Output: 8
//
// Example 2:
//
//	Input:  heights = [1, 3, 7]
//	Output: 7
//
// How to run (from the "local" folder):
//
//	go run ./136-largest-rectangle-in-histogram
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea: for every bar we want to know how far a rectangle of that bar's
// height can stretch to the left and to the right.
// Keep a stack of bars that are still "open" (their rectangle can still
// grow to the right). Heights in the stack go up.
// When a new bar is LOWER than the bar on top, the top bar's rectangle
// must end here: pop it and compute its area.
// The new bar can stretch back to where the popped bars started, because
// they were all taller.
// Time: O(n), every bar is pushed and popped once.

type openBar struct {
	start  int // the leftmost index this height reaches
	height int
}

func largestRectangleArea(heights []int) int {
	// the biggest area found so far
	best := 0
	stack := []openBar{}

	for i, height := range heights {
		// a new bar starts at its own index, unless it can reach further back
		start := i

		// taller bars on the stack end here
		for len(stack) > 0 && stack[len(stack)-1].height > height {
			// take the taller bar off the stack
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// its rectangle goes from top.start up to (not including) i
			area := top.height * (i - top.start)
			if area > best {
				best = area
			}
			start = top.start // the new, lower bar reaches back this far
		}

		// the new bar is now open
		stack = append(stack, openBar{start, height})
	}

	// bars still on the stack reach all the way to the end
	for _, bar := range stack {
		// width = from bar.start to the end of the histogram
		area := bar.height * (len(heights) - bar.start)
		if area > best {
			best = area
		}
	}
	return best
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: the last four bars (7, 2, 2, 4) are all at least 2 high: 4 * 2 = 8.
	check("example 1", `largestRectangleArea([]int{7, 1, 7, 2, 2, 4})`,
		largestRectangleArea([]int{7, 1, 7, 2, 2, 4}), 8)

	// Example 2 from the problem: the single bar of height 7.
	check("example 2", `largestRectangleArea([]int{1, 3, 7})`,
		largestRectangleArea([]int{1, 3, 7}), 7)

	// The classic example: the bars 5 and 6 give 5 * 2.
	check("classic", `largestRectangleArea([]int{2, 1, 5, 6, 2, 3})`,
		largestRectangleArea([]int{2, 1, 5, 6, 2, 3}), 10)

	// Two bars.
	check("two bars", `largestRectangleArea([]int{2, 4})`,
		largestRectangleArea([]int{2, 4}), 4)

	// One bar.
	check("one bar", `largestRectangleArea([]int{5})`,
		largestRectangleArea([]int{5}), 5)

	// All the same height: the whole histogram.
	check("all the same", `largestRectangleArea([]int{3, 3, 3})`,
		largestRectangleArea([]int{3, 3, 3}), 9)

	// Stairs: 3 * 3 (the bars 3, 4, 5).
	check("stairs up", `largestRectangleArea([]int{1, 2, 3, 4, 5})`,
		largestRectangleArea([]int{1, 2, 3, 4, 5}), 9)
	check("stairs down", `largestRectangleArea([]int{5, 4, 3, 2, 1})`,
		largestRectangleArea([]int{5, 4, 3, 2, 1}), 9)

	// A zero splits the histogram.
	check("zero in the middle", `largestRectangleArea([]int{2, 0, 2})`,
		largestRectangleArea([]int{2, 0, 2}), 2)

	// No bars.
	check("no bars", `largestRectangleArea([]int{})`,
		largestRectangleArea([]int{}), 0)

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
