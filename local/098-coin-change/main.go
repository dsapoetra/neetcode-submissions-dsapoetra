// Problem:    Coin Change
// Difficulty: Medium
// Link:       https://neetcode.io/problems/coin-change
// Source:     Data Structures & Algorithms/coin-change/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Given coin values and an amount, return the fewest coins needed to make
// exactly that amount. Every coin can be used as often as you like.
// Return -1 if the amount cannot be made.
//
// Example 1:
//
//	Input:  coins = [1, 5, 10], amount = 12
//	Output: 3             (10 + 1 + 1)
//
// Example 2:
//
//	Input:  coins = [2], amount = 3
//	Output: -1
//
// Example 3:
//
//	Input:  coins = [1], amount = 0
//	Output: 0
//
// How to run (from the "local" folder):
//
//	go run ./098-coin-change
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea (dynamic programming): dp[a] = fewest coins needed to make amount a.
// To make amount a, try every coin c as the last coin:
//
//	dp[a] = dp[a-c] + 1   (take the smallest over all coins)
//
// Time: O(amount * number of coins).
func coinChange(coins []int, amount int) int {
	// amount+1 is used as "impossible" because we never need more
	// than amount coins (the smallest coin is 1).
	impossible := amount + 1

	dp := make([]int, amount+1)
	for a := 1; a <= amount; a++ {
		dp[a] = impossible
	}
	// dp[0] is already 0: zero coins make amount 0

	for a := 1; a <= amount; a++ {
		for _, c := range coins {
			// coin c can be the last coin only if it is not bigger than a
			if c <= a && dp[a-c]+1 < dp[a] {
				dp[a] = dp[a-c] + 1
			}
		}
	}

	if dp[amount] == impossible {
		return -1
	}
	return dp[amount]
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: 10 + 1 + 1.
	check("example 1", `coinChange([]int{1, 5, 10}, 12)`,
		coinChange([]int{1, 5, 10}, 12), 3)

	// Example 2 from the problem: 3 cannot be made from 2s.
	check("example 2: impossible", `coinChange([]int{2}, 3)`,
		coinChange([]int{2}, 3), -1)

	// Example 3 from the problem: amount 0 needs no coins.
	check("example 3: amount 0", `coinChange([]int{1}, 0)`,
		coinChange([]int{1}, 0), 0)

	// The classic example: 5 + 5 + 1.
	check("classic", `coinChange([]int{1, 2, 5}, 11)`,
		coinChange([]int{1, 2, 5}, 11), 3)

	// The coins are not sorted: 10 + 10 + 5 + 2.
	check("unsorted coins", `coinChange([]int{2, 5, 10, 1}, 27)`,
		coinChange([]int{2, 5, 10, 1}, 27), 4)

	// One coin is exactly the amount.
	check("one coin is enough", `coinChange([]int{5}, 5)`,
		coinChange([]int{5}, 5), 1)

	// Tricky: taking the biggest coin first (4+1+1 = 3 coins) is NOT best.
	// 3 + 3 = 2 coins. This is why we need dynamic programming.
	check("greedy is wrong here", `coinChange([]int{1, 3, 4}, 6)`,
		coinChange([]int{1, 3, 4}, 6), 2)

	// Amount is smaller than every coin.
	check("amount smaller than all coins", `coinChange([]int{2}, 1)`,
		coinChange([]int{2}, 1), -1)

	// A big case to make sure it is fast enough.
	check("big case", `coinChange([]int{186, 419, 83, 408}, 6249)`,
		coinChange([]int{186, 419, 83, 408}, 6249), 20)

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
