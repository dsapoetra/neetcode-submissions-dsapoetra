// Problem: Two Sum
// Link:    https://neetcode.io/problems/two-integer-sum
// Source:  Data Structures & Algorithms/two-integer-sum/submission-9.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./two-integer-sum
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

func twoSum(nums []int, target int) []int {
	indexOf := make(map[int]int)

	for i, num := range nums {
		need := target - num

		j, ok := indexOf[need]
		if ok {
			return []int{j, i}
		}

		indexOf[num] = i
	}

	return nil
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: 3 + 4 = 7.
	check("example 1", `twoSum([]int{3, 4, 5, 6}, 7)`,
		twoSum([]int{3, 4, 5, 6}, 7), []int{0, 1})

	// Example 2 from the problem: 4 + 6 = 10 (first and last).
	check("example 2", `twoSum([]int{4, 5, 6}, 10)`,
		twoSum([]int{4, 5, 6}, 10), []int{0, 2})

	// Example 3 from the problem: the same value twice, two different indexes.
	check("example 3: same value twice", `twoSum([]int{5, 5}, 10)`,
		twoSum([]int{5, 5}, 10), []int{0, 1})

	// The pair is in the middle and at the end.
	check("pair at the end", `twoSum([]int{3, 2, 4}, 6)`,
		twoSum([]int{3, 2, 4}, 6), []int{1, 2})

	// Negative numbers.
	check("negative numbers", `twoSum([]int{-1, -2, -3, -4}, -6)`,
		twoSum([]int{-1, -2, -3, -4}, -6), []int{1, 3})

	// Target 0 with two zeros.
	check("zeros", `twoSum([]int{0, 4, 3, 0}, 0)`,
		twoSum([]int{0, 4, 3, 0}, 0), []int{0, 3})

	// The smallest possible input: exactly two numbers.
	check("only two numbers", `twoSum([]int{1, 2}, 3)`,
		twoSum([]int{1, 2}, 3), []int{0, 1})

	// A negative and a positive number.
	check("negative + positive", `twoSum([]int{-3, 4, 3, 90}, 0)`,
		twoSum([]int{-3, 4, 3, 90}, 0), []int{0, 2})

	// The answer is the last two numbers.
	check("last two numbers", `twoSum([]int{1, 5, 9, 13}, 22)`,
		twoSum([]int{1, 5, 9, 13}, 22), []int{2, 3})

	// No pair adds up to the target: this solution returns nothing.
	check("no answer", `twoSum([]int{1, 2, 3}, 100)`,
		twoSum([]int{1, 2, 3}, 100), []int{})

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
