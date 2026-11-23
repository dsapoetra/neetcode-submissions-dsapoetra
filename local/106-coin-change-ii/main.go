// Problem:    Coin Change II
// Difficulty: Medium
// Link:       https://neetcode.io/problems/coin-change-ii
// Source:     Data Structures & Algorithms/coin-change-ii/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Given coin values and an amount, return the number of different
// combinations of coins that make exactly that amount. Every coin can be
// used as often as you like; the order of the coins does not matter.
//
// Example 1:
//
//	Input:  amount = 4, coins = [1, 2, 3]
//	Output: 4             (1+1+1+1, 1+1+2, 2+2, 1+3)
//
// Example 2:
//
//	Input:  amount = 7, coins = [2, 4]
//	Output: 0
//
// How to run (from the "local" folder):
//
//	go run ./106-coin-change-ii
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea (dynamic programming): ways[a] = number of ways to make amount a.
// Add the coins ONE BY ONE. With a new coin, every way to make
// (a - coin) becomes a way to make a, by adding this coin.
// Because the coin loop is the OUTER loop, each combination is counted
// once (1+2 and 2+1 are the same combination, not two).
// Time: O(coins * amount).
func change(amount int, coins []int) int {
	ways := make([]int, amount+1)
	ways[0] = 1 // one way to make 0: use no coins

	for _, coin := range coins {
		for a := coin; a <= amount; a++ {
			ways[a] += ways[a-coin]
		}
	}
	return ways[amount]
}

// ===================== Test cases =====================

func main() {
	// The call is change(amount, coins).

	// Example 1 from the problem: 1+1+1+1, 1+1+2, 2+2, 1+3.
	check("example 1", `change(4, []int{1, 2, 3})`,
		change(4, []int{1, 2, 3}), 4)

	// Example 2 from the problem: 7 is odd, the coins are even.
	check("example 2: impossible", `change(7, []int{2, 4})`,
		change(7, []int{2, 4}), 0)

	// The classic examples.
	check("classic", `change(5, []int{1, 2, 5})`,
		change(5, []int{1, 2, 5}), 4)
	check("classic impossible", `change(3, []int{2})`,
		change(3, []int{2}), 0)
	check("classic one coin", `change(10, []int{10})`,
		change(10, []int{10}), 1)

	// Amount 0: one way (use no coins).
	check("amount 0", `change(0, []int{1, 2})`,
		change(0, []int{1, 2}), 1)

	// 1+1+1+1+1+1 or 1+5.
	check("two ways", `change(6, []int{1, 5})`,
		change(6, []int{1, 5}), 2)

	// Making change for 1 dollar with US coins: a well-known answer.
	check("100 with 1, 5, 10, 25, 50", `change(100, []int{1, 5, 10, 25, 50})`,
		change(100, []int{1, 5, 10, 25, 50}), 292)

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
