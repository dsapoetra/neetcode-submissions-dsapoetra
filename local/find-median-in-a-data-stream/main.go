// Problem: Find Median from Data Stream
// Link:    https://neetcode.io/problems/find-median-in-a-data-stream
// Source:  Data Structures & Algorithms/find-median-in-a-data-stream/submission-0.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./find-median-in-a-data-stream
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
	// Example from the problem.
	mf := Constructor()
	mf.AddNum(1)
	check("one number: the number itself", mf.FindMedian(), 1.0)
	mf.AddNum(2)
	check("two numbers: the average", mf.FindMedian(), 1.5)
	mf.AddNum(3)
	check("three numbers: the middle one", mf.FindMedian(), 2.0)

	// Numbers arriving from big to small.
	mf = Constructor()
	mf.AddNum(5)
	mf.AddNum(4)
	mf.AddNum(3)
	mf.AddNum(2)
	check("descending order", mf.FindMedian(), 3.5)

	// Negative numbers.
	mf = Constructor()
	mf.AddNum(-1)
	mf.AddNum(-2)
	check("two negatives", mf.FindMedian(), -1.5)
	mf.AddNum(-3)
	check("three negatives", mf.FindMedian(), -2.0)

	// The same number several times.
	mf = Constructor()
	mf.AddNum(2)
	mf.AddNum(2)
	mf.AddNum(2)
	check("duplicates", mf.FindMedian(), 2.0)

	// Numbers in random order: sorted they are 1, 3, 5, 15.
	mf = Constructor()
	mf.AddNum(5)
	mf.AddNum(15)
	mf.AddNum(1)
	mf.AddNum(3)
	check("random order", mf.FindMedian(), 4.0)

	// One very big number does not move the median much.
	mf = Constructor()
	mf.AddNum(1)
	mf.AddNum(2)
	mf.AddNum(1000)
	check("one huge number", mf.FindMedian(), 2.0)

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
