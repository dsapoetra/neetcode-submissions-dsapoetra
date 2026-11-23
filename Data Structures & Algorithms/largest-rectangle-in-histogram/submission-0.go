// Idea: for every bar we want to know how far a rectangle of that bar's
// height can stretch to the left and to the right.
// Keep a stack of bars that are still "open" (their rectangle can still
// grow to the right). Heights in the stack go up.
// When a new bar is LOWER than the bar on top, the top bar's rectangle
// must end here: pop it and compute its area.
// The new bar can stretch back to where the popped bars started, because
// they were all taller.
// Time: O(n), every bar is pushed and popped once.

type openBar struct {
	start  int // the leftmost index this height reaches
	height int
}

func largestRectangleArea(heights []int) int {
	best := 0
	stack := []openBar{}

	for i, height := range heights {
		start := i

		// taller bars on the stack end here
		for len(stack) > 0 && stack[len(stack)-1].height > height {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			area := top.height * (i - top.start)
			if area > best {
				best = area
			}
			start = top.start // the new, lower bar reaches back this far
		}

		stack = append(stack, openBar{start, height})
	}

	// bars still on the stack reach all the way to the end
	for _, bar := range stack {
		area := bar.height * (len(heights) - bar.start)
		if area > best {
			best = area
		}
	}
	return best
}
