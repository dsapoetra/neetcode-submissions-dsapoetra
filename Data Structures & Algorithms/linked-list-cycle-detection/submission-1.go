// Idea (slow and fast pointer): slow moves 1 step, fast moves 2 steps.
//   - no cycle: fast reaches the end of the list
//   - cycle:    fast keeps going around and catches up with slow,
//               like a faster runner on a round track
// Time: O(n), no extra memory.
func hasCycle(head *ListNode) bool {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next      // 1 step
		fast = fast.Next.Next // 2 steps

		if slow == fast {
			return true // they met: there is a cycle
		}
	}

	// fast reached the end: no cycle
	return false
}
