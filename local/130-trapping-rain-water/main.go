// Problem:    Trapping Rain Water
// Difficulty: Hard
// Link:       https://neetcode.io/problems/trapping-rain-water
// Source:     Data Structures & Algorithms/trapping-rain-water/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Each number is the height of a bar of width 1. After rain, water stays in
// the dips between the bars. Return how many units of water are trapped.
//
// Example 1:
//
//	Input:  height = [0, 2, 0, 3, 1, 0, 1, 3, 2, 1]
//	Output: 9
//
// How to run (from the "local" folder):
//
//	go run ./130-trapping-rain-water
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea: the water above one bar is limited by the tallest bar on its left
// and the tallest bar on its right. The lower of those two is the water
// level, so:
//
//	water above bar i = min(tallest left, tallest right) - height[i]
//
// First compute "tallest bar so far" from the left and from the right,
// then add up the water above every bar.
// Time: O(n), extra space: O(n).
func trap(height []int) int {
	n := len(height)
	// no bars, no water
	if n == 0 {
		return 0
	}

	// maxLeft[i] = tallest bar from index 0 to i
	maxLeft := make([]int, n)
	maxLeft[0] = height[0]
	for i := 1; i < n; i++ {
		// start with the tallest bar before i ...
		maxLeft[i] = maxLeft[i-1]
		// ... and replace it if bar i itself is taller
		if height[i] > maxLeft[i] {
			maxLeft[i] = height[i]
		}
	}

	// maxRight[i] = tallest bar from index i to the end
	maxRight := make([]int, n)
	maxRight[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		// same as above, but walking from right to left
		maxRight[i] = maxRight[i+1]
		if height[i] > maxRight[i] {
			maxRight[i] = height[i]
		}
	}

	water := 0
	// add up the water above every bar
	for i := 0; i < n; i++ {
		// the water level is the LOWER of the two walls
		level := maxLeft[i]
		if maxRight[i] < level {
			level = maxRight[i]
		}
		// water = level minus the bar itself (never negative, because
		// both walls are at least as tall as bar i)
		water += level - height[i]
	}
	return water
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem.
	check("example 1", `trap([]int{0, 2, 0, 3, 1, 0, 1, 3, 2, 1})`,
		trap([]int{0, 2, 0, 3, 1, 0, 1, 3, 2, 1}), 9)

	// The classic examples.
	check("classic 1", `trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})`,
		trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}), 6)
	check("classic 2", `trap([]int{4, 2, 0, 3, 2, 5})`,
		trap([]int{4, 2, 0, 3, 2, 5}), 9)

	// No bars.
	check("no bars", `trap([]int{})`,
		trap([]int{}), 0)

	// One bar cannot hold water.
	check("one bar", `trap([]int{5})`,
		trap([]int{5}), 0)

	// Stairs going up or down: the water runs off.
	check("stairs up", `trap([]int{1, 2, 3})`,
		trap([]int{1, 2, 3}), 0)
	check("stairs down", `trap([]int{3, 2, 1})`,
		trap([]int{3, 2, 1}), 0)

	// A simple bowl: 3 units in the middle.
	check("bowl", `trap([]int{3, 0, 3})`,
		trap([]int{3, 0, 3}), 3)

	// Flat.
	check("flat", `trap([]int{2, 2, 2})`,
		trap([]int{2, 2, 2}), 0)

	// The lower wall (1) decides the water level.
	check("low right wall", `trap([]int{5, 0, 0, 1})`,
		trap([]int{5, 0, 0, 1}), 2)

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
