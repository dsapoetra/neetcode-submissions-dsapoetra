// Problem: 3Sum
// Link:    https://neetcode.io/problems/three-integer-sum
// Source:  Data Structures & Algorithms/three-integer-sum/submission-4.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./three-integer-sum
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
	"sort"
)

// ===================== Solution =====================

// Idea: sort the numbers. Then fix the first number nums[i] and find
// the other two with two pointers (left and right) that move towards
// each other:
//   - sum too small -> move left forward (bigger number)
//   - sum too big   -> move right back (smaller number)
//   - sum is 0      -> save the triplet
//
// Skipping equal numbers avoids duplicate triplets.
// Time: O(n^2).
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		// same first number as before -> would give the same triplets
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum < 0 {
				left++ // need a bigger sum
			} else if sum > 0 {
				right-- // need a smaller sum
			} else {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				// skip the same second number
				for left < right && nums[left] == nums[left-1] {
					left++
				}
			}
		}
	}
	return res
}

// ===================== Test cases =====================

func main() {
	// The triplets may come back in any order, so we sort them before comparing.

	// Example 1 from the problem.
	check("example 1", `sortRows(threeSum([]int{-1, 0, 1, 2, -1, -4}))`,
		sortRows(threeSum([]int{-1, 0, 1, 2, -1, -4})), [][]int{{-1, -1, 2}, {-1, 0, 1}})

	// Example 2 from the problem: no triplet adds up to 0.
	check("example 2: no triplet", `sortRows(threeSum([]int{0, 1, 1}))`,
		sortRows(threeSum([]int{0, 1, 1})), [][]int{})

	// Example 3 from the problem: three zeros.
	check("example 3: three zeros", `sortRows(threeSum([]int{0, 0, 0}))`,
		sortRows(threeSum([]int{0, 0, 0})), [][]int{{0, 0, 0}})

	// Four zeros must still give only ONE triplet (no duplicates).
	check("four zeros, one triplet", `sortRows(threeSum([]int{0, 0, 0, 0}))`,
		sortRows(threeSum([]int{0, 0, 0, 0})), [][]int{{0, 0, 0}})

	// Only positive numbers can never add up to 0.
	check("all positive", `sortRows(threeSum([]int{1, 2, 3}))`,
		sortRows(threeSum([]int{1, 2, 3})), [][]int{})

	// Two different triplets share the same first number.
	check("same first number", `sortRows(threeSum([]int{-2, 0, 1, 1, 2}))`,
		sortRows(threeSum([]int{-2, 0, 1, 1, 2})), [][]int{{-2, 0, 2}, {-2, 1, 1}})

	// Exactly three numbers that work.
	check("exactly three numbers", `sortRows(threeSum([]int{-1, 0, 1}))`,
		sortRows(threeSum([]int{-1, 0, 1})), [][]int{{-1, 0, 1}})

	// Fewer than three numbers: no triplet is possible.
	check("too few numbers", `sortRows(threeSum([]int{1, -1}))`,
		sortRows(threeSum([]int{1, -1})), [][]int{})

	// Many duplicates: each triplet must appear only once.
	check("many duplicates", `sortRows(threeSum([]int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}))`,
		sortRows(threeSum([]int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6})), [][]int{{-4, -2, 6}, {-4, 0, 4}, {-4, 1, 3}, {-4, 2, 2}, {-2, -2, 4}, {-2, 0, 2}})

	report()
}

// ===================== Helpers for the tests =====================

// sortRows sorts the numbers inside every row and then sorts the rows.
// We use it when the problem says "return the answer in any order":
// after sorting, two correct answers always look the same.
func sortRows(rows [][]int) [][]int {
	for _, row := range rows {
		sort.Ints(row)
	}
	sort.Slice(rows, func(a, b int) bool {
		for i := 0; i < len(rows[a]) && i < len(rows[b]); i++ {
			if rows[a][i] != rows[b][i] {
				return rows[a][i] < rows[b][i]
			}
		}
		return len(rows[a]) < len(rows[b])
	})
	return rows
}

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
