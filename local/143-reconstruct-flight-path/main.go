// Problem:    Reconstruct Itinerary
// Difficulty: Hard
// Link:       https://neetcode.io/problems/reconstruct-flight-path
// Source:     Data Structures & Algorithms/reconstruct-flight-path/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Each ticket is [from, to]. Starting at "JFK", find the route that uses
// every ticket exactly once. If several routes are possible, return the one
// that comes first in alphabetical order. A valid route always exists.
//
// Example 1:
//
//	Input:  tickets = [["BUF", "HOU"], ["HOU", "SEA"], ["JFK", "BUF"]]
//	Output: ["JFK", "BUF", "HOU", "SEA"]
//
// Example 2:
//
//	Input:  tickets = [["HOU", "JFK"], ["SEA", "JFK"], ["JFK", "SEA"], ["JFK", "HOU"]]
//	Output: ["JFK", "HOU", "JFK", "SEA", "JFK"]
//
// How to run (from the "local" folder):
//
//	go run ./143-reconstruct-flight-path
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

// Idea: we must use every ticket exactly once, and prefer the smallest
// airport name whenever there is a choice.
//  1. for every airport keep its destinations sorted, biggest name first,
//     so the SMALLEST name is at the end and cheap to take off
//  2. fly (DFS): from an airport, keep taking the smallest unused ticket.
//     Only when an airport has no tickets left do we write it down.
//
// That writes the route down BACKWARDS (dead ends first), so we reverse
// it at the end.
// Time: O(tickets * log tickets).
func findItinerary(tickets [][]string) []string {
	// 1) destinations per airport, sorted so the smallest name is last
	sort.Slice(tickets, func(a, b int) bool {
		return tickets[a][1] > tickets[b][1]
	})
	destinations := map[string][]string{}
	for _, ticket := range tickets {
		destinations[ticket[0]] = append(destinations[ticket[0]], ticket[1])
	}

	route := []string{}

	// 2) use up all tickets leaving from this airport, then write it down
	var fly func(airport string)
	fly = func(airport string) {
		for len(destinations[airport]) > 0 {
			list := destinations[airport]
			next := list[len(list)-1] // smallest name
			destinations[airport] = list[:len(list)-1]
			fly(next)
		}
		route = append(route, airport)
	}
	fly("JFK")

	// the route was written backwards
	for i, j := 0, len(route)-1; i < j; i, j = i+1, j-1 {
		route[i], route[j] = route[j], route[i]
	}
	return route
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: only one possible route.
	check("example 1", `findItinerary([][]string{{"BUF", "HOU"}, {"HOU", "SEA"}, {"JFK", "BUF"}})`,
		findItinerary([][]string{{"BUF", "HOU"}, {"HOU", "SEA"}, {"JFK", "BUF"}}), []string{"JFK", "BUF", "HOU", "SEA"})

	// Example 2 from the problem: from JFK, "HOU" comes before "SEA" in the alphabet.
	check("example 2", `findItinerary([][]string{{"HOU", "JFK"}, {"SEA", "JFK"}, {"JFK", "SEA"}, {"JFK", "HOU"}})`,
		findItinerary([][]string{{"HOU", "JFK"}, {"SEA", "JFK"}, {"JFK", "SEA"}, {"JFK", "HOU"}}), []string{"JFK", "HOU", "JFK", "SEA", "JFK"})

	// The classic examples.
	check("classic 1", `findItinerary([][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}})`,
		findItinerary([][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}}), []string{"JFK", "MUC", "LHR", "SFO", "SJC"})
	check("classic 2", `findItinerary([][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}})`,
		findItinerary([][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}}), []string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"})

	// Tricky: "KUL" is smaller than "NRT", but flying to KUL first is a
	// dead end (the other tickets could not be used). So NRT must come first.
	check("smallest name is a dead end", `findItinerary([][]string{{"JFK", "KUL"}, {"JFK", "NRT"}, {"NRT", "JFK"}})`,
		findItinerary([][]string{{"JFK", "KUL"}, {"JFK", "NRT"}, {"NRT", "JFK"}}), []string{"JFK", "NRT", "JFK", "KUL"})

	// One ticket.
	check("one ticket", `findItinerary([][]string{{"JFK", "AAA"}})`,
		findItinerary([][]string{{"JFK", "AAA"}}), []string{"JFK", "AAA"})

	// The same ticket twice.
	check("same ticket twice", `findItinerary([][]string{{"JFK", "ANU"}, {"ANU", "JFK"}, {"JFK", "ANU"}})`,
		findItinerary([][]string{{"JFK", "ANU"}, {"ANU", "JFK"}, {"JFK", "ANU"}}), []string{"JFK", "ANU", "JFK", "ANU"})

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
