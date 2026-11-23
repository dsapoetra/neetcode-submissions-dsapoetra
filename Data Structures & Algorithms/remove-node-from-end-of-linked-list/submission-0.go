// Idea (two pointers): move "fast" n steps ahead first. Then move "slow"
// and "fast" together. When fast reaches the last node, slow is right
// before the node we must remove.
// A dummy node before head makes removing the head itself easy.
// Time: O(n).
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	slow := dummy
	fast := dummy

	// move fast n steps ahead
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	// move both until fast is at the last node
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	// slow.Next is the node to remove: skip over it
	slow.Next = slow.Next.Next
	return dummy.Next
}
