// Problem:    Detect Squares
// Difficulty: Medium
// Link:       https://neetcode.io/problems/count-squares
// Source:     Data Structures & Algorithms/count-squares/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Build a structure that receives points on a grid (the same point can be
// added several times):
//
//	Add(point)     adds a point [x, y]
//	Count(point)   the number of ways to pick three ADDED points that,
//	               together with the given point, form a square with sides
//	               parallel to the x and y axis (and with a size bigger than 0)
//
// Example:
//
//	Input:  Add([1, 1]), Add([2, 2]), Add([1, 2]), Count([2, 1]), Count([3, 3]),
//	        Add([2, 2]), Count([2, 1])
//	Output: -, -, -, 1, 0, -, 2
//
// How to run (from the "local" folder):
//
//	go run ./129-count-squares
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea: remember how often each point was added (count), plus a list of
// all added points.
// A square is fixed by two opposite corners. For a query point (x, y),
// try every added point (px, py) as the DIAGONAL corner. It only works if
// the horizontal and vertical distance are the same (and not 0).
// The other two corners are then (x, py) and (px, y). Each of them may
// exist several times, so multiply their counts.
// Going through the list of points (with duplicates) counts a diagonal
// point that was added twice two times, which is what we want.
// Add: O(1). Count: O(number of added points).

type CountSquares struct {
	count  map[[2]int]int // how many times each point was added
	points [][2]int       // every added point, duplicates included
}

func Constructor() CountSquares {
	return CountSquares{count: map[[2]int]int{}}
}

func (this *CountSquares) Add(point []int) {
	p := [2]int{point[0], point[1]}
	this.count[p]++
	this.points = append(this.points, p)
}

func (this *CountSquares) Count(point []int) int {
	x, y := point[0], point[1]
	total := 0

	for _, p := range this.points {
		px, py := p[0], p[1]

		// p must be a diagonal corner: same distance in x and in y, and not 0
		dx := px - x
		dy := py - y
		if dx == 0 || (dx != dy && dx != -dy) {
			continue
		}

		// the two remaining corners
		total += this.count[[2]int{x, py}] * this.count[[2]int{px, y}]
	}
	return total
}

// ===================== Test cases =====================

func main() {
	// The object remembers all added points, so the lines starting with
	// "--" show every Add in the output.

	// Example from the problem.
	squares := Constructor()
	squares.Add([]int{1, 1})
	squares.Add([]int{2, 2})
	squares.Add([]int{1, 2})
	fmt.Println(`-- squares.Add([]int{1, 1}); squares.Add([]int{2, 2}); squares.Add([]int{1, 2})`)
	check("one square with corner (2,1)", `squares.Count([]int{2, 1})`,
		squares.Count([]int{2, 1}), 1)
	check("no square with corner (3,3)", `squares.Count([]int{3, 3})`,
		squares.Count([]int{3, 3}), 0)
	squares.Add([]int{2, 2})
	fmt.Println(`-- squares.Add([]int{2, 2})   (the same point again)`)
	check("a duplicate point makes a second square", `squares.Count([]int{2, 1})`,
		squares.Count([]int{2, 1}), 2)

	// The classic example.
	classic := Constructor()
	classic.Add([]int{3, 10})
	classic.Add([]int{11, 2})
	classic.Add([]int{3, 2})
	fmt.Println(`-- classic.Add([]int{3, 10}); classic.Add([]int{11, 2}); classic.Add([]int{3, 2})`)
	check("classic: one square", `classic.Count([]int{11, 10})`,
		classic.Count([]int{11, 10}), 1)
	check("classic: no square", `classic.Count([]int{14, 8})`,
		classic.Count([]int{14, 8}), 0)
	classic.Add([]int{11, 2})
	fmt.Println(`-- classic.Add([]int{11, 2})   (the same point again)`)
	check("classic: two squares", `classic.Count([]int{11, 10})`,
		classic.Count([]int{11, 10}), 2)

	// A square needs a size: the same point four times is not a square.
	single := Constructor()
	single.Add([]int{5, 5})
	fmt.Println(`-- single.Add([]int{5, 5})`)
	check("a point alone is not a square", `single.Count([]int{5, 5})`,
		single.Count([]int{5, 5}), 0)

	// A rectangle that is not a square does not count.
	rectangle := Constructor()
	rectangle.Add([]int{0, 1})
	rectangle.Add([]int{2, 1})
	rectangle.Add([]int{2, 0})
	fmt.Println(`-- rectangle.Add([]int{0, 1}); rectangle.Add([]int{2, 1}); rectangle.Add([]int{2, 0})`)
	check("a 2 x 1 rectangle is not a square", `rectangle.Count([]int{0, 0})`,
		rectangle.Count([]int{0, 0}), 0)

	// Squares in two directions from the same corner: one above, one below.
	both := Constructor()
	both.Add([]int{0, 2})
	both.Add([]int{2, 2})
	both.Add([]int{2, 0})
	both.Add([]int{0, -2})
	both.Add([]int{2, -2})
	fmt.Println(`-- both: added (0,2), (2,2), (2,0), (0,-2), (2,-2)`)
	check("one square above and one below", `both.Count([]int{0, 0})`,
		both.Count([]int{0, 0}), 2)

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
