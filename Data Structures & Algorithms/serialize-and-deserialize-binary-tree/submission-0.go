// Idea: write the tree in preorder (node, left, right) separated by
// commas, and write "N" for an empty child.
// Example:   1
//           / \
//          2   3     ->  "1,2,N,N,3,N,N"
// Because the empty children are written too, we can rebuild the tree
// by reading the values back in the same order.

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "N"
	}
	// node, then left subtree, then right subtree
	return strconv.Itoa(root.Val) + "," + this.serialize(root.Left) + "," + this.serialize(root.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	parts := strings.Split(data, ",")
	index := 0 // the next value to read

	var build func() *TreeNode
	build = func() *TreeNode {
		value := parts[index]
		index++

		if value == "N" {
			return nil
		}

		number, _ := strconv.Atoi(value)
		node := &TreeNode{Val: number}
		node.Left = build()  // same order as serialize: left first
		node.Right = build() // then right
		return node
	}

	return build()
}
