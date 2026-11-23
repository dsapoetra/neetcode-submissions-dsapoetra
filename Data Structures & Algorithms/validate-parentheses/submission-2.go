// Idea: use a stack. Push every opening bracket. When we see a closing
// bracket, it must match the opening bracket on TOP of the stack
// (the most recent one that is still open). At the end nothing may be
// left open.
// Time: O(n).
func isValid(s string) bool {
	// closing bracket -> the opening bracket it needs
	pairs := map[byte]byte{')': '(', ']': '[', '}': '{'}
	stack := []byte{}

	for i := 0; i < len(s); i++ {
		ch := s[i]
		opening, isClosing := pairs[ch]

		if !isClosing {
			// opening bracket: remember it
			stack = append(stack, ch)
			continue
		}

		// closing bracket: the top of the stack must be its partner
		if len(stack) == 0 || stack[len(stack)-1] != opening {
			return false
		}
		stack = stack[:len(stack)-1] // pop
	}

	// valid only if every opening bracket was closed
	return len(stack) == 0
}
