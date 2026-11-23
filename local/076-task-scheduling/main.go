// Problem:    Task Scheduler
// Difficulty: Medium
// Link:       https://neetcode.io/problems/task-scheduling
// Source:     Data Structures & Algorithms/task-scheduling/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Each letter is a task; every task takes one unit of time. Between two
// tasks with the SAME letter there must be at least n units of time (filled
// with other tasks or with idle time). Return the shortest total time to
// finish all tasks.
//
// Example 1:
//
//	Input:  tasks = ["X", "X", "Y", "Y"], n = 2
//	Output: 5             (X, Y, idle, X, Y)
//
// Example 2:
//
//	Input:  tasks = ["A", "A", "A", "B", "C"], n = 3
//	Output: 9             (A, B, C, idle, A, idle, idle, idle, A)
//
// How to run (from the "local" folder):
//
//	go run ./076-task-scheduling
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea: the task that appears most often decides the schedule.
// Say A appears maxCount times and the cooldown is n:
//
//	A _ _ A _ _ A        (n = 2, maxCount = 3)
//
// That is (maxCount - 1) blocks of length (n + 1), plus the last A.
// If several tasks appear maxCount times, they all sit in the last block:
//
//	A B _ A B _ A B
//
// The other tasks fill the gaps. If there are more tasks than gaps, no
// idle time is needed at all and the answer is just the number of tasks.
// Time: O(n).
func leastInterval(tasks []byte, n int) int {
	var count [26]int
	for _, task := range tasks {
		count[task-'A']++
	}

	// the highest count, and how many tasks have that count
	maxCount := 0
	tasksWithMax := 0
	for _, c := range count {
		if c > maxCount {
			maxCount = c
			tasksWithMax = 1
		} else if c == maxCount {
			tasksWithMax++
		}
	}

	slots := (maxCount-1)*(n+1) + tasksWithMax
	if len(tasks) > slots {
		return len(tasks)
	}
	return slots
}

// ===================== Test cases =====================

func main() {
	// The tasks are written as a string: []byte("XXYY") is the tasks X, X, Y, Y.

	// Example 1 from the problem: X Y idle X Y.
	check("example 1", `leastInterval([]byte("XXYY"), 2)`,
		leastInterval([]byte("XXYY"), 2), 5)

	// Example 2 from the problem: A _ _ _ A _ _ _ A, with B and C in two of the gaps.
	check("example 2", `leastInterval([]byte("AAABC"), 3)`,
		leastInterval([]byte("AAABC"), 3), 9)

	// The classic example: A B idle A B idle A B.
	check("classic", `leastInterval([]byte("AAABBB"), 2)`,
		leastInterval([]byte("AAABBB"), 2), 8)

	// No cooldown: just do the tasks.
	check("no cooldown", `leastInterval([]byte("AAABBB"), 0)`,
		leastInterval([]byte("AAABBB"), 0), 6)

	// A longer cooldown: A B _ _ A B _ _ A B.
	check("longer cooldown", `leastInterval([]byte("AAABBB"), 3)`,
		leastInterval([]byte("AAABBB"), 3), 10)

	// One task.
	check("one task", `leastInterval([]byte("A"), 5)`,
		leastInterval([]byte("A"), 5), 1)

	// Only one kind of task: A _ _ A _ _ A _ _ A.
	check("one kind of task", `leastInterval([]byte("AAAA"), 2)`,
		leastInterval([]byte("AAAA"), 2), 10)

	// All different: no idle time needed.
	check("all different", `leastInterval([]byte("ABCDEF"), 2)`,
		leastInterval([]byte("ABCDEF"), 2), 6)

	// Enough other tasks to fill every gap: the answer is the number of tasks.
	check("gaps are all filled", `leastInterval([]byte("AAABBBCCDD"), 2)`,
		leastInterval([]byte("AAABBBCCDD"), 2), 10)

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
