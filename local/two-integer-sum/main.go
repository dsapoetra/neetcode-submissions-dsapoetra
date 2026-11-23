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
	check("example 1", twoSum([]int{3, 4, 5, 6}, 7), []int{0, 1})

	// Example 2 from the problem: 4 + 6 = 10 (first and last).
	check("example 2", twoSum([]int{4, 5, 6}, 10), []int{0, 2})

	// Example 3 from the problem: the same value twice, two different indexes.
	check("example 3: same value twice", twoSum([]int{5, 5}, 10), []int{0, 1})

	// The pair is in the middle and at the end.
	check("pair at the end", twoSum([]int{3, 2, 4}, 6), []int{1, 2})

	// Negative numbers.
	check("negative numbers", twoSum([]int{-1, -2, -3, -4}, -6), []int{1, 3})

	// Target 0 with two zeros.
	check("zeros", twoSum([]int{0, 4, 3, 0}, 0), []int{0, 3})

	// The smallest possible input: exactly two numbers.
	check("only two numbers", twoSum([]int{1, 2}, 3), []int{0, 1})

	// A negative and a positive number.
	check("negative + positive", twoSum([]int{-3, 4, 3, 90}, 0), []int{0, 2})

	// The answer is the last two numbers.
	check("last two numbers", twoSum([]int{1, 5, 9, 13}, 22), []int{2, 3})

	// No pair adds up to the target: this solution returns nothing.
	check("no answer", twoSum([]int{1, 2, 3}, 100), []int{})

	report()
}

// ===================== Helpers for the tests =====================

// passed and failed count the test results.
var passed, failed int

// check compares the answer we got with the answer we want, and prints
// PASS or FAIL for that test case.
func check(name string, got interface{}, want interface{}) {
	if fmt.Sprint(got) == fmt.Sprint(want) {
		passed++
		fmt.Printf("PASS  %s\n", name)
	} else {
		failed++
		fmt.Printf("FAIL  %s\n      got:  %v\n      want: %v\n", name, got, want)
	}
}

// report prints the total. The program exits with an error if a test failed.
func report() {
	fmt.Printf("\n%d passed, %d failed\n", passed, failed)
	if failed > 0 {
		os.Exit(1)
	}
}
