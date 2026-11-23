// Idea: add like on paper. The lists store the digits backwards (ones
// first), so we can go from left to right: add the two digits plus the
// carry, write down the last digit of the sum, and carry the rest.
// Keep going while there are digits left OR there is still a carry.
// Time: O(longer list).
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{} // fake first node, makes the code simpler
	tail := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		tail.Next = &ListNode{Val: sum % 10} // write the last digit
		tail = tail.Next
		carry = sum / 10 // carry the rest (0 or 1)
	}
	return dummy.Next
}
