// Problem: Encode and Decode Strings
// Link:    https://neetcode.io/problems/string-encode-and-decode
// Source:  Data Structures & Algorithms/string-encode-and-decode/submission-0.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./string-encode-and-decode
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
	"strconv"
)

// ===================== Solution =====================

// Idea: write every string as  <length>#<string>.
// Example: ["neet", "code"] -> "4#neet4#code"
// When decoding we read the number up to '#', so we know exactly how
// many characters belong to the string, even if it contains '#' itself.

type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	encoded := ""
	for _, str := range strs {
		encoded += strconv.Itoa(len(str)) + "#" + str
	}
	return encoded
}

func (s *Solution) Decode(encoded string) []string {
	res := []string{}
	i := 0
	for i < len(encoded) {
		// find the '#': the digits before it are the length
		j := i
		for encoded[j] != '#' {
			j++
		}
		length, _ := strconv.Atoi(encoded[i:j])

		// the string starts right after '#' and is "length" characters long
		start := j + 1
		res = append(res, encoded[start:start+length])

		// jump to the beginning of the next piece
		i = start + length
	}
	return res
}

// ===================== Test cases =====================

func main() {
	codec := &Solution{}

	// roundTrip encodes the list, decodes it again, and prints the result
	// with quotes (so "" and " " are visible). We must get the same list back.
	roundTrip := func(strs []string) string {
		return fmt.Sprintf("%q", codec.Decode(codec.Encode(strs)))
	}

	// What the encoded string looks like: <length>#<string> for every string.
	check("encoded format", codec.Encode([]string{"neet", "code"}), "4#neet4#code")

	// Example 1 from the problem.
	check("example 1", roundTrip([]string{"neet", "code", "love", "you"}), `["neet" "code" "love" "you"]`)

	// Example 2 from the problem: a string with a special character.
	check("example 2", roundTrip([]string{"we", "say", ":", "yes"}), `["we" "say" ":" "yes"]`)

	// An empty list.
	check("empty list", roundTrip([]string{}), `[]`)

	// A list with one empty string is NOT the same as an empty list.
	check("one empty string", roundTrip([]string{""}), `[""]`)
	check("two empty strings", roundTrip([]string{"", ""}), `["" ""]`)

	// Tricky: the strings contain '#', the character we use as a separator.
	check("strings with #", roundTrip([]string{"a#b", "#", "##"}), `["a#b" "#" "##"]`)

	// Tricky: the strings contain digits, which could be mixed up with the length.
	check("strings with digits", roundTrip([]string{"12", "3#4", "5"}), `["12" "3#4" "5"]`)

	// Spaces must survive.
	check("strings with spaces", roundTrip([]string{"hello world", " "}), `["hello world" " "]`)

	// A string longer than 9 characters has a length with two digits.
	check("long string", roundTrip([]string{"abcdefghijklmnop"}), `["abcdefghijklmnop"]`)

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
