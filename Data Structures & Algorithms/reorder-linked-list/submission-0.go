// Idea: put all the nodes in a slice so we can reach them by index.
// Then use two pointers, one at the front and one at the back, and
// link them in the order: first, last, second, second-last, ...
// Time: O(n), extra space: O(n).
func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	// collect all nodes in a slice
	nodes := []*ListNode{}
	for cur := head; cur != nil; cur = cur.Next {
		nodes = append(nodes, cur)
	}

	front := 0
	back := len(nodes) - 1
	for front < back {
		// front node points to the back node
		nodes[front].Next = nodes[back]
		front++
		if front == back {
			break
		}
		// back node points to the next front node
		nodes[back].Next = nodes[front]
		back--
	}

	// the node in the middle is now the last node
	nodes[front].Next = nil
}
