// Problem:    Best Time to Buy and Sell Stock with Cooldown
// Difficulty: Medium
// Link:       https://neetcode.io/problems/buy-and-sell-crypto-with-cooldown
// Source:     Data Structures & Algorithms/buy-and-sell-crypto-with-cooldown/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// prices[i] is the price on day i. You may buy and sell as often as you
// like, but you can hold only one coin at a time, and after selling you
// must wait one day before buying again. Return the biggest profit.
//
// Example 1:
//
//	Input:  prices = [1, 3, 4, 0, 4]
//	Output: 6             (buy 1, sell 3, wait, buy 0, sell 4)
//
// Example 2:
//
//	Input:  prices = [1]
//	Output: 0
//
// How to run (from the "local" folder):
//
//	go run ./105-buy-and-sell-crypto-with-cooldown
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea: at the end of every day we are in one of three situations.
// Keep the best profit for each of them:
//
//	holding = we own a coin
//	sold    = we sold today (tomorrow is the cooldown day)
//	resting = we own nothing and did not sell today (free to buy tomorrow)
//
// Each day the new values come from yesterday's values:
//
//	holding: keep holding, OR buy today (only allowed after resting)
//	sold:    we were holding and sell today
//	resting: keep resting, OR yesterday we sold (cooldown is over)
//
// Time: O(n), no extra memory.
func maxProfit(prices []int) int {
	holding := -prices[0] // bought on the first day
	sold := 0
	resting := 0

	for i := 1; i < len(prices); i++ {
		price := prices[i]

		newHolding := holding
		if resting-price > newHolding {
			newHolding = resting - price // buy today
		}

		newSold := holding + price // sell today

		newResting := resting
		if sold > newResting {
			newResting = sold // yesterday's sale, cooldown done
		}

		holding, sold, resting = newHolding, newSold, newResting
	}

	// the best result is never "still holding a coin"
	if sold > resting {
		return sold
	}
	return resting
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: buy 1, sell 3, wait one day, buy 0, sell 4.
	check("example 1", `maxProfit([]int{1, 3, 4, 0, 4})`,
		maxProfit([]int{1, 3, 4, 0, 4}), 6)

	// Example 2 from the problem: one day, nothing to do.
	check("example 2: one day", `maxProfit([]int{1})`,
		maxProfit([]int{1}), 0)

	// The classic example: buy 1, sell 2, cooldown, buy 0, sell 2.
	check("classic", `maxProfit([]int{1, 2, 3, 0, 2})`,
		maxProfit([]int{1, 2, 3, 0, 2}), 3)

	// Prices only fall: do nothing.
	check("falling prices", `maxProfit([]int{5, 4, 3})`,
		maxProfit([]int{5, 4, 3}), 0)

	// Two days.
	check("two days, rising", `maxProfit([]int{1, 2})`,
		maxProfit([]int{1, 2}), 1)
	check("two days, falling", `maxProfit([]int{2, 1})`,
		maxProfit([]int{2, 1}), 0)

	// Prices only rise: buy on the first day, sell on the last.
	check("rising prices", `maxProfit([]int{1, 2, 4})`,
		maxProfit([]int{1, 2, 4}), 3)

	// Buy 1, sell 6, cooldown, ..., buy 0, sell 2.
	check("two trades", `maxProfit([]int{6, 1, 6, 4, 3, 0, 2})`,
		maxProfit([]int{6, 1, 6, 4, 3, 0, 2}), 7)

	// Tricky: selling at 4 would force a cooldown on the cheap day (2).
	// Holding from 1 to 7 is better.
	check("cooldown blocks the cheap day", `maxProfit([]int{1, 4, 2, 7})`,
		maxProfit([]int{1, 4, 2, 7}), 6)

	// Sell at 4 (not 5!), cooldown on the 5, buy 2, sell 9.
	check("sell one day early", `maxProfit([]int{2, 1, 4, 5, 2, 9, 7})`,
		maxProfit([]int{2, 1, 4, 5, 2, 9, 7}), 10)

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
