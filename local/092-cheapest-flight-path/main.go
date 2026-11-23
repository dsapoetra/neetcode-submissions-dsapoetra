// Problem:    Cheapest Flights Within K Stops
// Difficulty: Medium
// Link:       https://neetcode.io/problems/cheapest-flight-path
// Source:     Data Structures & Algorithms/cheapest-flight-path/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// There are n cities (0 .. n-1). A flight [from, to, price] goes one way.
// Return the cheapest price to get from src to dst with at most k stops in
// between, or -1 if that is not possible.
//
// Example 1:
//
//	Input:  n = 4, flights = [[0, 1, 200], [1, 2, 100], [1, 3, 300], [2, 3, 100]],
//	        src = 0, dst = 3, k = 1
//	Output: 500           (0 -> 1 -> 3)
//
// Example 2:
//
//	Input:  n = 3, flights = [[1, 0, 100], [1, 2, 200], [0, 2, 100]],
//	        src = 1, dst = 2, k = 1
//	Output: 200
//
// How to run (from the "local" folder):
//
//	go run ./092-cheapest-flight-path
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea (Bellman-Ford with a limit): price[x] = the cheapest price we know
// to reach city x. At most k stops means at most k+1 flights.
// So do k+1 rounds; in each round every flight may be used ONCE.
// Important: a round must read the prices of the PREVIOUS round (a copy).
// Otherwise one round could chain several flights together.
// Time: O(k * flights).
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	const inf = 1000000000

	price := make([]int, n)
	for i := range price {
		price[i] = inf
	}
	price[src] = 0

	for round := 0; round <= k; round++ {
		// copy, so this round only extends routes from the round before
		next := make([]int, n)
		copy(next, price)

		for _, flight := range flights {
			from, to, cost := flight[0], flight[1], flight[2]
			if price[from] != inf && price[from]+cost < next[to] {
				next[to] = price[from] + cost
			}
		}
		price = next
	}

	if price[dst] == inf {
		return -1
	}
	return price[dst]
}

// ===================== Test cases =====================

func main() {
	// A flight is [from, to, price]. The call is
	// findCheapestPrice(cities, flights, start, destination, maxStops).

	// Example 1 from the problem: 0 -> 1 -> 3. (0 -> 1 -> 2 -> 3 is cheaper but has 2 stops.)
	check("example 1", `findCheapestPrice(4, [][]int{{0, 1, 200}, {1, 2, 100}, {1, 3, 300}, {2, 3, 100}}, 0, 3, 1)`,
		findCheapestPrice(4, [][]int{{0, 1, 200}, {1, 2, 100}, {1, 3, 300}, {2, 3, 100}}, 0, 3, 1), 500)

	// Example 2 from the problem: the direct flight 1 -> 2 costs the same as 1 -> 0 -> 2.
	check("example 2", `findCheapestPrice(3, [][]int{{1, 0, 100}, {1, 2, 200}, {0, 2, 100}}, 1, 2, 1)`,
		findCheapestPrice(3, [][]int{{1, 0, 100}, {1, 2, 200}, {0, 2, 100}}, 1, 2, 1), 200)

	// The classic example.
	check("classic", `findCheapestPrice(4, [][]int{{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200}}, 0, 3, 1)`,
		findCheapestPrice(4, [][]int{{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200}}, 0, 3, 1), 700)

	// With one stop the cheap route is allowed, with no stops it is not.
	check("one stop allowed", `findCheapestPrice(3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 1)`,
		findCheapestPrice(3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 1), 200)
	check("no stops allowed", `findCheapestPrice(3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 0)`,
		findCheapestPrice(3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 0), 500)

	// No route at all.
	check("no route", `findCheapestPrice(3, [][]int{{0, 1, 100}}, 0, 2, 1)`,
		findCheapestPrice(3, [][]int{{0, 1, 100}}, 0, 2, 1), -1)

	// A route exists, but it needs 2 stops.
	check("too many stops", `findCheapestPrice(4, [][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 1}}, 0, 3, 1)`,
		findCheapestPrice(4, [][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 1}}, 0, 3, 1), -1)
	check("exactly enough stops", `findCheapestPrice(4, [][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 1}}, 0, 3, 2)`,
		findCheapestPrice(4, [][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 1}}, 0, 3, 2), 3)

	// Tricky: with 0 stops only ONE flight may be used, so the chain
	// 0 -> 1 -> 2 -> 3 (price 3) is not allowed.
	check("one flight only", `findCheapestPrice(4, [][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 1}, {0, 3, 10}}, 0, 3, 0)`,
		findCheapestPrice(4, [][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 1}, {0, 3, 10}}, 0, 3, 0), 10)

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
