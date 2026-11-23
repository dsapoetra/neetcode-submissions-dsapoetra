// Problem:    Time Based Key-Value Store
// Difficulty: Medium
// Link:       https://neetcode.io/problems/time-based-key-value-store
// Source:     Data Structures & Algorithms/time-based-key-value-store/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Build a key-value store that remembers the value of a key at different
// times:
//
//	Set(key, value, timestamp)  stores the value for that time
//	Get(key, timestamp)         the value with the latest time that is
//	                            <= timestamp, or "" if there is none
//
// Set is always called with increasing timestamps.
//
// Example:
//
//	Input:  Set("alice", "happy", 1), Get("alice", 1), Get("alice", 2),
//	        Set("alice", "sad", 3), Get("alice", 3)
//	Output: -, "happy", "happy", -, "sad"
//
// How to run (from the "local" folder):
//
//	go run ./045-time-based-key-value-store
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea: for every key keep a list of (timestamp, value) entries.
// The problem promises that Set is called with increasing timestamps, so
// each list is already sorted by time.
// Get must find the LAST entry whose timestamp is <= the asked timestamp,
// which is a binary search on that sorted list.
// Set: O(1). Get: O(log n).

type timeEntry struct {
	timestamp int
	value     string
}

type TimeMap struct {
	entries map[string][]timeEntry
}

func Constructor() TimeMap {
	return TimeMap{entries: map[string][]timeEntry{}}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.entries[key] = append(this.entries[key], timeEntry{timestamp, value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	list := this.entries[key]
	answer := "" // stays empty if nothing is old enough

	left := 0
	right := len(list) - 1
	for left <= right {
		mid := (left + right) / 2
		if list[mid].timestamp <= timestamp {
			answer = list[mid].value // a candidate; look for a later one
			left = mid + 1
		} else {
			right = mid - 1 // too new
		}
	}
	return answer
}

// ===================== Test cases =====================

func main() {
	// The store remembers what was set before, so the lines starting with
	// "--" show every Set in the output.
	store := Constructor()

	// Nothing was set yet.
	check("empty store", `store.Get("alice", 1)`,
		store.Get("alice", 1), "")

	// Example from the problem.
	store.Set("alice", "happy", 1)
	fmt.Println(`-- store.Set("alice", "happy", 1)`)
	check("exact timestamp", `store.Get("alice", 1)`,
		store.Get("alice", 1), "happy")
	check("later timestamp: the latest value before it", `store.Get("alice", 2)`,
		store.Get("alice", 2), "happy")

	store.Set("alice", "sad", 3)
	fmt.Println(`-- store.Set("alice", "sad", 3)`)
	check("new value at its timestamp", `store.Get("alice", 3)`,
		store.Get("alice", 3), "sad")

	// The old value is still there for older timestamps.
	check("between the two values", `store.Get("alice", 2)`,
		store.Get("alice", 2), "happy")

	// Before the first value.
	check("before the first value", `store.Get("alice", 0)`,
		store.Get("alice", 0), "")

	// Far in the future.
	check("far in the future", `store.Get("alice", 100)`,
		store.Get("alice", 100), "sad")

	// A key that was never set.
	check("unknown key", `store.Get("bob", 5)`,
		store.Get("bob", 5), "")

	// Keys do not affect each other.
	store.Set("bob", "tired", 2)
	fmt.Println(`-- store.Set("bob", "tired", 2)`)
	check("second key", `store.Get("bob", 2)`,
		store.Get("bob", 2), "tired")
	check("second key, too early", `store.Get("bob", 1)`,
		store.Get("bob", 1), "")
	check("first key unchanged", `store.Get("alice", 2)`,
		store.Get("alice", 2), "happy")

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
