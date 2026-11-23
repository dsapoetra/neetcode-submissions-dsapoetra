// Idea: use a stack of numbers.
//   - a number   -> push it
//   - an operator -> pop the last two numbers, calculate, push the result
// Careful with the order: the number popped FIRST is the RIGHT side
// (for "a b -" the stack gives b first, and the answer is a - b).
// Time: O(n).
func evalRPN(tokens []string) int {
	stack := []int{}

	for _, token := range tokens {
		if token == "+" || token == "-" || token == "*" || token == "/" {
			// pop the right side, then the left side
			right := stack[len(stack)-1]
			left := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			result := 0
			if token == "+" {
				result = left + right
			} else if token == "-" {
				result = left - right
			} else if token == "*" {
				result = left * right
			} else {
				result = left / right // Go cuts off the decimals (rounds toward zero)
			}
			stack = append(stack, result)
		} else {
			number, _ := strconv.Atoi(token)
			stack = append(stack, number)
		}
	}
	return stack[0]
}
