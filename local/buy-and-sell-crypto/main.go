// Problem: Best Time to Buy and Sell Stock
// Link:    https://neetcode.io/problems/buy-and-sell-crypto
// Source:  Data Structures & Algorithms/buy-and-sell-crypto/submission-5.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./buy-and-sell-crypto
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea: walk through the prices once. Remember the cheapest price seen
// so far (the best day to buy). For every day, check the profit if we
// sold today, and keep the biggest profit.
// Time: O(n).
func maxProfit(prices []int) int {
	minPrice := prices[0] // cheapest price so far
	best := 0             // best profit so far (0 = do not trade)

	for _, price := range prices {
		if price < minPrice {
			minPrice = price // found a cheaper day to buy
		}
		profit := price - minPrice // buy at the cheapest, sell today
		if profit > best {
			best = profit
		}
	}
	return best
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: buy at 1, sell at 7.
	check("example 1", maxProfit([]int{10, 1, 5, 6, 7, 1}), 6)

	// Example 2 from the problem: the price only goes down, so do not trade.
	check("example 2: only falling", maxProfit([]int{10, 8, 7, 5, 2}), 0)

	// One day: we cannot sell after buying.
	check("one day", maxProfit([]int{5}), 0)

	// The price only goes up: buy first day, sell last day.
	check("only rising", maxProfit([]int{1, 2, 3, 4, 5}), 4)

	// The classic example: buy at 1, sell at 6.
	check("classic", maxProfit([]int{7, 1, 5, 3, 6, 4}), 5)

	// The price never changes.
	check("flat price", maxProfit([]int{2, 2, 2}), 0)

	// The lowest price comes AFTER the best trade (buy 3, sell 8).
	check("lowest price too late", maxProfit([]int{3, 8, 1, 2}), 5)

	// A new lowest price gives a better trade later (buy 1, sell 7).
	check("better trade after new low", maxProfit([]int{2, 4, 1, 7}), 6)

	// We must buy BEFORE we sell: the big 9 at the start is useless.
	check("high price first", maxProfit([]int{9, 1, 2}), 1)

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
