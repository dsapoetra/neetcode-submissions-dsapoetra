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
	check("encoded format", `codec.Encode([]string{"neet", "code"})`,
		codec.Encode([]string{"neet", "code"}), "4#neet4#code")

	// Example 1 from the problem.
	check("example 1", `roundTrip([]string{"neet", "code", "love", "you"})`,
		roundTrip([]string{"neet", "code", "love", "you"}), `["neet" "code" "love" "you"]`)

	// Example 2 from the problem: a string with a special character.
	check("example 2", `roundTrip([]string{"we", "say", ":", "yes"})`,
		roundTrip([]string{"we", "say", ":", "yes"}), `["we" "say" ":" "yes"]`)

	// An empty list.
	check("empty list", `roundTrip([]string{})`,
		roundTrip([]string{}), `[]`)

	// A list with one empty string is NOT the same as an empty list.
	check("one empty string", `roundTrip([]string{""})`,
		roundTrip([]string{""}), `[""]`)
	check("two empty strings", `roundTrip([]string{"", ""})`,
		roundTrip([]string{"", ""}), `["" ""]`)

	// Tricky: the strings contain '#', the character we use as a separator.
	check("strings with #", `roundTrip([]string{"a#b", "#", "##"})`,
		roundTrip([]string{"a#b", "#", "##"}), `["a#b" "#" "##"]`)

	// Tricky: the strings contain digits, which could be mixed up with the length.
	check("strings with digits", `roundTrip([]string{"12", "3#4", "5"})`,
		roundTrip([]string{"12", "3#4", "5"}), `["12" "3#4" "5"]`)

	// Spaces must survive.
	check("strings with spaces", `roundTrip([]string{"hello world", " "})`,
		roundTrip([]string{"hello world", " "}), `["hello world" " "]`)

	// A string longer than 9 characters has a length with two digits.
	check("long string", `roundTrip([]string{"abcdefghijklmnop"})`,
		roundTrip([]string{"abcdefghijklmnop"}), `["abcdefghijklmnop"]`)

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
