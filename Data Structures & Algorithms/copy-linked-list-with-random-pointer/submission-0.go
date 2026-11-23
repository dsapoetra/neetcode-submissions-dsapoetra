// Idea: copy in two passes, with a map from each old node to its copy.
//   pass 1: create a copy of every node (only the value)
//   pass 2: set Next and Random of every copy, using the map to find
//           the copy that belongs to the old node's Next / Random
// Time: O(n), extra space: O(n).
func copyRandomList(head *Node) *Node {
	copies := map[*Node]*Node{} // old node -> its copy

	// pass 1: create all the copies
	for cur := head; cur != nil; cur = cur.Next {
		copies[cur] = &Node{Val: cur.Val}
	}

	// pass 2: connect the copies.
	// copies[nil] gives nil, which is exactly what we want for "no node".
	for cur := head; cur != nil; cur = cur.Next {
		copies[cur].Next = copies[cur.Next]
		copies[cur].Random = copies[cur.Random]
	}

	return copies[head]
}
