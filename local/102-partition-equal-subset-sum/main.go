// Problem:    Partition Equal Subset Sum
// Difficulty: Medium
// Link:       https://neetcode.io/problems/partition-equal-subset-sum
// Source:     Data Structures & Algorithms/partition-equal-subset-sum/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Given a list of positive numbers, return true if it can be split into two
// groups that have the same sum.
//
// Example 1:
//
//	Input:  nums = [1, 2, 3, 4]
//	Output: true          ([1, 4] and [2, 3])
//
// Example 2:
//
//	Input:  nums = [1, 2, 3, 4, 5]
//	Output: false
//
// How to run (from the "local" folder):
//
//	go run ./102-partition-equal-subset-sum
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea: two parts with the same sum means each part is HALF of the total.
// So the question is: can some of the numbers add up to total / 2?
// possible[s] = true if some of the numbers seen so far add up to s.
// For each number, every sum s that was possible makes s + num possible.
// We go through the sums from big to small so that one number is not
// used twice in the same round.
// Time: O(n * total / 2).
func canPartition(nums []int) bool {
	// add up all the numbers
	total := 0
	for _, num := range nums {
		total += num
	}
	if total%2 == 1 {
		return false // an odd total cannot be split into two equal parts
	}
	half := total / 2

	possible := make([]bool, half+1)
	possible[0] = true // taking nothing gives sum 0

	for _, num := range nums {
		// sum - num was possible before, so with this number sum is possible.
		// Sums smaller than num cannot use this number, so we stop there.
		for sum := half; sum >= num; sum-- {
			if possible[sum-num] {
				possible[sum] = true
			}
		}
	}
	// can some numbers add up to exactly half? Then the rest is the other half
	return possible[half]
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: [1,4] and [2,3].
	check("example 1", `canPartition([]int{1, 2, 3, 4})`,
		canPartition([]int{1, 2, 3, 4}), true)

	// Example 2 from the problem: the total (15) is odd.
	check("example 2: odd total", `canPartition([]int{1, 2, 3, 4, 5})`,
		canPartition([]int{1, 2, 3, 4, 5}), false)

	// The classic examples: [1,5,5] and [11].
	check("classic true", `canPartition([]int{1, 5, 11, 5})`,
		canPartition([]int{1, 5, 11, 5}), true)
	check("classic false", `canPartition([]int{1, 2, 3, 5})`,
		canPartition([]int{1, 2, 3, 5}), false)

	// Two equal numbers.
	check("two equal numbers", `canPartition([]int{1, 1})`,
		canPartition([]int{1, 1}), true)

	// One number cannot be split.
	check("one number", `canPartition([]int{1})`,
		canPartition([]int{1}), false)
	check("one even number", `canPartition([]int{100})`,
		canPartition([]int{100}), false)

	// The total (12) is even, but no group adds up to 6.
	check("even total, still impossible", `canPartition([]int{2, 2, 3, 5})`,
		canPartition([]int{2, 2, 3, 5}), false)

	// One number is bigger than half of the total.
	check("one number too big", `canPartition([]int{1, 2, 5})`,
		canPartition([]int{1, 2, 5}), false)

	// 14 + 4 + 2 = 20 and 9 + 8 + 3 = 20.
	check("bigger case", `canPartition([]int{14, 9, 8, 4, 3, 2})`,
		canPartition([]int{14, 9, 8, 4, 3, 2}), true)

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
