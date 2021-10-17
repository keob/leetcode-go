package solution230

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BST struct {
	root    *TreeNode
	nodeNum map[*TreeNode]int
}

func (bst *BST) countNodeNum(node *TreeNode) int {
	if node == nil {
		return 0
	}
	bst.nodeNum[node] = 1 + bst.countNodeNum(node.Left) + bst.countNodeNum(node.Right)
	return bst.nodeNum[node]
}

func (bst *BST) kthSmallest(k int) int {
	node := bst.root
	for {
		leftNodeNum := bst.nodeNum[node.Left]
		if leftNodeNum < k-1 {
			node = node.Right
			k -= leftNodeNum + 1
		} else if leftNodeNum == k-1 {
			return node.Val
		} else {
			node = node.Left
		}
	}
}

func kthSmallest(root *TreeNode, k int) int {
	t := &BST{root, map[*TreeNode]int{}}
	t.countNodeNum(root)
	return t.kthSmallest(k)
}
