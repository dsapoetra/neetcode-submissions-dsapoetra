// Idea: handle one group of k nodes at a time.
//   1) check that k nodes are left (if not, we are done)
//   2) reverse the arrows inside the group
//   3) connect the reversed group with the part before and after it
// groupPrev is always the node right before the group we are working on.
// Time: O(n), no extra memory.
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head} // fake node before the head
	groupPrev := dummy

	for {
		// 1) find the last node of this group
		kth := groupPrev
		for i := 0; i < k && kth != nil; i++ {
			kth = kth.Next
		}
		if kth == nil {
			break // fewer than k nodes left: leave them as they are
		}
		groupNext := kth.Next // first node AFTER the group

		// 2) reverse the group. Starting with prev = groupNext makes the
		//    old first node point to the rest of the list.
		prev := groupNext
		cur := groupPrev.Next
		for cur != groupNext {
			next := cur.Next
			cur.Next = prev
			prev = cur
			cur = next
		}

		// 3) the old first node is now the last node of the group
		oldFirst := groupPrev.Next
		groupPrev.Next = kth // kth is the new first node
		groupPrev = oldFirst // the next group starts after it
	}
	return dummy.Next
}
