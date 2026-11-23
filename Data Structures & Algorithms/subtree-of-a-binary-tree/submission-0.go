// Idea: for every node in the big tree, check if the tree starting at
// that node is exactly the same as subRoot.
// Time: O(nodes in root * nodes in subRoot).
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	// an empty tree is a subtree of anything
	if subRoot == nil {
		return true
	}
	// nothing left to search in
	if root == nil {
		return false
	}
	// does the tree starting here match?
	if sameTree(root, subRoot) {
		return true
	}
	// otherwise try the left and right children
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

// sameTree returns true if the two trees are exactly the same.
func sameTree(a *TreeNode, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil || a.Val != b.Val {
		return false
	}
	return sameTree(a.Left, b.Left) && sameTree(a.Right, b.Right)
}
