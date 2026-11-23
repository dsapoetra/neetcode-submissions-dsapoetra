// Idea: swap the left and right child of every node (recursion).
// Time: O(n).
func invertTree(root *TreeNode) *TreeNode {
	// empty tree: nothing to do
	if root == nil {
		return nil
	}

	// swap the two children
	root.Left, root.Right = root.Right, root.Left

	// do the same for both subtrees
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
