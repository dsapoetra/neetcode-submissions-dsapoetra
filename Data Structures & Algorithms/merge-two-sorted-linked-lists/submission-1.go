// Idea: compare the first nodes of both lists and always attach the
// smaller one to the result. A dummy first node saves us from special
// code for the very first node.
// Time: O(n + m).
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{} // fake first node
	tail := dummy        // last node of the result so far

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}

	// one list is finished: attach what is left of the other
	if list1 != nil {
		tail.Next = list1
	} else {
		tail.Next = list2
	}

	return dummy.Next // skip the fake node
}
