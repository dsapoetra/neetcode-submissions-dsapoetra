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
	check("example 1", `maxProduct([]int{1, 2, -3, 4})`,
		maxProduct([]int{1, 2, -3, 4}), 4)

	// Example 2 from the problem: negative times negative is positive.
	check("example 2: two negatives", `maxProduct([]int{-2, -1})`,
		maxProduct([]int{-2, -1}), 2)

	// One number.
	check("one positive", `maxProduct([]int{5})`,
		maxProduct([]int{5}), 5)
	check("one negative", `maxProduct([]int{-5})`,
		maxProduct([]int{-5}), -5)

	// The classic example: [2,3].
	check("classic", `maxProduct([]int{2, 3, -2, 4})`,
		maxProduct([]int{2, 3, -2, 4}), 6)

	// A zero resets the product: the best is the 0 itself.
	check("zero in the middle", `maxProduct([]int{-2, 0, -1})`,
		maxProduct([]int{-2, 0, -1}), 0)

	// The whole list: -2 * 3 * -4 = 24.
	check("two negatives far apart", `maxProduct([]int{-2, 3, -4})`,
		maxProduct([]int{-2, 3, -4}), 24)

	// Starts with a zero.
	check("zero first", `maxProduct([]int{0, 2})`,
		maxProduct([]int{0, 2}), 2)

	// Three negatives: use only two of them.
	check("three negatives", `maxProduct([]int{-1, -2, -3})`,
		maxProduct([]int{-1, -2, -3}), 6)

	// -2 * -4 * 3 = 24.
	check("mixed", `maxProduct([]int{2, -5, -2, -4, 3})`,
		maxProduct([]int{2, -5, -2, -4, 3}), 24)

	// 2 * 3 * -5 * -2 = 60, after the zero.
	check("after a zero", `maxProduct([]int{1, 0, -1, 2, 3, -5, -2})`,
		maxProduct([]int{1, 0, -1, 2, 3, -5, -2}), 60)

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
