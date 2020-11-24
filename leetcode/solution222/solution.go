package solution222

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lH, rH := 0, 0
	lNode, rNode := root, root

	for lNode != nil {
		lH++
		lNode = lNode.Left
	}
	for rNode != nil {
		rH++
		rNode = rNode.Right
	}
	if lH == rH {
		return 1<<lH - 1
	}

	return 1 + countNodes(root.Left) + countNodes(root.Right)
}
