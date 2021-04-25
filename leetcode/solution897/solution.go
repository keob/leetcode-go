package solution897

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {
	dummyNode := &TreeNode{}
	resNode := dummyNode

	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)

		resNode.Right = node
		node.Left = nil
		resNode = node

		inorder(node.Right)
	}

	inorder(root)

	return dummyNode.Right
}
