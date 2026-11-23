// Idea: depth of a tree = 1 (this node) + the deeper of its two subtrees.
// Time: O(n).
func maxDepth(root *TreeNode) int {
	// an empty tree has depth 0
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	// take the deeper side and add 1 for the current node
	if left > right {
		return left + 1
	}
	return right + 1
}
