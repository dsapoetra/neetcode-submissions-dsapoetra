// Problem:    Container With Most Water
// Difficulty: Medium
// Link:       https://neetcode.io/problems/max-water-container
// Source:     Data Structures & Algorithms/max-water-container/submission-1.go (the solution below is a copy of it)
//
// THE PROBLEM
// Each number is the height of a vertical bar. Choose two bars; together
// with the ground they form a container. Return the largest amount of water
// a container can hold: (distance between the bars) * (height of the lower bar).
//
// Example 1:
//
//	Input:  heights = [1, 7, 2, 5, 4, 7, 3, 6]
//	Output: 36            (bars at index 1 and 7: width 6 * height 6)
//
// Example 2:
//
//	Input:  heights = [2, 2, 2]
//	Output: 4
//
// How to run (from the "local" folder):
//
//	go run ./037-max-water-container
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea (two pointers): start with the widest container (first and last
// wall). The water height is limited by the SHORTER wall. Moving the
// taller wall inwards can never help (width shrinks, height cannot grow),
// so always move the shorter wall and hope for a taller one.
// Time: O(n).
func maxArea(heights []int) int {
	// start with the two outer walls: the widest container
	left := 0
	right := len(heights) - 1
	best := 0

	for left < right {
		// water height = the shorter of the two walls
		height := heights[left]
		if heights[right] < height {
			height = heights[right]
		}

		// area = height * width
		area := height * (right - left)
		if area > best {
			best = area
		}

		// move the shorter wall inwards
		if heights[left] < heights[right] {
			// the left wall is the shorter one: try the next wall on the left side
			left++
		} else {
			// the right wall is shorter (or equal): try the next one on the right
			right--
		}
	}
	return best
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: walls 7 (index 1) and 6 (index 7): 6 * 6 = 36.
	check("example 1", `maxArea([]int{1, 7, 2, 5, 4, 7, 3, 6})`,
		maxArea([]int{1, 7, 2, 5, 4, 7, 3, 6}), 36)

	// Example 2 from the problem: first and last wall: 2 * 2 = 4.
	check("example 2", `maxArea([]int{2, 2, 2})`,
		maxArea([]int{2, 2, 2}), 4)

	// The smallest input: two walls.
	check("two walls", `maxArea([]int{1, 1})`,
		maxArea([]int{1, 1}), 1)

	// The classic example: walls 8 and 7, 7 apart: 7 * 7 = 49.
	check("classic", `maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})`,
		maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}), 49)

	// The two tallest walls are at the ends.
	check("tall walls at the ends", `maxArea([]int{4, 3, 2, 1, 4})`,
		maxArea([]int{4, 3, 2, 1, 4}), 16)

	// A tall wall in the middle does not help: the ends give 1 * 2 = 2.
	check("tall wall in the middle", `maxArea([]int{1, 2, 1})`,
		maxArea([]int{1, 2, 1}), 2)

	// Walls of height 0 hold no water.
	check("zero height", `maxArea([]int{0, 0})`,
		maxArea([]int{0, 0}), 0)

	// Heights going up: best is 2 * 3 = 6 (or 3 * 2 = 6).
	check("ascending", `maxArea([]int{1, 2, 3, 4, 5})`,
		maxArea([]int{1, 2, 3, 4, 5}), 6)

	// All the same height: widest container wins, 5 * 3 = 15.
	check("all the same height", `maxArea([]int{5, 5, 5, 5})`,
		maxArea([]int{5, 5, 5, 5}), 15)

	// Two very tall walls next to each other beat the wide but low container.
	check("narrow but tall", `maxArea([]int{1, 100, 100, 1})`,
		maxArea([]int{1, 100, 100, 1}), 100)

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
