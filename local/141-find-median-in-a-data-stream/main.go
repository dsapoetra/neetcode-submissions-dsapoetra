// Problem:    Find Median from Data Stream
// Difficulty: Hard
// Link:       https://neetcode.io/problems/find-median-in-a-data-stream
// Source:     Data Structures & Algorithms/find-median-in-a-data-stream/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Build a structure that receives numbers one by one and can report the
// median of all numbers so far:
//
//	AddNum(num)    adds a number
//	FindMedian()   the middle value; with an even count, the average of the
//	               two middle values
//
// Example:
//
//	Input:  AddNum(1), FindMedian(), AddNum(2), FindMedian(), AddNum(3), FindMedian()
//	Output: -, 1.0, -, 1.5, -, 2.0
//
// How to run (from the "local" folder):
//
//	go run ./141-find-median-in-a-data-stream
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea: keep all numbers in a slice that is ALWAYS sorted.
// AddNum puts the new number at the end and moves it left until it is
// in the right place (one step of insertion sort).
// FindMedian just reads the middle of the sorted slice.
//
// AddNum is O(n), FindMedian is O(1). This is the simple version.
// The faster version uses two heaps (O(log n) per AddNum), but it is
// much more code.

type MedianFinder struct {
	nums []int // always sorted
}

func Constructor() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
	this.nums = append(this.nums, num)

	// move the new number left while it is smaller than its left neighbor
	i := len(this.nums) - 1
	for i > 0 && this.nums[i] < this.nums[i-1] {
		this.nums[i], this.nums[i-1] = this.nums[i-1], this.nums[i]
		i--
	}
}

func (this *MedianFinder) FindMedian() float64 {
	n := len(this.nums)
	if n%2 == 1 {
		// odd count: the middle number
		return float64(this.nums[n/2])
	}
	// even count: average of the two middle numbers
	return float64(this.nums[n/2-1]+this.nums[n/2]) / 2.0
}

// ===================== Test cases =====================

func main() {
	// medianOf creates a new MedianFinder, adds the numbers one by one (in
	// the order given) and returns the median at the end.
	medianOf := func(nums ...int) float64 {
		finder := Constructor()
		for _, num := range nums {
			finder.AddNum(num)
		}
		return finder.FindMedian()
	}

	// Example from the problem, step by step.
	check("one number: the number itself", `medianOf(1)`,
		medianOf(1), 1.0)
	check("two numbers: the average", `medianOf(1, 2)`,
		medianOf(1, 2), 1.5)
	check("three numbers: the middle one", `medianOf(1, 2, 3)`,
		medianOf(1, 2, 3), 2.0)

	// Numbers arriving from big to small.
	check("descending order", `medianOf(5, 4, 3, 2)`,
		medianOf(5, 4, 3, 2), 3.5)

	// Negative numbers.
	check("two negatives", `medianOf(-1, -2)`,
		medianOf(-1, -2), -1.5)
	check("three negatives", `medianOf(-1, -2, -3)`,
		medianOf(-1, -2, -3), -2.0)

	// The same number several times.
	check("duplicates", `medianOf(2, 2, 2)`,
		medianOf(2, 2, 2), 2.0)

	// Numbers in random order: sorted they are 1, 3, 5, 15.
	check("random order", `medianOf(5, 15, 1, 3)`,
		medianOf(5, 15, 1, 3), 4.0)

	// One very big number does not move the median much.
	check("one huge number", `medianOf(1, 2, 1000)`,
		medianOf(1, 2, 1000), 2.0)

	// FindMedian can be called between AddNum calls on the SAME finder.
	finder := Constructor()
	finder.AddNum(10)
	finder.AddNum(20)
	fmt.Println(`-- same finder: AddNum(10), AddNum(20)`)
	check("same finder, after two numbers", `finder.FindMedian()`,
		finder.FindMedian(), 15.0)
	finder.AddNum(30)
	fmt.Println(`-- same finder: AddNum(30)`)
	check("same finder, after a third number", `finder.FindMedian()`,
		finder.FindMedian(), 20.0)

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
