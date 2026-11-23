// Problem:    Network Delay Time
// Difficulty: Medium
// Link:       https://neetcode.io/problems/network-delay-time
// Source:     Data Structures & Algorithms/network-delay-time/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// A network has n nodes (numbered from 1). times[i] = [u, v, t] means a
// signal travels from node u to node v in t units of time (one direction
// only). A signal is sent from node k. Return the time until ALL nodes have
// received it, or -1 if some node never does.
//
// Example 1:
//
//	Input:  times = [[1, 2, 1], [2, 3, 1], [1, 4, 4], [3, 4, 1]], n = 4, k = 1
//	Output: 3
//
// Example 2:
//
//	Input:  times = [[1, 2, 1], [2, 3, 1]], n = 3, k = 2
//	Output: -1
//
// How to run (from the "local" folder):
//
//	go run ./090-network-delay-time
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea (Bellman-Ford): dist[x] = the shortest time we know so far to
// reach node x. At the start only k is reachable (time 0).
// Then repeat: look at every edge u -> v. If going through u is faster
// than what we know for v, update v.
// A shortest path uses at most n-1 edges, so after n-1 rounds all the
// times are final. The signal is everywhere when the SLOWEST node got it.
// Time: O(n * edges). (Dijkstra with a heap is faster, but needs more code.)
func networkDelayTime(times [][]int, n int, k int) int {
	const inf = 1000000000

	dist := make([]int, n+1) // nodes are numbered from 1
	for i := range dist {
		dist[i] = inf
	}
	dist[k] = 0

	for round := 0; round < n-1; round++ {
		for _, edge := range times {
			from, to, time := edge[0], edge[1], edge[2]
			if dist[from] != inf && dist[from]+time < dist[to] {
				dist[to] = dist[from] + time
			}
		}
	}

	// the answer is the slowest node
	slowest := 0
	for node := 1; node <= n; node++ {
		if dist[node] == inf {
			return -1 // this node never gets the signal
		}
		if dist[node] > slowest {
			slowest = dist[node]
		}
	}
	return slowest
}

// ===================== Test cases =====================

func main() {
	// An entry [u, v, t] means: a signal goes from node u to node v in time t.

	// Example 1 from the problem: 1 -> 2 -> 3 -> 4 takes 3 (faster than 1 -> 4 directly).
	check("example 1", `networkDelayTime([][]int{{1, 2, 1}, {2, 3, 1}, {1, 4, 4}, {3, 4, 1}}, 4, 1)`,
		networkDelayTime([][]int{{1, 2, 1}, {2, 3, 1}, {1, 4, 4}, {3, 4, 1}}, 4, 1), 3)

	// Example 2 from the problem: starting at 2, node 1 is never reached.
	check("example 2: unreachable node", `networkDelayTime([][]int{{1, 2, 1}, {2, 3, 1}}, 3, 2)`,
		networkDelayTime([][]int{{1, 2, 1}, {2, 3, 1}}, 3, 2), -1)

	// The classic example.
	check("classic", `networkDelayTime([][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 2)`,
		networkDelayTime([][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 2), 2)

	// One node: nothing to wait for.
	check("one node", `networkDelayTime([][]int{}, 1, 1)`,
		networkDelayTime([][]int{}, 1, 1), 0)

	// Two nodes.
	check("two nodes", `networkDelayTime([][]int{{1, 2, 1}}, 2, 1)`,
		networkDelayTime([][]int{{1, 2, 1}}, 2, 1), 1)

	// The edge only goes one way.
	check("wrong direction", `networkDelayTime([][]int{{1, 2, 1}}, 2, 2)`,
		networkDelayTime([][]int{{1, 2, 1}}, 2, 2), -1)

	// The detour 1 -> 3 -> 2 (time 2) beats the direct edge 1 -> 2 (time 5).
	check("detour is faster", `networkDelayTime([][]int{{1, 2, 5}, {1, 3, 1}, {3, 2, 1}}, 3, 1)`,
		networkDelayTime([][]int{{1, 2, 5}, {1, 3, 1}, {3, 2, 1}}, 3, 1), 2)

	// Edges in both directions with different times.
	check("both directions", `networkDelayTime([][]int{{1, 2, 1}, {2, 1, 3}}, 2, 2)`,
		networkDelayTime([][]int{{1, 2, 1}, {2, 1, 3}}, 2, 2), 3)

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
