// Idea: the longest path that bends at a node goes down its left side as
// deep as possible and down its right side as deep as possible:
//   path through node = height of left subtree + height of right subtree
// So compute the height of every subtree (recursion) and, at every node,
// check whether left height + right height is the best so far.
// Time: O(n).
func diameterOfBinaryTree(root *TreeNode) int {
	best := 0

	// height returns the number of nodes on the longest way down from node.
	var height func(node *TreeNode) int
	height = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := height(node.Left)
		right := height(node.Right)

		// the path that bends at this node has left + right edges
		if left+right > best {
			best = left + right
		}

		// height of this node = 1 + the taller side
		if left > right {
			return left + 1
		}
		return right + 1
	}

	height(root)
	return best
}
