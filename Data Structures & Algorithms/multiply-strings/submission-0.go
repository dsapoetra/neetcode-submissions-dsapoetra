// Idea: multiply like on paper. The digit at position i of num1 times the
// digit at position j of num2 belongs at position i + j + 1 of the
// result (and its carry at position i + j).
// A number with a digits times a number with b digits has at most a + b
// digits, so the result fits in a list of that length.
// Time: O(len(num1) * len(num2)).
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	result := make([]int, len(num1)+len(num2))

	// go from the last digits to the first, like on paper
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			product := int(num1[i]-'0') * int(num2[j]-'0')

			sum := product + result[i+j+1] // add what is already there
			result[i+j+1] = sum % 10       // keep one digit
			result[i+j] += sum / 10        // carry the rest to the left
		}
	}

	// turn the digits into text, without the zeros at the front
	text := ""
	for i, digit := range result {
		if text == "" && digit == 0 && i < len(result)-1 {
			continue
		}
		text += string(rune('0' + digit))
	}
	return text
}
