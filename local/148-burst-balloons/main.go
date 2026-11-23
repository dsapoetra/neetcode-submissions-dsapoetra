// Problem:    Burst Balloons
// Difficulty: Hard
// Link:       https://neetcode.io/problems/burst-balloons
// Source:     Data Structures & Algorithms/burst-balloons/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Each balloon has a number. When you burst balloon i you get
// nums[i-1] * nums[i] * nums[i+1] coins (a missing neighbor counts as 1).
// Afterwards its two neighbors become neighbors of each other. Return the
// most coins you can collect by bursting all balloons.
//
// Example 1:
//
//	Input:  nums = [4, 2, 3, 7]
//	Output: 143
//
// How to run (from the "local" folder):
//
//	go run ./148-burst-balloons
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea: thinking about which balloon to burst FIRST is hard, because the
// neighbors change. Think about which balloon is burst LAST instead.
// First put a balloon with value 1 at both ends (they are never burst).
// best[left][right] = the most coins from bursting all balloons strictly
// BETWEEN left and right.
// If balloon k is the last one burst in that range, then at that moment
// its neighbors are exactly left and right, so we get
//
//	best[left][k] + best[k][right] + value[left]*value[k]*value[right]
//
// Try every k and keep the best. Short ranges are computed first.
// Time: O(n^3).
func maxCoins(nums []int) int {
	// balloons with a 1 on both ends
	balloons := []int{1}
	balloons = append(balloons, nums...)
	balloons = append(balloons, 1)
	n := len(balloons)

	best := make([][]int, n)
	for i := range best {
		best[i] = make([]int, n)
	}

	// gap = distance between left and right; start with the smallest ranges
	for gap := 2; gap < n; gap++ {
		for left := 0; left+gap < n; left++ {
			right := left + gap

			for k := left + 1; k < right; k++ {
				coins := best[left][k] + best[k][right] + balloons[left]*balloons[k]*balloons[right]
				if coins > best[left][right] {
					best[left][right] = coins
				}
			}
		}
	}
	return best[0][n-1]
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem.
	check("example 1", `maxCoins([]int{4, 2, 3, 7})`,
		maxCoins([]int{4, 2, 3, 7}), 143)

	// The classic example: burst 1, 5, 3, 8 in that order: 15 + 120 + 24 + 8.
	check("classic", `maxCoins([]int{3, 1, 5, 8})`,
		maxCoins([]int{3, 1, 5, 8}), 167)

	// Two balloons: burst the 1 first (1*1*5 = 5), then the 5 (1*5*1 = 5).
	check("two balloons", `maxCoins([]int{1, 5})`,
		maxCoins([]int{1, 5}), 10)

	// One balloon: 1 * value * 1.
	check("one balloon", `maxCoins([]int{5})`,
		maxCoins([]int{5}), 5)
	check("one balloon of 1", `maxCoins([]int{1})`,
		maxCoins([]int{1}), 1)

	// 1*2*2 = 4, then 1*2*1 = 2.
	check("two equal balloons", `maxCoins([]int{2, 2})`,
		maxCoins([]int{2, 2}), 6)

	// Burst the middle first: 1*2*3 = 6, then 1*1*3 = 3, then 3.
	check("three balloons", `maxCoins([]int{1, 2, 3})`,
		maxCoins([]int{1, 2, 3}), 12)

	// A balloon with 0 gives no coins, but it must still be burst.
	check("a zero", `maxCoins([]int{0, 5})`,
		maxCoins([]int{0, 5}), 5)

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
