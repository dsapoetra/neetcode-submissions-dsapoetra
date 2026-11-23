// Idea: we already know how to merge TWO sorted lists.
// So merge the lists in pairs: list 0 with list 1, list 2 with list 3, ...
// That cuts the number of lists in half. Repeat until one list is left.
// (Merging everything into one big list one by one also works, but it
// walks over the big list again and again, which is too slow.)
// Time: O(total nodes * log k).
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	for len(lists) > 1 {
		merged := []*ListNode{}

		// merge neighbors: (0,1), (2,3), (4,5), ...
		for i := 0; i < len(lists); i += 2 {
			if i+1 < len(lists) {
				merged = append(merged, mergeTwo(lists[i], lists[i+1]))
			} else {
				// odd one out: no partner, keep it for the next round
				merged = append(merged, lists[i])
			}
		}

		lists = merged // now we have half as many lists
	}
	return lists[0]
}

// mergeTwo merges two sorted lists into one sorted list.
func mergeTwo(a *ListNode, b *ListNode) *ListNode {
	dummy := &ListNode{} // fake first node, makes the code simpler
	tail := dummy

	// always attach the smaller node
	for a != nil && b != nil {
		if a.Val <= b.Val {
			tail.Next = a
			a = a.Next
		} else {
			tail.Next = b
			b = b.Next
		}
		tail = tail.Next
	}

	// one list is finished: attach what is left of the other
	if a != nil {
		tail.Next = a
	} else {
		tail.Next = b
	}
	return dummy.Next
}
