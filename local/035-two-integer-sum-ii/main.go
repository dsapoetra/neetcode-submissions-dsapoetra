// Problem:    Two Sum II - Input Array Is Sorted
// Difficulty: Medium
// Link:       https://neetcode.io/problems/two-integer-sum-ii
// Source:     Data Structures & Algorithms/two-integer-sum-ii/submission-8.go (the solution below is a copy of it)
//
// THE PROBLEM
// Given a SORTED list of numbers and a target, find the two numbers that add
// up to the target and return their positions, counted from 1 (smaller
// position first). There is exactly one answer. Use only O(1) extra memory.
//
// Example 1:
//
//	Input:  numbers = [1, 2, 3, 4], target = 3
//	Output: [1, 2]        (1 + 2 = 3)
//
// How to run (from the "local" folder):
//
//	go run ./035-two-integer-sum-ii
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea (two pointers): the list is sorted. Start with the smallest and the
// biggest number.
//   - sum too small -> move the left pointer right (bigger number)
//   - sum too big   -> move the right pointer left (smaller number)
//
// The answer uses positions that start at 1, not 0.
// Time: O(n), no extra memory.
func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1} // positions start at 1
		}
		if sum < target {
			left++
		} else {
			right--
		}
	}
	return []int{}
}

// ===================== Test cases =====================

func main() {
	// The answer uses positions that start at 1.

	// Example 1 from the problem: 1 + 2 = 3.
	check("example 1", `twoSum([]int{1, 2, 3, 4}, 3)`,
		twoSum([]int{1, 2, 3, 4}, 3), []int{1, 2})

	// The classic example: 2 + 7 = 9.
	check("classic", `twoSum([]int{2, 7, 11, 15}, 9)`,
		twoSum([]int{2, 7, 11, 15}, 9), []int{1, 2})

	// First and last number.
	check("first and last", `twoSum([]int{2, 3, 4}, 6)`,
		twoSum([]int{2, 3, 4}, 6), []int{1, 3})

	// Negative numbers.
	check("negative numbers", `twoSum([]int{-1, 0}, -1)`,
		twoSum([]int{-1, 0}, -1), []int{1, 2})

	// The pair is in the middle: 4 + 5 = 9.
	check("pair in the middle", `twoSum([]int{1, 3, 4, 5, 7, 10, 11}, 9)`,
		twoSum([]int{1, 3, 4, 5, 7, 10, 11}, 9), []int{3, 4})

	// A negative target.
	check("negative target", `twoSum([]int{-5, -3, 0, 2, 8}, -8)`,
		twoSum([]int{-5, -3, 0, 2, 8}, -8), []int{1, 2})

	// The same value twice.
	check("same value twice", `twoSum([]int{1, 1}, 2)`,
		twoSum([]int{1, 1}, 2), []int{1, 2})

	// The last two numbers.
	check("last two", `twoSum([]int{5, 25, 75}, 100)`,
		twoSum([]int{5, 25, 75}, 100), []int{2, 3})

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
