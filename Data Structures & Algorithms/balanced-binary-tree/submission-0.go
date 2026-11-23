// Idea: a tree is balanced if, at EVERY node, the heights of the left and
// right subtree differ by at most 1.
// The helper returns the height of a subtree, or -1 as a signal that
// something below is already unbalanced. That way we visit every node
// only once.
// Time: O(n).
func isBalanced(root *TreeNode) bool {
	return balancedHeight(root) != -1
}

// balancedHeight returns the height of the tree, or -1 if it is not balanced.
func balancedHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left := balancedHeight(node.Left)
	right := balancedHeight(node.Right)

	// a subtree is already unbalanced: pass the bad news up
	if left == -1 || right == -1 {
		return -1
	}

	// the two sides differ by more than 1: unbalanced here
	if left-right > 1 || right-left > 1 {
		return -1
	}

	if left > right {
		return left + 1
	}
	return right + 1
}
