// Idea: use the BST order. Start at the root:
//   - if p and q are both smaller, the answer is in the left subtree
//   - if p and q are both bigger, the answer is in the right subtree
//   - otherwise they split here (or one of them IS this node) -> answer
// Time: O(height of tree).
func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	cur := root
	for cur != nil {
		if p.Val < cur.Val && q.Val < cur.Val {
			cur = cur.Left // both on the left
		} else if p.Val > cur.Val && q.Val > cur.Val {
			cur = cur.Right // both on the right
		} else {
			return cur // they split here
		}
	}
	return nil
}
