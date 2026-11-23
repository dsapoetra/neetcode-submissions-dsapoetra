// Problem: Maximum Product Subarray
// Link:    https://neetcode.io/problems/maximum-product-subarray
// Source:  Data Structures & Algorithms/maximum-product-subarray/submission-0.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./maximum-product-subarray
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea: like Maximum Subarray, but with multiplication a negative number
// can turn the SMALLEST product into the BIGGEST one (negative * negative).
// So for the subarrays ending at the current number we remember both
// the biggest product (high) and the smallest product (low).
// Time: O(n).
func maxProduct(nums []int) int {
	best := nums[0]
	high := 1 // biggest product of a subarray ending at the previous number
	low := 1  // smallest product of a subarray ending at the previous number

	for _, n := range nums {
		// three choices for a subarray ending here:
		// start fresh with n, or extend high, or extend low
		a := high * n
		b := low * n

		high = n
		if a > high {
			high = a
		}
		if b > high {
			high = b
		}

		low = n
		if a < low {
			low = a
		}
		if b < low {
			low = b
		}

		if high > best {
			best = high
		}
	}
	return best
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: just [4].
	check("example 1", maxProduct([]int{1, 2, -3, 4}), 4)

	// Example 2 from the problem: negative times negative is positive.
	check("example 2: two negatives", maxProduct([]int{-2, -1}), 2)

	// One number.
	check("one positive", maxProduct([]int{5}), 5)
	check("one negative", maxProduct([]int{-5}), -5)

	// The classic example: [2,3].
	check("classic", maxProduct([]int{2, 3, -2, 4}), 6)

	// A zero resets the product: the best is the 0 itself.
	check("zero in the middle", maxProduct([]int{-2, 0, -1}), 0)

	// The whole list: -2 * 3 * -4 = 24.
	check("two negatives far apart", maxProduct([]int{-2, 3, -4}), 24)

	// Starts with a zero.
	check("zero first", maxProduct([]int{0, 2}), 2)

	// Three negatives: use only two of them.
	check("three negatives", maxProduct([]int{-1, -2, -3}), 6)

	// -2 * -4 * 3 = 24.
	check("mixed", maxProduct([]int{2, -5, -2, -4, 3}), 24)

	// 2 * 3 * -5 * -2 = 60, after the zero.
	check("after a zero", maxProduct([]int{1, 0, -1, 2, 3, -5, -2}), 60)

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
