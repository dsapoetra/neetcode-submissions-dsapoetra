// Problem: Top K Frequent Elements
// Link:    https://neetcode.io/problems/top-k-elements-in-list
// Source:  Data Structures & Algorithms/top-k-elements-in-list/submission-0.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./top-k-elements-in-list
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

func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	freq := make([][]int, len(nums)+1)

	for _, num := range nums {
		count[num]++
	}
	for num, cnt := range count {
		freq[cnt] = append(freq[cnt], num)
	}

	res := []int{}
	for i := len(freq) - 1; i > 0; i-- {
		for _, num := range freq[i] {
			res = append(res, num)
			if len(res) == k {
				return res
			}
		}
	}
	return res
}

// ===================== Test cases =====================

func main() {
	// The answer may be in any order, so we sort it before comparing.

	// Example 1 from the problem: 3 appears 3 times, 2 appears 2 times.
	check("example 1", `sortedInts(topKFrequent([]int{1, 2, 2, 3, 3, 3}, 2))`,
		sortedInts(topKFrequent([]int{1, 2, 2, 3, 3, 3}, 2)), []int{2, 3})

	// Example 2 from the problem: only one different number.
	check("example 2", `sortedInts(topKFrequent([]int{7, 7}, 1))`,
		sortedInts(topKFrequent([]int{7, 7}, 1)), []int{7})

	// A single number.
	check("one number", `sortedInts(topKFrequent([]int{1}, 1))`,
		sortedInts(topKFrequent([]int{1}, 1)), []int{1})

	// The classic example.
	check("classic", `sortedInts(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))`,
		sortedInts(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)), []int{1, 2})

	// k is the number of different values: return all of them.
	check("k = all different values", `sortedInts(topKFrequent([]int{4, 4, 4, 5, 5, 6}, 3))`,
		sortedInts(topKFrequent([]int{4, 4, 4, 5, 5, 6}, 3)), []int{4, 5, 6})

	// Negative numbers.
	check("negative numbers", `sortedInts(topKFrequent([]int{-1, -1, 2}, 1))`,
		sortedInts(topKFrequent([]int{-1, -1, 2}, 1)), []int{-1})

	// The most frequent number is spread around the list.
	check("spread around", `sortedInts(topKFrequent([]int{1, 2, 3, 1, 2, 1}, 1))`,
		sortedInts(topKFrequent([]int{1, 2, 3, 1, 2, 1}, 1)), []int{1})

	// Two numbers tie for first place, k = 2 takes both.
	check("tie for first", `sortedInts(topKFrequent([]int{3, 3, 1, 1, 2}, 2))`,
		sortedInts(topKFrequent([]int{3, 3, 1, 1, 2}, 2)), []int{1, 3})

	report()
}

// ===================== Helpers for the tests =====================

// sortedInts sorts the numbers. We use it when the answer may be in any order.
func sortedInts(nums []int) []int {
	sort.Ints(nums)
	return nums
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
