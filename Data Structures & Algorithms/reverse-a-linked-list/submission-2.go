// Idea: walk through the list and turn every arrow around, so each
// node points to the node BEFORE it instead of the node after it.
// Time: O(n).
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode // the node before cur (nothing at the start)
	cur := head

	for cur != nil {
		next := cur.Next // save the next node before we lose it
		cur.Next = prev  // turn the arrow around
		prev = cur       // move both pointers one step forward
		cur = next
	}

	// prev is the old last node, which is the new head
	return prev
}
