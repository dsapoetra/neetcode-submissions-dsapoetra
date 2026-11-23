// Idea: keep a list ("queue") of INDEXES of numbers that could still be
// the maximum of a window. The numbers at those indexes are always in
// decreasing order, so the front of the queue is the current maximum.
// For every new number:
//   1) remove smaller numbers from the back: they can never be the maximum
//      again, because the new number is bigger AND stays longer
//   2) add the new index at the back
//   3) remove the front if it has slid out of the window
// Time: O(n), every index is added once and removed at most once.
func maxSlidingWindow(nums []int, k int) []int {
	res := []int{}
	queue := []int{} // indexes; their numbers are in decreasing order

	for right := 0; right < len(nums); right++ {
		// 1) drop smaller numbers from the back
		for len(queue) > 0 && nums[queue[len(queue)-1]] < nums[right] {
			queue = queue[:len(queue)-1]
		}

		// 2) add the new index
		queue = append(queue, right)

		// 3) the window is [right-k+1 .. right]; drop the front if it is outside
		if queue[0] < right-k+1 {
			queue = queue[1:]
		}

		// once the first full window is reached, the front is its maximum
		if right >= k-1 {
			res = append(res, nums[queue[0]])
		}
	}
	return res
}
