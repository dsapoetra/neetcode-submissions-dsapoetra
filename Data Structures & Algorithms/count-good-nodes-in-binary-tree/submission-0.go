// Idea: a node is "good" if no node on the way from the root down to it
// has a bigger value. So walk down the tree and carry along the biggest
// value seen so far on this path. A node is good if its value is at
// least that big.
// Time: O(n).
func goodNodes(root *TreeNode) int {
	return countGood(root, root.Val)
}

// countGood counts the good nodes in this subtree.
// maxSoFar is the biggest value on the path from the root to this node.
func countGood(node *TreeNode, maxSoFar int) int {
	if node == nil {
		return 0
	}

	count := 0
	if node.Val >= maxSoFar {
		count = 1          // this node is good
		maxSoFar = node.Val // and it is the new biggest value for the nodes below
	}

	count += countGood(node.Left, maxSoFar)
	count += countGood(node.Right, maxSoFar)
	return count
}
