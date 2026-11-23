// Idea: an inorder traversal (left, node, right) of a BST visits the
// values in sorted order. So collect them and take the k-th one.
// Time: O(n).
func kthSmallest(root *TreeNode, k int) int {
	values := inorderValues(root)
	return values[k-1] // k starts from 1, the slice starts from 0
}

// inorderValues returns all values of the tree in sorted order.
func inorderValues(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}
	values := inorderValues(node.Left)                     // smaller values first
	values = append(values, node.Val)                      // then this node
	values = append(values, inorderValues(node.Right)...) // then bigger values
	return values
}
