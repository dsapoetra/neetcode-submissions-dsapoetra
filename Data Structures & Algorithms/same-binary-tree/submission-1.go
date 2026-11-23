// Idea: two trees are the same if the roots have the same value
// and the left subtrees are the same and the right subtrees are the same.
// Time: O(n).
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// both empty -> same
	if p == nil && q == nil {
		return true
	}
	// only one is empty, or the values differ -> not the same
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	// check both subtrees
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
